package model

import (
	"hubimage/swagger"
)

// swagger.ProjectReq
type RequestUserCreate struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Realname     string `json:"realname"`
	Comment      string `json:"comment"`
	RoleName     string `json:"role_name"`
	HasAdminRole int32  `json:"has_admin_role"`
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
type RequestUserList struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
}

type ResponseUserList struct {
	UserList []swagger.User
}
