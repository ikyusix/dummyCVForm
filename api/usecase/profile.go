package usecase

import (
	"dummyCVForm/models"
	"github.com/gin-gonic/gin"
)

type ProfileControllers struct {
	Repositories models.ProfileRepositories
}

func NewProfileControllers(PUsecase models.ProfileRepositories) *ProfileControllers {
	return &ProfileControllers{Repositories: PUsecase}
}

func (ct *ProfileControllers) Get(c *gin.Context, id string) (*models.Profile, error) {
	data, err := ct.Repositories.Get(c, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ct *ProfileControllers) Create(c *gin.Context, req *models.Profile) error {
	err := ct.Repositories.Create(c, req)
	if err != nil {
		return err
	}
	return nil
}

func (ct *ProfileControllers) Update(c *gin.Context, req *models.Profile) error {
	err := ct.Repositories.Update(c, req)
	if err != nil {
		return err
	}
	return nil
}
