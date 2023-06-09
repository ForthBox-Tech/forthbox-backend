package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Resp keeps API responses aligned across handlers.
type Resp struct {
	ginc *gin.Context
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewResp(c *gin.Context) *Resp {
	return &Resp{ginc: c}
}

// placeholder
func (resp *Resp) Output() {
	var statusCode int
	if resp.Code == http.StatusBadRequest {
		statusCode = http.StatusOK
	} else {
		statusCode = resp.Code
	}
	resp.ginc.JSON(statusCode, resp) 
}

// placeholder
func (resp *Resp) Out(data interface{}) {
	resp.Data = data
	resp.Code = http.StatusOK
	if len(resp.Msg) == 0 {
		resp.Msg = "OK" 
	}
	resp.Output()
}

// placeholder
func (resp *Resp) Err(msg string) {
	if resp.Code == 0 {
		resp.Code = http.StatusBadRequest
	}
	resp.Msg = msg
	resp.Data = map[string]string{}
	resp.Output()
}


