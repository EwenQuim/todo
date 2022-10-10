package model

// Todo is a struct that represents a todo item
type Todo struct {
	UUID   string `gorm:"primary_key"`
	Title  string
	Public bool
	Items  []Item  `gorm:"-" json:",omitempty"`
	Groups []Group `gorm:"-" json:",omitempty"`
}

type Group struct {
	Name  string
	Items []Item
}
