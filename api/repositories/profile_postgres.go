package repositories

import (
	"database/sql"
	"dummyCVForm/models"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	db *sql.DB
}

func NewControllersRepositories(db *sql.DB) *Controllers {
	return &Controllers{db: db}
}

func (ct *Controllers) Get(c *gin.Context, id string) (*models.Profile, error) {
	var rowScan models.Profile

	err := ct.db.QueryRow("select profile_code, wanted_job_title, first_name, last_name, email, phone, country, city, address, postal_code, driving_license, nationality, place_of_birth, date_of_birth, photo_url from cv_form.user_dtls.profile_dtls where profile_code = $1", id).Scan(&rowScan.ProfileCode, &rowScan.WantedJobTitle, &rowScan.FirstName, &rowScan.LastName, &rowScan.Email, &rowScan.Phone, &rowScan.Country, &rowScan.City, &rowScan.Address, &rowScan.PostalCode, &rowScan.DrivingLicense, &rowScan.Nationality, &rowScan.PlaceOfBirth, &rowScan.DateOfBirth, &rowScan.PhotoUrl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &rowScan, nil
}

func (ct *Controllers) Create(c *gin.Context, req *models.Profile) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (ct *Controllers) Update(c *gin.Context, req *models.Profile) (int, error) {
	//TODO implement me
	panic("implement me")
}
