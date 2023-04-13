package models

import "github.com/gin-gonic/gin"

type Response struct {
	RespCode    string      `json:"responseCode,omitempty"`
	RespMessage string      `json:"respMessage,omitempty"`
	IdMessage   string      `json:"idMessage,omitempty"`
	Response    interface{} `json:"response,omitempty"`
}

type DataArr struct {
	DataRow []Data `json:"data,omitempty"`
}

type Data struct {
	Id          int    `json:"id,omitempty"`
	JobTitle    string `json:"jobTitle,omitempty"`
	Employer    string `json:"employer,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
	City        string `json:"city,omitempty"`
	Description string `json:"description,omitempty"`
}

type Profile struct {
	ProfileCode    int    `json:"profileCode,omitempty"`
	WantedJobTitle string `json:"wantedJobTitle,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Country        string `json:"country,omitempty"`
	City           string `json:"city,omitempty"`
	Address        string `json:"address,omitempty"`
	PostalCode     int    `json:"postalCode,omitempty"`
	DrivingLicense string `json:"drivingLicense,omitempty"`
	Nationality    string `json:"nationality,omitempty"`
	PlaceOfBirth   string `json:"placeOfBirth,omitempty"`
	DateOfBirth    string `json:"dateOfBirth,omitempty"`
	PhotoUrl       string `json:"photoUrl,omitempty"`
}

func CreateResponse(c *gin.Context, respCode, respMessage, idMessage string, response interface{}) Response {
	c.Set("RespCode", respCode)
	return Response{RespCode: respCode, RespMessage: respMessage, IdMessage: idMessage, Response: response}
}
