package controller

import (
	"net/http"

	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	err := deleteRepoUser(c)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusOK)
}

func deleteRepoUser(c echo.Context) error {
	db := storage.GetDBInstance()

	user_id, _ := GetDetails(c)

	_, err := db.Exec("UPDATE \"user\"	SET status=$1 WHERE user_id = $2 AND status=TRUE", false, user_id)

	return err
}
