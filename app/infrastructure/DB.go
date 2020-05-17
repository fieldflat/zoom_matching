package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is a struct
type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Connect  *gorm.DB
}

// NewDB is a function
func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

// NewTestDB is a function
func NewTestDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Test.Host,
		Username: c.DB.Test.Username,
		Password: c.DB.Test.Password,
		DBName:   c.DB.Test.DBName,
	})
}

// newDB is a function
func newDB(d *DB) *DB {
	// https://github.com/go-sql-driver/mysql#examples
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connect = db
	return d
}

// First is a function
func (db *DB) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.First(out, where...)
}

// Find is a function
func (db *DB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.Find(out, where...)
}

// Create is a function
func (db *DB) Create(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.Create(out)
}

// Save is a function
func (db *DB) Save(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.Save(out)
}

// Delete is a function
func (db *DB) Delete(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.Delete(out, where...)
}

// Where is a function
func (db *DB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Connect.Where(query, args...)
}
