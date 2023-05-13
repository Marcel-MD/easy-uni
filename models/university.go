package models

type University struct {
	Base

	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Ranking string `json:"ranking"`

	Faculties []Faculty `json:"faculties" gorm:"foreignKey:UniversityID;constraint:OnDelete:CASCADE"`
}

type CreateUniversity struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
	City    string `json:"city" binding:"required"`
	Ranking string `json:"ranking" binding:"required"`
}
