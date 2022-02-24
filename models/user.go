package models

import (
	"UserCenter/tools"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DefaultModel
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

type RUser struct {
	Id       string `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

/*
添加用户
*/
func AddUser(username string, email string, password string) RUser {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	id := tools.GenerateSnowflakeId()
	user := User{
		DefaultModel: DefaultModel{ID: id},
		Username:     username,
		Email:        email,
		Password:     string(encryptPassword),
	}
	rUser := RUser{
		Id:       strconv.FormatInt(id, 10),
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	DB.Table("user").Create(&user)
	return rUser
}

/*
检查邮箱是否存在
*/
func CheckEmailIsExist(email string) bool {
	var result RUser
	sqlStr := `select email from user where email=@email`
	DB.Raw(sqlStr, map[string]interface{}{
		"email": email,
	}).Scan(&result)
	if result.Email == "" {
		return false
	} else {
		return true
	}
}

/*
检查用户名是否存在
*/
func CheckUsernameIsExist(username string) bool {
	var result RUser
	sqlStr := `select username from user where username=@username`
	DB.Raw(sqlStr, map[string]interface{}{
		"username": username,
	}).Scan(&result)
	if result.Username == "" {
		return false
	} else {
		return true
	}
}

/*
校验邮箱密码
*/
func CheckPasswordByEmail(email string, password string) RUser {
	var user User
	var initUser RUser
	DB.Table("user").Where("email=?", email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return initUser
	} else {
		initUser = RUser{
			Id:       strconv.FormatInt(user.DefaultModel.ID, 10),
			Username: user.Username,
			Email:    user.Email,
			Avatar:   user.Avatar,
		}
	}
	return initUser
}

/*
根据用户编号检查密码是否正确
*/
func CheckPasswordByUserId(id string, password string) bool {
	var user User
	DB.Table("user").Where("id=?", id).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

/*
根据用户编号修改密码
*/
func ChangePasswordByUserId(id string, password string) {
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	DB.Exec("UPDATE `user` SET `password` = ? WHERE id = ?", string(encryptPassword), id)
}

/*
根据用户邮箱修改密码
*/
func ChangePasswordByEmail(email string, password string) {
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	DB.Exec("UPDATE `user` SET `password` = ? WHERE email = ?", string(encryptPassword), email)
}
