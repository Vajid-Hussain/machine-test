package domain

type Users struct {
	ID             int `gorm:"primary key; autoIncrement"`
	Name           string
	Email          string
	Address        string
	UserType       string
	Password       string
	ProfileSummary string
}

type Profile struct{
	ID int `gorm:"primay key; autoIncrement"`
	UserID int 
	FKUserID Users `gorm:"foreignkey:ID; references:UserID`
	Skills []string
	Qualification string 
	Experience string
	Phone string 
}
