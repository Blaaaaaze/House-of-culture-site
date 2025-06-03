package models

type ChildWithParent struct {
	ChildID       int    `json:"child_id"`
	ChildName     string `json:"child_name"`
	ChildSurname  string `json:"child_surname"`
	ChildLastname string `json:"child_lastname"`
	ChildPhone    string `json:"child_phone"`

	ParentID       int    `json:"parent_id"`
	ParentName     string `json:"parent_name"`
	ParentSurname  string `json:"parent_surname"`
	ParentLastname string `json:"parent_lastname"`
	ParentPhone    string `json:"parent_phone"`
}
