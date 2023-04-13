package repositories

import (
	"database/sql"
	"dummyCVForm/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ProfileControllers struct {
	db *sql.DB
}

func NewProfileControllers(db *sql.DB) *ProfileControllers {
	return &ProfileControllers{db: db}
}

func (ct *ProfileControllers) Get(c *gin.Context, id string) (*models.Profile, error) {
	var rowScan models.Profile

	err := ct.db.QueryRow("select profile_code, wanted_job_title, first_name, last_name, email, phone, country, city, address, postal_code, driving_license, nationality, place_of_birth, date_of_birth, photo_url from cv_form.user_dtls.profile_dtls where profile_code = $1 and del_flg != 'Y'", id).Scan(&rowScan.ProfileCode, &rowScan.WantedJobTitle, &rowScan.FirstName, &rowScan.LastName, &rowScan.Email, &rowScan.Phone, &rowScan.Country, &rowScan.City, &rowScan.Address, &rowScan.PostalCode, &rowScan.DrivingLicense, &rowScan.Nationality, &rowScan.PlaceOfBirth, &rowScan.DateOfBirth, &rowScan.PhotoUrl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &rowScan, nil
}

func (ct *ProfileControllers) Create(c *gin.Context, req *models.Profile) error {
	timeStamp := time.Now().Format("2006-01-02T15:04:05")
	pcode := fmt.Sprintf("%v", req.ProfileCode)
	_, err := ct.db.Exec("insert into cv_form.user_dtls.profile_dtls (profile_code, wanted_job_title, first_name, last_name, email, phone, country, city, address, postal_code, driving_license, nationality, place_of_birth, date_of_birth, photo_url, del_flg, mod_id, mod_time, cre_id, cre_time) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, 'N', $16, $17, $18, $19)", pcode, req.WantedJobTitle, req.FirstName, req.LastName, req.Email, req.Phone, req.Country, req.City, req.Address, req.PostalCode, req.DrivingLicense, req.Nationality, req.PlaceOfBirth, req.DateOfBirth, req.PhotoUrl, pcode, timeStamp, pcode, timeStamp)
	if err != nil {
		return err
	}
	return nil
}

func (ct *ProfileControllers) Update(c *gin.Context, req *models.Profile) error {
	timeStamp := time.Now().Format("2006-01-02T15:04:05")
	id := c.Param("id")
	code, _ := strconv.Atoi(id)
	_, err := ct.db.Exec("update cv_form.user_dtls.profile_dtls set wanted_job_title = $1, first_name = $2, last_name = $3, email = $4, phone = $5, country = $6, city = $7, address = $8, postal_code = $9, driving_license = $10, nationality = $11, place_of_birth = $12, date_of_birth = $13, mod_id = $14, mod_time = $15 where profile_code = $16 and del_flg = 'N'", req.WantedJobTitle, req.FirstName, req.LastName, req.Email, req.Phone, req.Country, req.City, req.Address, req.PostalCode, req.DrivingLicense, req.Nationality, req.PlaceOfBirth, req.DateOfBirth, id, timeStamp, code)
	if err != nil {
		return err
	}
	return nil
}
