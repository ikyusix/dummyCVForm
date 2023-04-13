package usecase

import (
	"dummyCVForm/models"
	"github.com/gin-gonic/gin"
)

type JobControllers struct {
	Repositories models.JobDtlsRepositories
}

func NewJobControllers(repositories models.JobDtlsRepositories) *JobControllers {
	return &JobControllers{Repositories: repositories}
}

func (j *JobControllers) Get(c *gin.Context) (*models.DataArr, error) {
	data, err := j.Repositories.Get(c)
	if err != nil {
		return nil, err
	}
	return data, nil
}
