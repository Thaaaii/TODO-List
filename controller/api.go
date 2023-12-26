package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Thaaaii/TODO-List/models"
	"github.com/Thaaaii/TODO-List/utils"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var newUser models.User
	var err error

	if err = ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Password, err = models.HashPassword(newUser.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = models.InsertUserIntoTable(newUser.Name, newUser.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task could not be inserted",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func Login(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := models.LoginCheck(user.Name, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password is incorrect",
		})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(60 * time.Minute),
	})
}

func Logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
	})
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("jwt")
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		token, err := utils.TokenValid(tokenString)

		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		var username string

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username = claims["username"].(string)
		} else {
			ctx.String(http.StatusInternalServerError, "Claims could not be extracted")
			ctx.Abort()
			return
		}

		if ctx.Param("user") != username {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func GetTasks(ctx *gin.Context) {
	var userTasks []models.Task
	user := ctx.Param("user")
	userID, err := models.SelectUserID(user)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "User could not be found",
		})
		return
	}

	userTasks, err = models.SelectUserTasks(userID)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Tasks of the user could not be found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, userTasks)
}

func PostTasks(ctx *gin.Context) {
	user := ctx.Param("user")
	var newTask models.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := models.SelectUserID(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User could not be found",
		})
		return
	}

	taskID, err := models.InsertTaskIntoTable(newTask.Title, newTask.Description, newTask.IsDone, userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task could not be inserted",
		})
		return
	}

	newTask.ID = int(taskID)
	err = models.InsertCategoriesIntoTable(newTask.Categories, taskID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Categories of the task could not be inserted",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newTask)
}

func PatchTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedTask models.Task

	if err = ctx.BindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.UpdateUserTask(taskID, updatedTask.Title, updatedTask.Description, updatedTask.IsDone)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Task to update could not be found",
		})
		return
	}

	err = models.UpdateTaskCategories(taskID, updatedTask.Categories)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Task to update categories could not be found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource with ID " + ctx.Param("taskID") + " updated.",
	})
}

func PatchTaskOrder(ctx *gin.Context) {
	type Sequence struct {
		SequenceNumber int `json:"sequenceNumber"`
	}

	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var sequence Sequence

	if err = ctx.BindJSON(&sequence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.UpdateUserTaskOrder(taskID, int64(sequence.SequenceNumber))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Task to update could not be found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource with ID " + ctx.Param("taskID") + " updated.",
	})
}

func DeleteTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.DeleteUserTask(taskID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task could not be deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource deleted",
		"id":      taskID,
	})
}
