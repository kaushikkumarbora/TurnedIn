package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/kaushikkumarbora/TurnedIn/model"
	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
)

//GetName
func GetName(c echo.Context) error {
	Names, err := GetRepoNames(c)
	if err != nil {
		c.Error(err)
	}
	return c.JSON(http.StatusOK, Names)
}

// GetUsers
func GetUsers(c echo.Context) error {
	Users, err := GetRepoUsers(c)
	if err != nil {
		c.Error(err)
	}
	return c.JSON(http.StatusOK, Users)
}

//GetUser
func GetUser(c echo.Context) error {
	db := storage.GetDBInstance()

	user_id, _ := GetDetails(c)
	User := model.User{}

	row := db.QueryRow("SELECT user_id, first_name, last_name, to_char(\"user\".dob, 'MM-DD-YYYY'), email_id, contact_no, skills, year_of_admission, year_of_completion, semester, bio, por.name, education, department, course, hostel, internship FROM \"user\" LEFT JOIN por ON \"user\".por_id=por.por_id WHERE user_id=$1 AND status=TRUE", user_id)

	if err := row.Scan(
		&User.Id,
		&User.FirstName,
		&User.LastName,
		&User.DOB,
		&User.Email,
		&User.Contact,
		&User.Skills,
		&User.YearOfJoining,
		&User.YearOfCompletion,
		&User.Semester,
		&User.Bio,
		&User.Residence,
		&User.Education,
		&User.Department,
		&User.Course,
		&User.Hostel,
		&User.Internship,
	); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, User)
}

func GetRepoUsers(c echo.Context) ([]model.User, error) {
	db := storage.GetDBInstance()
	Users := make([]model.User, 0)

	sem, err := strconv.Atoi(c.QueryParam("semseter"))
	semPresent := true

	if err != nil {
		semPresent = false
	}

	var rows *sql.Rows

	if semPresent {
		rows, err = db.Query("SELECT user_id, first_name, last_name, to_char(\"user\".dob, 'MM-DD-YYYY'), email_id, contact_no, skills, year_of_admission, year_of_completion, semester, bio, por.name, education, department, course, hostel, internship FROM \"user\" LEFT JOIN por ON \"user\".por_id=por.por_id WHERE first_name LIKE $1 AND last_name LIKE $2 AND semester=$3 AND department LIKE $4 AND course LIKE $5 AND hostel LIKE $6 AND skills LIKE $7 AND status=TRUE", string('%')+c.QueryParam("first_name")+string('%'), string('%')+c.QueryParam("last_name")+string('%'), sem, string('%')+c.QueryParam("department")+string('%'), string('%')+c.QueryParam("course")+string('%'), string('%')+c.QueryParam("hostel")+string('%'), string('%')+c.QueryParam("skills")+string('%'))
	} else {
		rows, err = db.Query("SELECT user_id, first_name, last_name, to_char(\"user\".dob, 'MM-DD-YYYY'), email_id, contact_no, skills, year_of_admission, year_of_completion, semester, bio, por.name, education, department, course, hostel, internship FROM \"user\" LEFT JOIN por ON \"user\".por_id=por.por_id WHERE first_name LIKE $1 AND last_name LIKE $2 AND department LIKE $3 AND course LIKE $4 AND hostel LIKE $5 AND skills LIKE $6 AND status=TRUE", string('%')+c.QueryParam("first_name")+string('%'), string('%')+c.QueryParam("last_name")+string('%'), string('%')+c.QueryParam("department")+string('%'), string('%')+c.QueryParam("course")+string('%'), string('%')+c.QueryParam("hostel")+string('%'), string('%')+c.QueryParam("skills")+string('%'))
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp model.User
		if err := rows.Scan(
			&tmp.Id,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.DOB,
			&tmp.Email,
			&tmp.Contact,
			&tmp.Skills,
			&tmp.YearOfJoining,
			&tmp.YearOfCompletion,
			&tmp.Semester,
			&tmp.Bio,
			&tmp.Residence,
			&tmp.Education,
			&tmp.Department,
			&tmp.Course,
			&tmp.Hostel,
			&tmp.Internship,
		); err != nil {
			return nil, err
		}
		Users = append(Users, tmp)
	}

	return Users, nil
}

func GetRepoNames(c echo.Context) ([]model.Name, error) {
	db := storage.GetDBInstance()
	Names := make([]model.Name, 0)

	rows, err := db.Query("SELECT first_name, last_name FROM \"user\" WHERE (first_name LIKE $1 OR last_name LIKE $2) AND status=TRUE LIMIT 5", string('%')+c.QueryParam("name")+string('%'), string('%')+c.QueryParam("name")+string('%'))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp model.Name
		if err := rows.Scan(
			&tmp.FirstName,
			&tmp.LastName); err != nil {
			return nil, err
		}
		Names = append(Names, tmp)
	}

	return Names, nil
}
