package jwt_token

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type JWTRefreshClaim struct {
	UserID string `json:"user_id"`
	//Email       string `json:"email"`
	//RefreshId   string
	//TokenString string
	jwt.RegisteredClaims
}

//type JWTClaim struct {
//	UserID string `json:"user_id"`
//	Email  string `json:"email"`
//	jwt.StandardClaims
//}
