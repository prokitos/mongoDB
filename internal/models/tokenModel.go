package models

type TokenDB struct {
	GUID         string
	RefreshToken string
}

type TokenResponser struct {
	AccessToken  string
	RefreshToken string
}
