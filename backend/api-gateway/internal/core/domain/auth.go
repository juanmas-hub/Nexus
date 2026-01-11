package domain

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type registerRequest struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=6"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
}

type UserDTO struct {
    ID    string   `json:"id"`
    Email string `json:"email"`
}

type LoginResponse struct {
    Token string  `json:"token"`
    User  UserDTO `json:"user"`
}

type RegisterResponse struct {
    Message string `json:"message"`
    User    struct {
        ID        string `json:"id"`
        Email     string `json:"email"`
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
    } `json:"user"`
}