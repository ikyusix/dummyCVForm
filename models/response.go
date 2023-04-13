package models

import "github.com/gin-gonic/gin"

type Response struct {
	RespCode    string      `json:"responseCode,omitempty"`
	RespMessage string      `json:"respMessage,omitempty"`
	IdMessage   string      `json:"idMessage,omitempty"`
	Response    interface{} `json:"response,omitempty"`
}

type Profile struct {
	ProfileCode    int    `json:"profileCode"`
	WantedJobTitle string `json:"wantedJobTitle"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Address        string `json:"address"`
	PostalCode     int    `json:"postalCode"`
	DrivingLicense string `json:"drivingLicense"`
	Nationality    string `json:"nationality"`
	PlaceOfBirth   string `json:"placeOfBirth"`
	DateOfBirth    string `json:"dateOfBirth"`
	PhotoUrl       string `json:"photoUrl"`
}

func CreateResponse(c *gin.Context, respCode, respMessage, idMessage string, response interface{}) Response {
	c.Set("RespCode", respCode)
	return Response{RespCode: respCode, RespMessage: respMessage, IdMessage: idMessage, Response: response}
}
