package controller

import (
	"net/http"

	"github.com/kaushikkumarbora/TurnedIn/model"
	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

// GetUsers
func GetUsers(c echo.Context) error {
	User, _ := GetRepoUsers(c)
	return c.JSON(http.StatusOK, User)
}

func GetRepoUsers(c echo.Context) ([]model.User, error) {
	db := storage.GetDBInstance()
	Users := []model.User{}

	queryString := "SELECT * FROM users WHERE "

	rows, err := db.Query(queryString)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp model.User
		if err := rows.Scan(
			&tmp.Id,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.DOB,
			&tmp.Education,
			&tmp.Email,
			&tmp.Contact,
			&tmp.Skills,
			&tmp.Internship,
			&tmp.YearOfJoining,
			&tmp.YearOfCompletion,
			&tmp.Department,
			&tmp.Course,
			&tmp.Semester,
			&tmp.Bio,
			&tmp.Residence,
			&tmp.Hostel); err != nil {
			return nil, err
		}
		Users = append(Users, tmp)
	}

	return Users, nil
}
