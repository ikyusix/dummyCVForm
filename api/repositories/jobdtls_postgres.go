package repositories

import (
	"database/sql"
	"dummyCVForm/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type JobControllers struct {
	db *sql.DB
}

func NewJobControllers(db *sql.DB) *JobControllers {
	return &JobControllers{db: db}
}

func (j *JobControllers) Get(c *gin.Context) (*models.DataArr, error) {

	var dataArr []models.Data
	var data models.Data
	var test models.DataArr
	id := c.Param("id")
	code, _ := strconv.Atoi(id)
	rows, err := j.db.Query("select id, job_title, employer, start_date, end_date, city, job_desc from cv_form.user_dtls.job_dtls where prof_id = $1 and del_flg != 'Y'", code)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&data.Id, &data.JobTitle, &data.Employer, &data.StartDate, &data.EndDate, &data.City, &data.Description)
		if err != nil {
			return nil, err
		}
		dataArr = append(dataArr, data)
	}

	test = models.DataArr{DataRow: dataArr}

	return &test, nil
}

func (j *JobControllers) Create(c *gin.Context, req *models.Data) error {
	_ = time.Now().Format("2006-01-02T15:04:05")
	id := c.Param("id")
	_, _ = strconv.Atoi(id)

	//TODO: Get by id, if nil create indicator ID from zero

	_, err := j.db.Exec("")
	if err != nil {
		return err
	}
	return nil
}
