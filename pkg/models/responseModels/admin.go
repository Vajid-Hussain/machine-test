package responsemodels

import "time"

type AdminLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Job struct {
	ID               string
	Title            string    `json:"title" gorm:"column:title"`
	Description      string    `json:"description"`
	PostOn           time.Time `json:"post_on"`
	TotalApplication int       `json:"total_application"`
	CompanyName      string    `json:"company_name"`
	PostedBy         string    `json:"posted_by"`
	Status           string    `json:"status"`
}
