package middlewares

import (
	"github.com/HeadcrabJ/go-gin-boilerplate/config"
	"github.com/HeadcrabJ/go-gin-boilerplate/db"
	"github.com/HeadcrabJ/go-gin-boilerplate/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var auth *jwt.GinJWTMiddleware

func InitJWT() {
	c := config.GetConfig()
	var authErr error
	auth, authErr = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "realm",
		Key:         []byte(c.GetString("auth.secret_key")),
		Timeout:     time.Hour * 24 * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserRes); ok {
				return jwt.MapClaims{
					"id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.UserRes{
				ID: uint(claims["id"].(float64)),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			userTel := loginVals.Tel
			userPass := loginVals.Password

			var user models.User
			db.GetDB().First(&user, "tel = ?", userTel)

			if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPass)) == nil {
				return &models.UserRes{
					ID:         user.ID,
					Tel:        user.Tel,
					Email:      user.Email,
					FirstName:  user.FirstName,
					LastName:   user.LastName,
					Patronymic: user.Patronymic,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.UserRes); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if authErr != nil {
		return
	}

	authErrInit := auth.MiddlewareInit()
	if authErrInit != nil {
		return
	}
}

func GetJWT() *jwt.GinJWTMiddleware {
	return auth
}
