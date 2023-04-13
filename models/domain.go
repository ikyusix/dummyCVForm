package models

import "github.com/gin-gonic/gin"

type ProfileUsecase interface {
	Get(c *gin.Context, id string) (*Profile, error)
	Create(c *gin.Context, req *Profile) error
	Update(c *gin.Context, req *Profile) error
}

type ProfileRepositories interface {
	Get(c *gin.Context, id string) (*Profile, error)
	Create(c *gin.Context, req *Profile) error
	Update(c *gin.Context, req *Profile) error
}

type JobDtlsUsecase interface {
	Get(c *gin.Context) (*DataArr, error)
}

type JobDtlsRepositories interface {
	Get(c *gin.Context) (*DataArr, error)
}
