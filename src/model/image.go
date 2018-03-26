package model

import (
	"io"

	"hubimage/swagger"
)

type RequestImageCreate struct {
	Reader      io.Reader `json:"reader"`
	File        string    `json:"file"`
	ProjectName string    `json:"projectName"`
	Name        string    `json:"name"`
	Tag         string    `json:"tag"`
	EnableWait  bool      `json:"enableWait"`
}

/**
 * Get repositories accompany with relevant project and repo name.
 * This endpoint let user search repositories accompanying with relevant project ID and repo name.
 *
 * @param projectId Relevant project ID.
 * @param page The page nubmer, default is 1.
 * @param pageSize The size of per page, default is 10, maximum is 100.
 * @return []Repository
 */
type RequestImageList struct {
	ProjectId   int32  `json:"projectId"`
	ProjectName string `json:"projectName"`
	Page        int32  `json:"page"`
	PageSize    int32  `json:"pageSize"`
}

type ResponseImageList struct {
	ImageList []swagger.Repository
}

type RequestImageDetail struct {
	ProjectName string `json:"projectName"`
	ImageName   string `json:"imageName"`
	Tag         string `json:"tag"`
}

type ImageDetail struct {
	Image    swagger.DetailedTag `json:"image"`
	Manifest swagger.Manifest    `json:"manifest"`
}

type ResponseImageDetail struct {
	ImageTags []ImageDetail
}

type RequestImageDelete struct {
	ProjectName string `json:"projectName"`
	ImageName   string `json:"imageName"`
	Tag         string `json:"tag"`
}
