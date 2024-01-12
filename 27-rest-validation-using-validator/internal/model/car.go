package model

// Car represents a car data model
type Car struct {
    Make  string `json:"make" validate:"required,alpha" errMsg:"make is required and must be alphabetic"`
    Model string `json:"model" validate:"required,alpha" errMsg:"model is required and must be alphabetic"`
    Year  int    `json:"year" validate:"required,numeric,min=1980,max=2022" errMsg:"year is required and must be between 1980 and 2022"`
}