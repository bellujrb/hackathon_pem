package db

import (
	"gorm.io/gorm"
)

func Create(db *gorm.DB, table interface{}) error {
	return db.Create(table).Error
}

func Read(db *gorm.DB, table interface{}, condition interface{}) error {
	return db.Where(condition).First(table).Error
}

func UpdateByID(db *gorm.DB, table interface{}, id int64) error {
	return db.Where("id = ?", id).Updates(table).Error
}

func Delete(db *gorm.DB, table interface{}, condition interface{}) error {
	return db.Where(condition).Delete(table).Error
}

func GetByID(db *gorm.DB, table interface{}, id int64) (bool, error) {
	err := db.First(table, id).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func GetAll(db *gorm.DB, tableSlice interface{}) error {
	return db.Find(tableSlice).Error
}

func GetAllWithCondition(db *gorm.DB, tableSlice interface{}, condition string, args ...interface{}) error {
	return db.Where(condition, args...).Find(tableSlice).Error
}

func InScope(db *gorm.DB, tableSlice interface{}) error {
	return db.Where("in_scope = ?", true).Find(tableSlice).Error
}
func GetAllByCategory(db *gorm.DB, tableSlice interface{}, category string) error {
	return db.Where("category = ?", category).Find(tableSlice).Error
}
