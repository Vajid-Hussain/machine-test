package requestmodels

type UserSignup struct {
	Name            string `json:"name" validate:"required,alpha"`
	Email           string `json:"email" validate:"email"`
	Address         string `json:"address" validate:"required"`
	ProfileHeadline string `json:"profileHeadline" validate:"max=30"`
	Password        string `json:"password" validate:"min=5"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
}

type UserLogin struct {
	Email           string `json:"email" validate:"email"`
	Password        string `json:"password" validate:"min=5"`
}
