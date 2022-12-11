package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"log"
	"main/internal/app/ds"
	"main/internal/app/role"
	"main/swagger"
	"net/http"
	"strings"
)

func (a *Application) GetRoleByToken(jwtStr string) (role role.Role) {
	if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
		return // завершаем обработку
	}
	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.JWT.Token), nil
	})
	if err != nil {
		log.Println(err)

		return
	}

	myClaims := token.Claims.(*ds.JWTClaims)
	log.Println(myClaims)

	return myClaims.Role
}

func (a *Application) Role(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	roleByToken := a.GetRoleByToken(jwtStr)

	gCtx.JSON(http.StatusOK, roleByToken)
}

func (a *Application) GetUser(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	userName, err := a.repo.GetUserByUUID(UUID)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&swagger.Error{
				Description: "Can't get a user",
				Error:       swagger.Err500,
				Type:        swagger.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, userName)

}
