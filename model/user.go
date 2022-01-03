package model

type User struct {
	Id               int64  `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DOB              string `json:"dob"`
	Email            string `json:"email"`
	Contact          string `json:"contact_no"`
	Skills           string `json:"skills"`
	YearOfJoining    string `json:"year_of_admission"`
	YearOfCompletion string `json:"year_of_completion"`
	Semester         int16  `json:"semester"`
	Bio              string `json:"bio"`
	Residence        string `json:"residence"`
	Education        string `json:"education"`
	Department       string `json:"department"`
	Course           string `json:"course"`
	Hostel           string `json:"hostel"`
	Internship       string `json:"internship"`
}
