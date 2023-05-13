package models

type University struct {
	Base

	Name    string `json:"name"`
	About   string `json:"about"`
	Country string `json:"country"`
	City    string `json:"city"`
	Ranking string `json:"ranking"`
	ImgLink string `json:"img_link"`

	Faculties []Faculty `json:"faculties" gorm:"foreignKey:UniversityID;constraint:OnDelete:CASCADE"`
}

type CreateUniversity struct {
	Name    string `json:"name" binding:"required"`
	About   string `json:"about" binding:"required"`
	Country string `json:"country" binding:"required"`
	City    string `json:"city" binding:"required"`
	Ranking string `json:"ranking" binding:"required"`
	ImgLink string `json:"img_link" binding:"required"`
}
