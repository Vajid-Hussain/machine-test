package requestmodels

import (
	"mime/multipart"
	"time"
)

type Resume struct {
	UserID string
	Resume *multipart.FileHeader
}

type Education struct {
	Name string `json:"name"`
}

type Experience struct {
	Name string `json:"name"`
}

type UserResumeData struct {
	UserID     string
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Skills     []string     `json:"skills"`
	Education  []Education  `json:"education"`
	Experience []Experience `json:"experience"`
	Resume     string       `json:"resume"`
}

type CreateJob struct {
	Title            string    `json:"title" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	PostOn           time.Time `json:"post_on"`
	TotalApplication int       `json:"total_application"`
	CompanyName      string    `json:"company_name" validate:"required"`
	PostedBy         string    `json:"posted_by"`
}

type DeleteJob struct {
	JobID string
}

type Pagination struct {
	Offset string `query:"Page" validate:"required"`
	Limit  string `query:"Limit" validate:"required"`
}

type JobSearch struct {
	Job string `query:"search"`
}

type JobApplication struct {
	UserID           string `json:"user_id" validate:"required"`
	JobApplicationID string `json:"job_application_id" validate:"required"`
}
