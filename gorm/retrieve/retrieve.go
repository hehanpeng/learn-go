package main

import (
	"errors"
	"fmt"
	"learn-go/gorm/model"

	"gorm.io/gorm/clause"

	"github.com/golang-module/carbon"

	"gorm.io/gorm"
)

func main() {
	db := model.Gorm()

	var user model.User
	var users []model.User
	// 检索单个对象
	// 获取第一条记录（主键升序）
	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 用主键检索
	db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	// 检索全部对象
	// 获取全部记录
	result = db.Find(&users)
	// SELECT * FROM users;
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// 条件
	// String 条件

	// 获取第一条匹配的记录
	db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// 获取全部匹配的记录
	db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	// 三周前
	lastWeek := carbon.Parse(carbon.Now().ToDateString()).SubWeeks(1).ToDateTimeString() // 2020-02-08 13:14:15
	db.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", lastWeek, carbon.Parse(carbon.Now().ToDateString()).ToDateTimeString()).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// Struct & Map 条件
	// Struct
	db.Where(&model.User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键切片条件
	db.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	// 指定结构体查询字段
	db.Where(&model.User{Name: "jinzhu", Age: 0}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	db.Where(&model.User{Name: "jinzhu", Age: 0}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;

	// 内联条件
	// SELECT * FROM users WHERE id = 23;
	// 根据主键获取记录，如果是非整型主键
	db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// Plain SQL
	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, model.User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;

	// Not条件
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(model.User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// 不在主键切片中的记录
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	// Or 条件
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(model.User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// 选择特定字段
	db.Select("name", "age").Find(&users)
	// SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age FROM users;

	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	// SELECT COALESCE(age,'42') FROM users;

	// Order
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// 多个 order
	db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&model.User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

	// Limit & Offset
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	var users1 []model.User
	var users2 []model.User
	// 通过 -1 消除 Limit 条件
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// 通过 -1 消除 Offset 条件
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)

	// Group & Having
	var result1 []Result
	db.Model(&model.User{}).Select("name, sum(age) as total").Where("name LIKE ?", "jin%").Group("name").Find(&result1)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "jin%" GROUP BY `name`

	db.Model(&model.User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "Jinzhu").Find(&result1)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

	rows, _ := db.Table("users").Select("date(created_at) as date, sum(age) as ages").Group("date(created_at)").Rows()
	for rows.Next() {
	}

	rows, _ = db.Table("users").Select("date(created_at) as date, sum(age) as ages").Group("date(created_at)").Having("ages > ?", 30).Rows()
	for rows.Next() {
	}
	db.Table("users").Select("name, sum(age) as total").Group("name").Having("total >= ?", 10).Scan(&result1)
	// Distinct
	db.Where("name is not null").Distinct("name", "age").Order("name, age desc").Find(&users)
	fmt.Println("complete")
}

type Result struct {
	Name  string
	Total int
}
