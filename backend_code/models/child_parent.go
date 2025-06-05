package models

// ChildParent описывает связь ребенка и родителя
type ChildParent struct {
	ID       int `json:"id"`
	ParentID int `json:"parent_id"`
	ChildID  int `json:"child_id"`
}
