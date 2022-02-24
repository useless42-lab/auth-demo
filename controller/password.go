package controller

import (
	"UserCenter/models"
	"UserCenter/response"
	"UserCenter/tools"
	"os"

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	userId := c.PostForm("user_id")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")
	reNewPassword := c.PostForm("re_new_password")
	if reNewPassword != newPassword {
		response.Error(c, 4006, "新密码不一致")
		return
	} else {
		isRight := models.CheckPasswordByUserId(userId, oldPassword)
		if isRight {
			encryptCode := c.PostForm("encrypt_code")
			decryptCode := tools.AesDecrypt(encryptCode, os.Getenv("KEY"))
			if decryptCode == "changepassword"+userId {
				models.ChangePasswordByUserId(userId, newPassword)
				response.Success(c, 200, "")
			} else {
				response.Error(c, 5000, "密钥有误")
			}
		} else {
			response.Error(c, 4007, "密码有误")
		}
	}
}

func ResetPassword(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	encryptCode := c.PostForm("encrypt_code")
	decryptCode := tools.AesDecrypt(encryptCode, os.Getenv("KEY"))
	if decryptCode == "resetpassword"+email {
		models.ChangePasswordByEmail(email, password)
		response.Success(c, 200, "")
	} else {
		response.Error(c, 5000, "密钥有误")
	}
}
