package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func get_user_data(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	User, err := services.GetUserData(claims)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, user not found"})
		return
	}

	response := services.NewProfileDataResponse(User)

	ctx.JSON(http.StatusOK, response)
}

func add_points(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user_id, err := strconv.Atoi(ctx.PostForm("user_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	points, err := strconv.Atoi(ctx.PostForm("points"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid points"})
		return
	}

	err = services.AddPoints(claims.User_id, user_id, points)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func get_all_users(ctx *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, services.NewAllUsersResponse(users))
}

func get_all_users_with_points(ctx *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, services.NewAllUsersWithPointsResponse(users))
}

func get_all_users_notifications(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	users, err := services.RetriveAllUserNotification(claims.User_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func get_users_points_and_stars_stats(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	users, err := services.GetUser(claims.User_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	gold_stars_number, err := services.CountStarsType(claims.User_id, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	silver_stars_number, err := services.CountStarsType(claims.User_id, 1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	bronze_stars_number, err := services.CountStarsType(claims.User_id, 2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"points": users.Points, "gold_stars": gold_stars_number, "silver_stars": silver_stars_number, "bronze_stars": bronze_stars_number})
}

func modify_profile_data(ctx *gin.Context) {

	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := services.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	temp_user := &models.User{}
	temp_user.ID = claims.User_id

	temp_user.Personal_description = ctx.PostForm("description")
	temp_user.Snapchat_id = ctx.PostForm("snapchat")
	temp_user.Instagram_id = ctx.PostForm("instagram")
	temp_user.Facebook_id = ctx.PostForm("facebook")
	temp_user.Google_id = ctx.PostForm("google")
	temp_user.Hometown = ctx.PostForm("hometown")
	temp_user.Studies = ctx.PostForm("studies")

	err = services.ModifyProfileData(temp_user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func Register_profile_routes(rg *gin.RouterGroup) {
	router_group := rg.Group("/profile")
	router_group.GET("/user/data", get_user_data)
	router_group.POST("/points/add", add_points)
	router_group.GET("/user/all", get_all_users)
	router_group.GET("/user/all/points", get_all_users_with_points)
	router_group.GET("/user/notifications", get_all_users_notifications)
	router_group.GET("/user/stats", get_users_points_and_stars_stats)
	router_group.POST("/user/modify", modify_profile_data)
}
