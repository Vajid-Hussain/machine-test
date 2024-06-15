package domain

import (
	"time"

	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
)

type Users struct {
	ID              int `gorm:"primary key; autoIncrement"`
	Name            string
	Email           string
	Address         string
	UserType        string
	Password        string
	ProfileHeadline string
	UserSince       time.Time
	Status          string `gorm:"default:active"`
}

type UserResumeData struct {
	ID         uint                       `gorm:"primaryKey"`
	UserID     int                        `gorm:"unique"`
	FKUserID   Users                      `gorm:"foreignkey:UserID; references:ID"`
	Name       string                     `gorm:"size:100"`
	Email      string                     `gorm:"size:100"`
	Phone      string                     `gorm:"size:20"`
	Skills     []string                   `gorm:"type:text[]"`
	Education  []requestmodels.Education  `gorm:"type:jsonb"`
	Experience []requestmodels.Experience `gorm:"type:jsonb"`
	Resume     string                     `gorm:"size:300"`
}

type Jobs struct {
	ID               int `gorm:"primarykey; autoIncrement"`
	Title            string
	Description      string
	PostOn           time.Time
	TotalApplication int `gorm:"default:0"`
	CompanyName      string
	PostedBy         string
	Status           string `gorm:"default:active"`
}

type JobApply struct {
	ID        int `gorm:"primarykey; autoIncrement"`
	UserID    int
	FKUserID  Users `gorm:"foreignkey:UserID; references:ID"`
	JobID     int
	FKJobID   Jobs `gorm:"foreignkey:JobID; references:ID"`
	ApplyTime time.Time
	Status    string `gorm:"default:pending"`
}
