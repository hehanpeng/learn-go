package main

import "learn-go/gorm/model"

func main() {
	db := model.Gorm()

	//更新

	//保存所有字段
	var user model.User
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
}
