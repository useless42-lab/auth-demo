package controller

import (
	"UserCenter/models"
	"UserCenter/response"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	_ "github.com/joho/godotenv/autoload"
)

type AuthForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (form AuthForm) ValidateLoginForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required.Error("邮箱不能为空"), is.Email.Error("邮箱格式有误")),
		validation.Field(&form.Password, validation.Required.Error("密码不能为空")),
	)
}

func (form AuthForm) ValidateRegisterForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Username, validation.Required.Error("用户名不能为空")),
		validation.Field(&form.Email, validation.Required.Error("邮箱不能为空"), is.Email.Error("请输入邮箱")),
		validation.Field(&form.Password, validation.Required.Error("密码不能为空")),
	)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	loginForm := AuthForm{
		Email:    email,
		Password: password,
	}
	err := loginForm.ValidateLoginForm()
	if err != nil {
		response.Error(c, 4001, response.ConvertValidationErrorToString(err))
		return
	}
	result := models.CheckPasswordByEmail(email, password)
	if result.Id == "" {
		response.Error(c, 4002, "邮箱密码错误")
	} else {
		response.Success(c, 200, result)
	}
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	registerForm := AuthForm{
		Username: username,
		Email:    email,
		Password: password,
	}
	err := registerForm.ValidateRegisterForm()
	if err != nil {
		response.Error(c, 4003, response.ConvertValidationErrorToString(err))
		return
	}
	isExistEmail := models.CheckEmailIsExist(email)
	isExistUsername := models.CheckUsernameIsExist(username)
	if isExistEmail {
		response.Error(c, 4004, "邮箱已存在")
	} else if isExistUsername {
		response.Error(c, 4005, "用户名已存在")
	} else {
		user := models.AddUser(username, email, password)
		response.Success(c, 200, user)
	}
}
