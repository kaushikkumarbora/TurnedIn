package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kaushikkumarbora/TurnedIn/model"
	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if user_id, auth, _ := verifyLogin(username, password); auth {

		// Set custom claims
		claims := &model.Auth{
			Id:       user_id,
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// generate encoded token and send it as response
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return c.NoContent(http.StatusUnauthorized)
}

func verifyLogin(username string, password string) (int64, bool, error) {
	db := storage.GetDBInstance()

	row := db.QueryRow("SELECT user_id, password FROM \"user\" WHERE username=$1 AND status=TRUE", username)

	var user_id int64
	var password_db string

	if err := row.Scan(
		&user_id,
		&password_db,
	); err != nil {
		return 0, false, err
	}

	if password != password_db {
		return user_id, false, nil
	}

	return user_id, true, nil
}

func GetDetails(c echo.Context) (int64, string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Auth)
	return claims.Id, claims.Username
}
