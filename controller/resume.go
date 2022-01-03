package controller

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

func GetResume(c echo.Context) error {
	db := storage.GetDBInstance()

	user_id, _ := GetDetails(c)

	row := db.QueryRow("SELECT resume FROM \"user\" WHERE user_id=$1 AND status=TRUE", user_id)

	var resume []byte

	if err := row.Scan(
		&resume,
	); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	http.ServeContent(c.Response().Writer, c.Request(), "resume.pdf", time.Now(), bytes.NewReader(resume))
	return nil
}

func PutResume(c echo.Context) error {
	db := storage.GetDBInstance()

	user_id, _ := GetDetails(c)

	resume, _, err := c.Request().FormFile("resume")
	if err != nil {
		return c.NoContent(http.StatusRequestTimeout)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, resume); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	row := db.QueryRow("SELECT status FROM \"user\" WHERE user_id=$1", user_id)

	var status bool

	if err := row.Scan(
		&status,
	); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if !status {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err = db.Exec("UPDATE \"user\" SET resume=$1 WHERE user_id = $2", buf.Bytes(), user_id)

	return err
}
