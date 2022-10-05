package controller

import (
	"crud-user-mvc/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateNewUser(ctx *gin.Context) {
	var userInput model.UserBodyParam
	if err := c.BindBody(ctx, &userInput); err != nil {
		ctx.JSON(400, c.helper.ResponseFormat("Bad request", false, err.Error()))
		return
	}

	hashedPassword, err := c.helper.HashPassword(userInput.CurrentPassword)
	if err != nil {
		ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
		return
	}

	user := model.User{
		Username: userInput.Username,
		Password: hashedPassword,
	}

	if err := c.db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			ctx.JSON(400, c.helper.ResponseFormat("Username already exists", false, err.Error()))
		} else {
			ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
		}
		return
	}

	ctx.JSON(200, c.helper.ResponseFormat("User created successfully", true, user))
}

func (c *Controller) GetUserList(ctx *gin.Context) {
	var userParam model.UserParam
	if err := c.BindParam(ctx, &userParam); err != nil {
		ctx.JSON(400, c.helper.ResponseFormat("Bad request", false, err.Error()))
		return
	}

	if userParam.Limit == 0 {
		userParam.Limit = 10
	}

	var users []model.User
	if err := c.db.Limit(userParam.Limit).Offset((userParam.Page - 1) * userParam.Limit).Find(&users).Error; err != nil {
		ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
		return
	}

	if len(users) == 0 {
		ctx.JSON(404, c.helper.ResponseFormat("Users not found", false, nil))
		return
	}

	ctx.JSON(200, c.helper.ResponseFormat("Successfully get user's list", true, users))
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	var userParam model.UserParam
	if err := c.BindParam(ctx, &userParam); err != nil {
		ctx.JSON(400, c.helper.ResponseFormat("Bad request", false, err.Error()))
		return
	}

	var userBody model.UserBodyParam
	if err := c.BindBody(ctx, &userBody); err != nil {
		ctx.JSON(400, c.helper.ResponseFormat("Bad request", false, err.Error()))
		return
	}

	var user model.User
	if err := c.db.First(&user, userParam.UserID).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			ctx.JSON(404, c.helper.ResponseFormat("User not found", false, err.Error()))
		} else {
			ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
		}
		return
	}

	if !c.helper.ComparePassword(user.Password, userBody.CurrentPassword) {
		ctx.JSON(400, c.helper.ResponseFormat("Wrong password", false, nil))
		return
	}

	if userBody.NewPassword == "" {
		hashedPassword, err := c.helper.HashPassword(userBody.NewPassword)
		if err != nil {
			ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
			return
		}
		user.Password = hashedPassword
	}
	user.Username = userBody.Username

	if err := c.db.Save(&user).Error; err != nil {
		ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
		return
	}

	ctx.JSON(200, c.helper.ResponseFormat("User updated successfully", true, user))
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	var userParam model.UserParam
	if err := c.BindParam(ctx, &userParam); err != nil {
		ctx.JSON(400, c.helper.ResponseFormat("Bad request", false, err.Error()))
		return
	}

	queryTx := c.db.Delete(&model.User{}, userParam.UserID)

	if err := queryTx.Error; err != nil {
		ctx.JSON(500, c.helper.ResponseFormat("Internal server error", false, err.Error()))
		return
	}

	if queryTx.RowsAffected == 0 {
		ctx.JSON(404, c.helper.ResponseFormat("User not found", false, nil))
		return
	}

	ctx.JSON(200, c.helper.ResponseFormat("User deleted successfully", true, nil))
}
