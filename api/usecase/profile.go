package usecase

import (
	"dummyCVForm/models"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	PRepositories models.ProfileRepositories
}

func NewControllersUsecase(PUsecase models.ProfileRepositories) *Controllers {
	return &Controllers{PRepositories: PUsecase}
}

func (ct *Controllers) Get(c *gin.Context, id string) (*models.Profile, error) {
	data, err := ct.PRepositories.Get(c, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ct *Controllers) Create(c *gin.Context, req *models.Profile) error {
	err := ct.PRepositories.Create(c, req)
	if err != nil {
		return err
	}
	return nil
}

func (ct *Controllers) Update(c *gin.Context, req *models.Profile) (int, error) {
	//TODO implement me
	panic("implement me")
}
