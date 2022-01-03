package model

type Connection struct {
	LUserID int64  `json:"l_user_id"`
	RUserID int64  `json:"r_user_id"`
	Status  string `json:"status"`
}
