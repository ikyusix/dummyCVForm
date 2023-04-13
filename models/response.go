package models

import "github.com/gin-gonic/gin"

type Response struct {
	RespCode    string      `json:"responseCode,omitempty"`
	RespMessage string      `json:"respMessage,omitempty"`
	IdMessage   string      `json:"idMessage,omitempty"`
	Response    interface{} `json:"response,omitempty"`
}

func CreateResponse(c *gin.Context, respCode, respMessage, idMessage string, response interface{}) Response {
	c.Set("RespCode", respCode)
	return Response{RespCode: respCode, RespMessage: respMessage, IdMessage: idMessage, Response: response}
}
