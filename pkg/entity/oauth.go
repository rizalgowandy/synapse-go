package entity

import "github.com/rizalgowandy/synapse-go/pkg/enum"

type CreateOAuthKeyReq struct {
	UserID        string       `json:"-"`
	UserIPAddress string       `json:"-"`
	ValidationPin string       `json:"validation_pin"`
	PhoneNumber   string       `json:"phone_number"`
	Scope         []enum.Scope `json:"scope"`
	RefreshToken  string       `json:"refresh_token"`
}

type CreateOAuthKeyResp struct {
	ClientID         string       `json:"client_id"`
	ClientName       string       `json:"client_name"`
	ExpiresAt        string       `json:"expires_at"`
	ExpiresIn        string       `json:"expires_in"`
	OauthKey         string       `json:"oauth_key"`
	RefreshExpiresIn int          `json:"refresh_expires_in"`
	RefreshToken     string       `json:"refresh_token"`
	Scope            []enum.Scope `json:"scope"`
	UserID           string       `json:"user_id"`
}

type GenerateRefreshTokenReq struct {
	UserID        string `json:"-"`
	UserIPAddress string `json:"-"`
	Email         string `json:"email"`
	Password      string `json:"password"`
}

type GenerateRefreshTokenResp struct {
	ID           string       `json:"_id"`
	V            int          `json:"_v"`
	ChangeIn     int          `json:"change_in"`
	ClientID     string       `json:"client_id"`
	ClientName   string       `json:"client_name"`
	FpRegistered bool         `json:"fp_registered"`
	IsActive     bool         `json:"is_active"`
	RefreshToken string       `json:"refresh_token"`
	Scope        []enum.Scope `json:"scope"`
	UserID       string       `json:"user_id"`
}
