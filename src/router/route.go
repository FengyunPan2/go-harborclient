package router

import (
	"hubimage/src/api"
	"hubimage/src/model"
	"hubimage/src/router/filter"
	"hubimage/swagger"

	restful "github.com/emicklei/go-restful"
)

// CreateHTTPRouter create http router for app
func CreateHTTPRouter() *restful.Container {

	wsContainer := restful.NewContainer()
	wsContainer.Filter(filter.LogRequestAndResponse)
	ws := new(restful.WebService)

	ws.Path("/api/v1").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	/************** images *************/
	// create image
	ws.Route(ws.POST("/image").
		To(api.HandleCreateImage).
		Reads(model.RequestImageCreate{}).
		Writes(model.ImageDetail{}))

	// get image list
	ws.Route(ws.GET("/image").
		To(api.HandleGetImageList).
		Reads(model.RequestImageList{}).
		Writes(model.ResponseImageList{}))

	// get image detail
	ws.Route(ws.GET("/image/{name}").
		To(api.HandleGetImageDetail).
		Reads(model.RequestImageDetail{}).
		Writes(model.ResponseImageDetail{}))

	// delete image
	ws.Route(ws.DELETE("/image/{name}").
		To(api.HandleDeleteImage).
		Reads(model.RequestImageDelete{}))

	/************** projects *************/
	// create project
	ws.Route(ws.POST("/project").
		To(api.HandleCreateProject).
		Reads(model.RequestProjectCreate{}).
		Writes(swagger.Project{}))

	// get project list
	ws.Route(ws.GET("/project").
		To(api.HandleGetProjectList).
		Reads(model.RequestProjectList{}).
		Writes(model.ResponseProjectList{}))

	// get project detail
	ws.Route(ws.GET("/project/{name}").
		To(api.HandleGetProjectDetail).
		Writes(swagger.Project{}))

	// delete project
	ws.Route(ws.DELETE("/project/{name}").
		To(api.HandleDeleteProject))

	/************** users *************/
	// create user
	ws.Route(ws.POST("/user").
		To(api.HandleCreateUser).
		Reads(model.RequestUserCreate{}).
		Writes(swagger.User{}))

	// get user list
	ws.Route(ws.GET("/user").
		To(api.HandleGetUserList).
		Reads(model.RequestUserList{}).
		Writes(model.ResponseUserList{}))

	// get user detail
	ws.Route(ws.GET("/user/{name}").
		To(api.HandleGetUserDetail).
		Writes(swagger.User{}))

	// delete user
	ws.Route(ws.DELETE("/user/{name}").
		To(api.HandleDeleteUser))

	// test
	ws.Route(ws.GET("/test").
		To(api.HandleTest))

	wsContainer.Add(ws)
	return wsContainer
}
