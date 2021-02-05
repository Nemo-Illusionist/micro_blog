package auth

type PasswordRequest struct {
	Password string `json:"password"`
}

type LoginRequest struct {
	PasswordRequest
	Login string `json:"Login"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegistrationRequest struct {
	LoginRequest
	Name string `json:"name"`
}

type RegistrationResponse struct {
	LoginResponse
	UserId uint64 `json:"user_id"`
}
