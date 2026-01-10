package domain

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserDTO struct {
    ID    string   `json:"id"`
    Email string `json:"email"`
}

type LoginResponse struct {
    Token string  `json:"token"`
    User  UserDTO `json:"user"`
}