package requestmodels


type AdminLogin struct {
	Email           string `json:"email" validate:"email"`
	Password        string `json:"password" validate:"min=5"`
}

