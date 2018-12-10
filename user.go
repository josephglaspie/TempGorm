package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//  db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")

type User struct {
	gorm.Model
	Name  string
	Email string
}

//used for populating DB
var users []User = []User{
	User{Name: "Ricky", Email: "Sydney"},
	User{Name: "Adam", Email: "Brisbane"},
	User{Name: "Justin", Email: "California"},
}

func InitialMigration() {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err, " Database connection issue")
	}
	fmt.Println("Database connection successful")
	defer db.Close()
	//db.Debug().DropTableIfExists(&User{})
	db.Debug().AutoMigrate(&User{})

	//Populate users in DB
	for _, user := range users {
		db.Create(&user)
	}
}
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err, " Could not connect to databbase AllUsers")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpoint")
}

func DeletUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete user endpoint")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user endpoint")
}
