package query

import (
	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/database"
)

func NewItem(s database.Service, item model.Item) (model.Item, error) {
	item.ID = model.NewUUID()
	err := s.DB.Create(&item).Error
	return item, err
}

func GetItemsForList(s database.Service, todoUUID string) ([]model.Item, error) {
	items := []model.Item{}
	err := s.DB.Where("todo_uuid = ?", todoUUID).Find(&items).Error
	return items, err
}

func DeleteItem(s database.Service, id string) error {
	return s.DB.Delete(&model.Item{}, "id = ?", id).Error
}

func ChangeItem(s database.Service, id, newContent string) error {
	itemToEdit := model.Item{}
	err := s.DB.Find(&itemToEdit, "id = ?", id).Error //.Update("done", true)
	if err != nil {
		return err
	}

	itemToEdit.Content = newContent

	return s.DB.Save(&itemToEdit).Error
}

func SwitchItem(s database.Service, id string) error {
	itemToEdit := model.Item{}
	err := s.DB.Find(&itemToEdit, "id = ?", id).Error //.Update("done", true)
	if err != nil {
		return err
	}

	itemToEdit.Done = !itemToEdit.Done

	return s.DB.Save(&itemToEdit).Error
}
