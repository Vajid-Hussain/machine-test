package responsemodels

import "time"

type UserSignup struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Address         string `json:"address"`
	ProfileHeadline string `json:"profileHeadline"`
	UserSince       string `json:"userSince"`
	UserType        string `json:"userType"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
}

type UserLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserResumeData struct {
	ID         uint
	UserID     uint
	Name       string
	Email      string
	Phone      string
	Skills     []string `json:"-" gorm:"-"`
	Education  []string `json:"-" gorm:"-"`
	Experience []string `json:"-" gorm:"-"`
	Resume     string
}

type JobApplication struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"posted_by"`
	Email     string `json:"email"`
	JobID     string `json:"job_id"`
	Title     string
	ApplyTime time.Time `json:"apply_time"`
	Status    string    `json:"status"`
}

type JobApplicationAdmin struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"posted_by"`
	Email     string `json:"email"`
	JobID     string `json:"job_id"`
	Address   string
	Title     string
	ApplyTime time.Time `json:"apply_time"`
	Status    string    `json:"status"`
}
