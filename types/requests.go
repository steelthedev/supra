package types

type UserSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
