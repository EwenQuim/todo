package model

// An item is a struct that represents a todo item
type Item struct {
	ID       uint   `gorm:"primaryKey"` // The id of the item. This is auto incremented
	TodoUUID string `json:"-"`          // The list that the item belongs to
	Content  string
	Done     bool
}
