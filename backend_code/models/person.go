package models

type RoleType string

const (
	RoleTeacher   RoleType = "teacher"
	RoleChild     RoleType = "child"
	RoleEmployee  RoleType = "employee"
	RoleParent    RoleType = "parent"
	RoleApplicant RoleType = "applicant"
)

type Person struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Lastname    *string  `json:"lastname,omitempty"`
	GroupID     *int     `json:"group_id,omitempty"`
	PhoneNumber string   `json:"phone_number"`
	Role        RoleType `json:"role"`
	IsActive    bool     `json:"is_active"`
	VacancyID   *int     `json:"vacancy_id,omitempty"`
}
