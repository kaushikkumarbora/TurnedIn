package controller

import (
	"net/http"
	"strconv"

	"github.com/kaushikkumarbora/TurnedIn/model"
	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

func GetRequests(c echo.Context) error {
	db := storage.GetDBInstance()

	Requests := make([]model.Connection, 0)

	user_id, _ := GetDetails(c)
	var status bool

	row := db.QueryRow("SELECT status FROM \"user\" WHERE user_id=$1 AND status=TRUE", user_id)
	if err := row.Scan(
		&status,
	); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	rows, err := db.Query("SELECT l_user_id, connection.status FROM connection LEFT JOIN \"user\" as u ON l_user_id=u.user_id WHERE r_user_id=$1 AND connection.status='p'", user_id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	defer rows.Close()

	for rows.Next() {
		var tmp model.Connection
		tmp.RUserID = user_id
		if err := rows.Scan(
			&tmp.LUserID,
			&tmp.Status,
		); err != nil {
			return err
		}
		Requests = append(Requests, tmp)
	}

	return c.JSON(http.StatusOK, Requests)
}

func SendConnection(c echo.Context) error {
	success, err := repoConnection(c, true)
	if err != nil || !success {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusOK)
}

func AcceptConnection(c echo.Context) error {
	success, err := repoConnection(c, false)
	if err != nil || !success {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusOK)
}

func repoConnection(c echo.Context, newconnec bool) (bool, error) {
	db := storage.GetDBInstance()

	user_id, _ := GetDetails(c)

	from, _ := strconv.Atoi(c.QueryParam("from"))
	to, _ := strconv.Atoi(c.QueryParam("to"))
	var status bool

	row := db.QueryRow("SELECT status FROM \"user\" WHERE user_id=$1 AND status=TRUE", from)
	if err := row.Scan(
		&status,
	); err != nil {
		return false, err
	}

	row = db.QueryRow("SELECT status FROM \"user\" WHERE user_id=$1 AND status=TRUE", to)
	if err := row.Scan(
		&status,
	); err != nil {
		return false, err
	}

	var err error

	if newconnec {
		if user_id != int64(from) {
			return false, nil
		}
		_, err = db.Exec("INSERT INTO connection (l_user_id, r_user_id, status) VALUES ($1, $2, 'p')", from, to)
	} else {
		row = db.QueryRow("SELECT * FROM connection WHERE l_user_id=$1 AND r_user_id=$2 AND status='p'", from, to)

		tmp := model.Connection{}
		if err := row.Scan(
			&tmp.LUserID,
			&tmp.RUserID,
			&tmp.Status,
		); err != nil {
			return false, err
		}

		if user_id != int64(to) {
			return false, nil
		}
		_, err = db.Exec("UPDATE connection SET status = 'a' WHERE l_user_id = $1 AND r_user_id = $2", from, to)
	}

	return true, err
}
