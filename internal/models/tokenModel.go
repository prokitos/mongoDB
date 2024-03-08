package models

type TokenDB struct {
	GUID         string `json:"GUID,omitempty"`
	RefreshToken string `json:"RefreshToken,omitempty"`
}

type TokenResponser struct {
	AccessToken  string
	RefreshToken string
}
