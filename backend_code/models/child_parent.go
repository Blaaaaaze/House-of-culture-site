package models

type ChildParent struct {
	ID       int `json:"id"`
	ParentID int `json:"parent_id"`
	ChildID  int `json:"child_id"`
}
