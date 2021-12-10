package model

// Todo is a struct that represents a todo item
type Todo struct {
	UUID   string `gorm:"primary_key"`
	Title  string
	Public bool
	Items  []Item            `gorm:"-" json:",omitempty"`
	Groups map[string][]Item `gorm:"-" json:",omitempty"`
}
