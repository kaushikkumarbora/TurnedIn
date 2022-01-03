package controller

import (
	"net/http"
	"strconv"

	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	db := storage.GetDBInstance()

	username := c.FormValue("username")

	row := db.QueryRow("SELECT username FROM \"user\" WHERE username=$1", username)

	var temp string

	if err := row.Scan(
		&temp,
	); err == nil {
		return c.NoContent(http.StatusBadRequest)
	}

	sem, err := strconv.Atoi(c.FormValue("semester"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Need Integer Semester"})
	}

	_, err = db.Exec("INSERT INTO \"user\" (first_name, last_name, dob, email_id, contact_no, skills, year_of_admission, year_of_completion, semester, bio, education, department, course, hostel, status, por_id, username, password, internship) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, (select por_id from por where name=$16), $17, $18, $19)", c.FormValue("first_name"), c.FormValue("last_name"), c.FormValue("dob"), c.FormValue("email_id"), c.FormValue("contact_no"), c.FormValue("skills"), c.FormValue("year_of_admission"), c.FormValue("year_of_completion"), sem, c.FormValue("bio"), c.FormValue("education"), c.FormValue("department"), c.FormValue("course"), c.FormValue("hostel"), true, c.FormValue("residence"), username, c.FormValue("password"), c.FormValue("internship"))

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	} else {
		return c.NoContent(http.StatusOK)
	}
}
