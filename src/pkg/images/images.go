package images

import (
	"github.com/pkg/errors"
	"log"
	"strings"

	"encoding/base64"
	"fmt"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"hubimage/src/client"
	"hubimage/src/config"
	"hubimage/src/model"
	"hubimage/src/pkg/projects"
	"hubimage/swagger"
	"io/ioutil"
	"os"
	"time"
)

// CreateImage upload image
func CreateImage(request *model.RequestImageCreate) (*model.ImageDetail, error) {
	log.Printf("call CreateImage:%v", request)
	if request.ProjectName == "" || request.Name == "" || request.Tag == "" {
		return &model.ImageDetail{}, errors.New("Please specified projectName, imageName and imageTag")
	}

	if request.File != "" {
		file, err := os.Open(request.File)
		if err != nil {
			return &model.ImageDetail{}, err
		}
		defer file.Close()
		request.Reader = file
	} else if request.Reader == nil {
		return &model.ImageDetail{}, errors.New("Please specified image file or reader")
	}

	// project should exist
	exist, err := projects.HeadProject(request.ProjectName)
	if err != nil {
		return &model.ImageDetail{}, err
	}
	if !exist {
		msg := fmt.Sprintf("project \"%s\" not exist", request.ProjectName)
		return &model.ImageDetail{}, errors.New(msg)
	}

	dockerClient := client.GetDockerClient()
	conf := config.GetConfig()
	// docker login
	auth := types.AuthConfig{
		Username:      conf.Username,
		Password:      conf.Password,
		Email:         conf.Email,
		ServerAddress: conf.HarborHost,
	}
	authBody, err := dockerClient.RegistryLogin(context.Background(), auth)
	if err != nil {
		return &model.ImageDetail{}, err
	}
	log.Printf("authBody: %#v---%s", authBody, authBody.IdentityToken)
	// test save
	//
	//reader, err := dockerClient.ImageSave(ctx, []string{"busybox:latest"})
	//if err != nil {
	//	log.Printf("Failed to save image %v", err)
	//}
	//defer reader.Close()
	//request.Reader = reader

	// docker load
	resp, err := dockerClient.ImageLoad(context.Background(), request.Reader, true)
	if err != nil {
		return &model.ImageDetail{}, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &model.ImageDetail{}, err
	}
	log.Printf("result: %s", result)

	results := strings.Split(string(result), "sha256:")
	imageID := results[1]
	imageID = strings.Replace(imageID, " ", "", -1)
	imageID = strings.Replace(imageID, "\n", "", -1)

	// docker tag
	target := fmt.Sprintf("%s/%s/%s:%s", conf.HarborHost, request.ProjectName, request.Name, request.Tag)
	err = dockerClient.ImageTag(context.Background(), imageID, target)
	if err != nil {
		return &model.ImageDetail{}, err
	}
	log.Printf("docker tag target: %s", target)

	// docker push
	registryAuth := fmt.Sprintf("{ \"username\": \"%s\", \"password\": \"%s\", \"email\": \"%s\" }",
		conf.Username, conf.Password, conf.Email)
	data := []byte(registryAuth)
	encodeString := base64.StdEncoding.EncodeToString(data)
	pushRes, err := dockerClient.ImagePush(context.Background(), target, types.ImagePushOptions{
		All:          true,
		RegistryAuth: encodeString,
	})
	if err != nil {
		return &model.ImageDetail{}, err
	}
	log.Printf("********docker push Res: %#v----%s", pushRes, pushRes)

	if !request.EnableWait {
		return &model.ImageDetail{}, nil
	}

	// get
	// get tag
	harborClient := client.GetHarborClient()
	var imageTag swagger.DetailedTag
	found := false
	repoName := request.ProjectName + "/" + request.Name
	for {
		imageTags, _, err := harborClient.RepositoriesRepoNameTagsGet(repoName)
		if err != nil {
			msg := fmt.Sprintf("Failed to get RepositoriesRepoNameTagsGet: %v", err)
			return &model.ImageDetail{}, errors.New(msg)
		}
		for _, tag := range imageTags {
			if tag.Name == request.Tag {
				imageTag = tag
				found = true
			}
		}
		log.Printf("pushing...")
		if found {
			break
		}
		time.Sleep(5 * time.Second)
	}

	// get mainfest
	manifest, _, err := harborClient.RepositoriesRepoNameTagsTagManifestGet(repoName, imageTag.Name, "")
	if err != nil {
		msg := fmt.Sprintf("Failed to get RepositoriesRepoNameTagsTagManifestGet: %v", err)
		return &model.ImageDetail{}, errors.New(msg)
	}
	var imageDetail model.ImageDetail
	imageDetail.Image = imageTag
	imageDetail.Manifest = *manifest
	return &imageDetail, nil
}

// GetImageList list app
func GetImageList(request *model.RequestImageList) ([]swagger.Repository, error) {
	log.Printf("call GetImageList:%v", request)
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PageSize == 0 {
		request.PageSize = 100
	}
	if request.ProjectId == 0 {
		if request.ProjectName == "" {
			return []swagger.Repository{}, errors.New("please specified project name or project id.")
		}
		project, err := projects.GetProjectDetail(request.ProjectName)
		if err != nil {
			return []swagger.Repository{}, err
		}
		request.ProjectId = project.ProjectId
	}
	log.Printf("request.ProjectId:%v", request.ProjectId)

	harborClient := client.GetHarborClient()
	repoList, _, err := harborClient.RepositoriesGet(request.ProjectId, "", request.Page, request.PageSize)
	if err != nil {
		msg := fmt.Sprintf("Failed to get repositories: %v", err)
		return []swagger.Repository{}, errors.New(msg)
	}

	log.Printf("repoList:%#v", repoList)
	return repoList, nil
}

// GetImageDetail get detail app
func GetImageDetail(request *model.RequestImageDetail) ([]model.ImageDetail, error) {
	log.Printf("call GetImageDetail:%v", request)
	var err error
	var imageTags []swagger.DetailedTag
	var imageDetails []model.ImageDetail
	harborClient := client.GetHarborClient()

	repoName := request.ProjectName + "/" + request.ImageName
	// get tags
	if request.Tag == "" {
		imageTags, _, err = harborClient.RepositoriesRepoNameTagsGet(repoName)
		if err != nil {
			msg := fmt.Sprintf("Failed to get RepositoriesRepoNameTagsGet: %v", err)
			return imageDetails, errors.New(msg)
		}
	} else {
		imageTag, _, err := harborClient.RepositoriesRepoNameTagsTagGet(repoName, request.Tag)
		if err != nil {
			msg := fmt.Sprintf("Failed to get RepositoriesRepoNameTagsTagGet: %v", err)
			return imageDetails, errors.New(msg)
		}
		imageTags = append(imageTags, *imageTag)
	}
	log.Printf("imageTags:%#v", imageTags)

	// get manifest
	for _, imageTag := range imageTags {
		var imageDetail model.ImageDetail
		manifest, _, err := harborClient.RepositoriesRepoNameTagsTagManifestGet(repoName, imageTag.Name, "")
		if err != nil {
			msg := fmt.Sprintf("Failed to get RepositoriesRepoNameTagsTagManifestGet: %v", err)
			return imageDetails, errors.New(msg)
		}
		imageDetail.Image = imageTag
		imageDetail.Manifest = *manifest
		imageDetails = append(imageDetails, imageDetail)
	}

	log.Printf("imageDetails:%#v", imageDetails)
	return imageDetails, nil
}

// DeleteImage delete app
func DeleteImage(request *model.RequestImageDelete) error {
	log.Printf("call DeleteImage:%v", request)
	harborClient := client.GetHarborClient()

	repoName := request.ProjectName + "/" + request.ImageName
	if request.Tag == "" {
		res, err := harborClient.RepositoriesRepoNameDelete(repoName)
		if err != nil {
			return err
		}
		log.Printf("res:%#v", res)
	} else {
		res, err := harborClient.RepositoriesRepoNameTagsTagDelete(repoName, request.Tag)
		if err != nil {
			return err
		}
		log.Printf("res:%#v", res)
	}

	return nil
}
