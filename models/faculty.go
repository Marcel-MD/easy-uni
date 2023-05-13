package models

import "github.com/lib/pq"

type Faculty struct {
	Base

	Name                 string         `json:"name"`
	Domains              pq.StringArray `json:"domains" gorm:"type:text[]"`
	About                string         `json:"about"`
	Budget               int            `json:"budget"`
	Duration             float32        `json:"duration"`
	ApplyDate            string         `json:"apply_date"`
	StartDate            string         `json:"start_date"`
	AcademicRequirements string         `json:"academic_requirements"`
	LanguageRequirements string         `json:"language_requirements"`
	OtherRequirements    string         `json:"other_requirements"`

	UniversityID string     `json:"university_id"`
	University   University `json:"university" gorm:"foreignKey:UniversityID"`
}

type CreateFaculty struct {
	Name                 string   `json:"name" binding:"required"`
	Domains              []string `json:"domains" binding:"required"`
	About                string   `json:"about"`
	Budget               int      `json:"budget" binding:"required"`
	Duration             float32  `json:"duration"`
	ApplyDate            string   `json:"apply_date"`
	StartDate            string   `json:"start_date"`
	AcademicRequirements string   `json:"academic_requirements"`
	LanguageRequirements string   `json:"language_requirements"`
	OtherRequirements    string   `json:"other_requirements"`
}
