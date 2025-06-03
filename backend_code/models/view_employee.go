package models

type EmployeeView struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Lastname    *string  `json:"lastname,omitempty"`
	PhoneNumber string   `json:"phone_number"`
	Role        RoleType `json:"role"` // ENUM: teacher, employee
	GroupID     *int     `json:"group_id,omitempty"`
	IsActive    bool     `json:"is_active"`
}
