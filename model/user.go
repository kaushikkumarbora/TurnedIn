package model

type User struct {
	Id               int64  `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DOB              string `json:"dob"`
	Education        string `json:"education"`
	Email            string `json:"email"`
	Contact          string `json:"contact"`
	Skills           string `json:"skills"`
	Internship       string `json:"internship"`
	YearOfJoining    string `json:"yoj"`
	YearOfCompletion string `json:"yoc"`
	Department       string `json:"department"`
	Course           string `json:"course"`
	Semester         string `json:"semester"`
	Bio              string `json:"bio"`
	Residence        string `json:"residence"`
	Hostel           string `json:"hostel"`
}
