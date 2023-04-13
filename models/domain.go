package models

import "github.com/gin-gonic/gin"

type ProfileUsecase interface {
	Get(c *gin.Context, id string) (*Profile, error)
	Create(c *gin.Context, req *Profile) (int, error)
	Update(c *gin.Context, req *Profile) (int, error)
}

type ProfileRepositories interface {
	Get(c *gin.Context, id string) (*Profile, error)
	Create(c *gin.Context, req *Profile) (int, error)
	Update(c *gin.Context, req *Profile) (int, error)
}
