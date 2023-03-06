package model

// An item is a struct that represents a todo item
type Item struct {
	ID       UUID   `gorm:"primaryKey"` // The id of the item
	TodoUUID string `json:"todoID"`     // The list that the item belongs to
	Content  string
	Done     bool
}
