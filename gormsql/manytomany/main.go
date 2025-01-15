package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	// 使用 gorm:"many2many 建立多对多关系
	Languages []Language `gorm:"many2many:student_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

type Student1 struct {
	gorm.Model
	Languages1 []*Language1 `gorm:"many2many:student1_languages1;"`
}

type Language1 struct {
	gorm.Model
	Name     string
	Student1 []*Student1 `gorm:"many2many:student1_languages1;"`
}

type Student2 struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:student2_profiles;foreignKey:Refer;joinForeignKey:Student2ReferID;References:Student2Refer;joinReferences:ProfileRefer"`
	Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
	gorm.Model
	Name          string
	Student2Refer uint `gorm:"index:,unique"`
}

type People struct {
	ID      int
	Name    string
	Address []PeopleAddress `gorm:"many2many:people_address;"`
}

type Address struct {
	ID   int
	Name string
}

type PeopleAddress struct {
	PeopleID  int `gorm:"primaryKey"`
	AddressID int `gorm:"primaryKey"`
	CreatedAt time.Time
}

type Book struct {
	ID        int `gorm:"primarykey"`
	Name      string
	Author    string
	BookTypes []BookType `gorm:"many2many:book_book_type;"`
}

type BookType struct {
	ID   int `gorm:"primarykey"`
	Name string
	Book []Book `gorm:"many2many:book_book_type;"`
}

func main() {
	dsn := "root:123456@tcp(10.1.131.32:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//DB.AutoMigrate(&Book{}, &BookType{})
	//booktype := []BookType{
	//	{Name: "1111"},
	//	{Name: "2222"},
	//	{Name: "3333"},
	//}
	//book := []Book{
	//	{Name: "xiyou", Author: "zhangsan", BookType: []BookType{
	//		{ID: 1},
	//		{ID: 2},
	//	}},
	//	{Name: "sanguo", Author: "lisi", BookType: []BookType{
	//		{ID: 2},
	//		{ID: 3},
	//	}},
	//}
	//
	var books Book = Book{ID: 1}
	var books2 = []Book{{ID: 1}, {ID: 2}}
	var book_types []BookType
	err = DB.Model(&books).Association("BookTypes").Error
	fmt.Println(err)
	fmt.Println(DB.Debug().Model(&books).Where("id = ?", 2).Association("BookTypes").Count())
	DB.Debug().Model(&books2).Association("BookTypes").Find(&book_types)
	fmt.Println(book_types)
	//err = DB.Debug().Model(&books).Where("id = ?", 1).Association("BookTypes").Find(&book_types)
	//fmt.Println(book_types)
	//DB.Preload("BookTypes").Find(&books)
	//DB.Debug().Model(&books).Association("BookTypes").Append(&BookType{ID: 3})
	//fmt.Println("1111")
	//DB.Debug().Model(&books).Association("BookTypes").Find(&book_types)
	//fmt.Println(book_types)
	var book3 []Book
	DB.Debug().Preload("BookTypes").Find(&book3)
	fmt.Println(book3)

}
