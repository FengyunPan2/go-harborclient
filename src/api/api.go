package api

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"

	"hubimage/src/model"
	"hubimage/src/pkg/images"
	"hubimage/src/pkg/projects"
	"hubimage/src/pkg/users"
)

// HandleCreateImage create Image
func HandleCreateImage(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleCreateImage")
	requestRaw := new(model.RequestImageCreate)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := images.CreateImage(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetImageList get image list
func HandleGetImageList(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetImageList")
	requestRaw := new(model.RequestImageList)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := images.GetImageList(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetImageDetail get detail image info
func HandleGetImageDetail(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetImageDetail")
	name := request.PathParameter("name")
	requestRaw := new(model.RequestImageDetail)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	requestRaw.ImageName = name
	result, err := images.GetImageDetail(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleDeleteImage delete app
func HandleDeleteImage(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleDeleteImage")
	name := request.PathParameter("name")
	requestRaw := new(model.RequestImageDelete)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	requestRaw.ImageName = name
	err = images.DeleteImage(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

/******************* projects *********************/
// HandleCreateProject create Project
func HandleCreateProject(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleCreateProject")
	requestRaw := new(model.RequestProjectCreate)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := projects.CreateProject(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetProjectList get Project list
func HandleGetProjectList(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetProjectList")
	requestRaw := new(model.RequestProjectList)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := projects.GetProjectList(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetProjectDetail get detail Project info
func HandleGetProjectDetail(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetProjectDetail")
	name := request.PathParameter("name")
	result, err := projects.GetProjectDetail(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleDeleteProject delete Project
func HandleDeleteProject(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleDeleteProject")
	name := request.PathParameter("name")
	err := projects.DeleteProject(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

/******************* users *********************/
// HandleCreateUser create user
func HandleCreateUser(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleCreateUser")
	requestRaw := new(model.RequestUserCreate)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := users.CreateUser(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetUserList get user list
func HandleGetUserList(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetUserList")
	requestRaw := new(model.RequestUserList)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := users.GetUserList(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetUserDetail get detail user info
func HandleGetUserDetail(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetUserDetail")
	name := request.PathParameter("name")
	result, err := users.GetUserDetail(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleDeleteUser delete user
func HandleDeleteUser(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleDeleteUser")
	name := request.PathParameter("name")
	err := users.DeleteUser(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

// test
func HandleTest(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleTest")

	response.WriteHeader(http.StatusOK)
}
