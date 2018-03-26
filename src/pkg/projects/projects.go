package projects

import (
	"log"

	"fmt"
	"github.com/pkg/errors"
	"hubimage/src/client"
	"hubimage/src/model"
	"hubimage/swagger"
)

// CreateProject create project
func CreateProject(request *model.RequestProjectCreate) (*swagger.Project, error) {
	log.Printf("call CreateProject:%#v", request)
	if request.Name == "" {
		return &swagger.Project{}, errors.New("please specified project name.")
	}

	harborClient := client.GetHarborClient()
	req := swagger.ProjectReq{}
	req.ProjectName = request.Name
	if request.Public {
		req.Public = int32(1)
	} else {
		req.Public = int32(0)
	}

	log.Printf("req:%#v", req)
	APIResponse, err := harborClient.ProjectsPost(req)
	if err != nil {
		return &swagger.Project{}, err
	}
	log.Printf("post APIResponse:%#v", APIResponse.Response)
	if APIResponse.Response.StatusCode/100 != 2 {
		msg := fmt.Sprintf("Failed to create project %s: %v", request.Name, APIResponse.Response.Status)
		return &swagger.Project{}, errors.New(msg)
	}

	// get
	project, err := getProject(harborClient, request.Name, false, request.Public)
	if err != nil {
		return &swagger.Project{}, err
	}

	log.Printf("get project:%#v", project)
	return project, nil
}

// GetProjectList list project
func GetProjectList(request *model.RequestProjectList) ([]swagger.Project, error) {
	log.Printf("call GetProjectList:%#v", request)
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PageSize == 0 {
		request.PageSize = 100
	}
	harborClient := client.GetHarborClient()

	projects, APIResponse, err := harborClient.ProjectsGet(request.Name, request.All, request.Public,
		request.Owner, request.Page, request.PageSize)
	if err != nil {
		log.Printf("failed to GetProjectList:%v", err)
		return []swagger.Project{}, err
	}
	log.Printf("list APIResponse:%#v", APIResponse.Response)

	log.Printf("projects:%#v", projects)
	return projects, nil
}

func getProject(harborClient *swagger.ProductsApi, name string, all bool, public bool) (*swagger.Project, error) {
	log.Printf("call GetProject")
	projects, APIResponse, err := harborClient.ProjectsGet(name, all, public, "", 1, 20)
	if err != nil {
		log.Printf("Failed to GetProject:%v", err)
		return &swagger.Project{}, err
	}
	log.Printf("get APIResponse:%#v", APIResponse.Response)

	if len(projects) == 0 {
		msg := fmt.Sprintf("Error: the project %s not found", name)
		return &swagger.Project{}, errors.New(msg)
	}

	for _, item := range projects {
		if item.Name == name {
			return &item, nil
		}
	}

	msg := fmt.Sprintf("Failed to get project %s", name)
	return &swagger.Project{}, errors.New(msg)
}

// HeadProject check if project exist
func HeadProject(name string) (bool, error) {
	log.Printf("call HeadProject:%s", name)
	harborClient := client.GetHarborClient()

	_, err := getProject(harborClient, name, true, true)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetProjectDetail get detail project
func GetProjectDetail(name string) (*swagger.Project, error) {
	log.Printf("call GetProjectDetail:%s", name)
	harborClient := client.GetHarborClient()

	project, err := getProject(harborClient, name, true, true)
	if err != nil {
		return &swagger.Project{}, err
	}

	log.Printf("project:%#v", project)
	return project, nil
}

// DeleteProject delete project
func DeleteProject(name string) error {
	log.Printf("call DeleteProject: %s", name)
	harborClient := client.GetHarborClient()
	project, err := getProject(harborClient, name, true, true)
	if err != nil {
		return err
	}

	_, err = harborClient.ProjectsProjectIdDelete(int64(project.ProjectId))
	if err != nil {
		return err
	}

	log.Printf("success to delete project:%#v", project)
	return nil
}
