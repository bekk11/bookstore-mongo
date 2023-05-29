package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mongo_go/pkg/common/models"
	"net/http"
	"time"
)

//---------------- LIST USERS -----------------------------

func (h handler) ListUsers(ctx *gin.Context) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		ctx.AbortWithError(http.StatusNoContent, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &users)
}

//=========================================================

//---------------- CREATE USER ----------------------------

type CreateUserRequestBody struct {
	Firstname string `json:"firstname" form:"firstname"`
	Lastname  string `json:"lastname" form:"lastname"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
}

func (h handler) CreateUser(ctx *gin.Context) {
	body := CreateUserRequestBody{}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Firstname = body.Firstname
	user.Lastname = body.Lastname
	user.Username = body.Username
	user.Password = body.Password

	if result := h.DB.Create(&user); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &user)
}

//=========================================================

//---------------- CREATE USER ----------------------------

func (h handler) RetrieveUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &user)
}

// =========================================================

//---------------- LOGIN USER ----------------------------

func createToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte("my-secret-key"))
}

// =========================================================
