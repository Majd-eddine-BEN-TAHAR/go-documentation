package model

type User struct {
    Name  string `json:"name" validate:"required,alpha" errMsg:"Name is required and must be alphabetic"`
    Email string `json:"email" validate:"required,email" errMsg:"Email is required and must be a valid email"`
    Age   int    `json:"age" validate:"required,numeric,min=18,max=100" errMsg:"Age is required and must be between 18 and 100"`
}