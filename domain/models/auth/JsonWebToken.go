package auth

type JsonWebToken struct {
	Token     string `json:"token"`
	Email     string `json:"email"`
	UserId    string `json:"userId"`
	ExpiresIn int64  `json:"expiresIn"`
}
