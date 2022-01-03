package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	db := storage.GetDBInstance()

	user_id, _ := GetDetails(c)

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

	sem, err := strconv.Atoi(c.FormValue("semester"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Need Integer Semester"})
	}

	temp := c.FormValue("year_of_completion")
	fmt.Print(temp)

	_, err = db.Exec("UPDATE \"user\" SET first_name=$1, last_name=$2, dob=$3, email_id=$4, contact_no=$5, skills=$6, year_of_admission=$7, year_of_completion=$8, semester=$9, bio=$10, education=$11, department=$12, course=$13, hostel=$14, por_id=(select por_id from por where name=$15), password=$16 WHERE \"user\".user_id=$17", c.FormValue("first_name"), c.FormValue("last_name"), c.FormValue("dob"), c.FormValue("email_id"), c.FormValue("contact_no"), c.FormValue("skills"), c.FormValue("year_of_admission"), c.FormValue("year_of_completion"), sem, c.FormValue("bio"), c.FormValue("education"), c.FormValue("department"), c.FormValue("course"), c.FormValue("hostel"), c.FormValue("residence"), c.FormValue("password"), user_id)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	} else {
		return c.NoContent(http.StatusOK)
	}
}
