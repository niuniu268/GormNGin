package models

import "GinNGorm/dao"

type User struct {
	ID      int
	Name    string
	Contact string
	Payment string
	History string
	Booking int
	Type    string
}

func (User) TableName() string {
	return "users"
}

func GetUserTest(id int) User {

	var user User

	//Db, _ := gorm.Open("mysql", config.Mysqldb)
	//Db, _ := gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{})

	dao.Db.Where("id=?", id).First(&user)

	return user

}
