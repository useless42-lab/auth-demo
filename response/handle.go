package response

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type PaginationData struct {
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	PageSize  int         `json:"page_size"`
	PageIndex int         `json:"page_index"`
}

func Error(c *gin.Context, code int, errors interface{}) {
	c.JSON(200, gin.H{
		"code":  code,
		"error": errors,
	})
	return
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
	})
	return
}

func ConvertValidationErrorToString(data interface{}) string {
	for _, val := range data.(validation.Errors) {
		return val.Error()
	}
	return ""
}
