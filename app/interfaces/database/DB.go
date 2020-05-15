package database

import (
	"github.com/jinzhu/gorm"
)

// DB is a type
type DB interface {
	First(out interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Create(out interface{}, where ...interface{}) *gorm.DB
	Delete(out interface{}, where ...interface{}) *gorm.DB
}
