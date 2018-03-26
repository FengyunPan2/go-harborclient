package model

import (
	"hubimage/swagger"
)

// swagger.ProjectReq
type RequestProjectCreate struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

/**
 * List projects
 * This endpoint returns all projects created by Harbor, and can be filtered by project name.
 *
 * @param name The name of project.
 * @param public The project is public or private.
 * @param owner The name of project owner.
 * @param page The page nubmer, default is 1.
 * @param pageSize The size of per page, default is 10, maximum is 100.
 * @return []Project
 */
type RequestProjectList struct {
	Name     string `json:"name"`
	All      bool   `json:"all"`
	Public   bool   `json:"public"`
	Owner    string `json:"owner"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
}

type ResponseProjectList struct {
	ProjectList []swagger.Project
}
