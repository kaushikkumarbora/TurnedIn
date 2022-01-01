package model

type ConnectionRequest struct {
	RequesterID   string `json:"reqid"`
	RequesterName string `json:"reqname"`
}
