package models

type PaginationQuery struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type UniversityQuery struct {
	Name    string `form:"name"`
	Country string `form:"country"`
	City    string `form:"city"`
}

type FacultyQuery struct {
	Name    string `form:"name"`
	Country string `form:"country"`
	City    string `form:"city"`
	Domain  string `form:"domain"`
	Budget  int    `form:"budget"`
}
