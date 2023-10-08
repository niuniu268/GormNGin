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

func GetUser(id int) User {

	var user User

	//Db, _ := gorm.Open("mysql", config.Mysqldb)
	//Db, _ := gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{})

	dao.Db.Where("id=?", id).First(&user)

	return user

}

func AddUser(Name string,
	Contact string,
	Payment string,
	History string,
	Booking int,
	Type string) (int, error) {

	user := User{

		Name:    Name,
		Contact: Contact,
		Payment: Payment,
		History: History,
		Booking: Booking,
		Type:    Type,
	}

	err := dao.Db.Create(&user).Error
	return user.ID, err

}

func UpdateUser(user *User) (int, error) {

	err := dao.Db.Save(&user).Error

	return user.ID, err
}

func DeleteUser(id int) (user User, err error) {

	err = dao.Db.Where("id=?", id).Delete(&user).Error

	return user, err

}
