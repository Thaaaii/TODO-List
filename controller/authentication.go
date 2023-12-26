package controller

import (
	"github.com/Thaaaii/TODO-List/models"
	"github.com/Thaaaii/TODO-List/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Register is a handler to process user signups. It processes POST requests, extracts the given password, hashes it
// and inserts the new user with the hashed password into the database
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

// Login is a handler that processes the user login. It checks the given data and compares it with the hashed password
// in the database to confirm the identity. The handler also generates a JWT Token and sets the token as cookie.
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

// AuthenticationMiddleware is a handler that is used before we access protected routes. It checks whether the user has
// permission to gain further access to user specific data and routes by checking the given claims.
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
