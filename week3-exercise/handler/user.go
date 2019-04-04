package handler

import (
	"../model"
	"../repo"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// Khai bao la mot interface
func UserSignin(c *gin.Context, repo repo.UserRepo) (*model.UserSigninResponse, error) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		return nil, err
	}
	password := []byte(user.Password)
	hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	createdUser, err := repo.Create(user)
	if err != nil {
		return nil, err
	}
	userSigninResponse := &model.UserSigninResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Fullname: createdUser.Fullname,
		Bod:      createdUser.Bod,
	}
	return userSigninResponse, nil
}

func UserLogin(c *gin.Context, repo repo.UserRepo) (*model.UserLoginReponse, error) {
	form := model.UserLoginForm{}
	if err := c.ShouldBind(&form); err != nil {
		return nil, err
	}
	user, err := repo.FindByUserLogin(form.Login)
	if err != nil {
		return nil, err
	}
	password := []byte(form.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	userLoginResponse := &model.UserLoginReponse{
		ID:       user.ID,
		Fullname: user.Fullname,
	}
	return userLoginResponse, err
}
