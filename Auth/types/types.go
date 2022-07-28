package types

import "time"

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    time.Duration
	RtExpires    time.Duration
}

type AccessDetails struct {
	AccessUuid string
	UserId     int64
}

type AuthToken struct {
	Token string `reqHeader:"Authorization"`
}
