package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//User
type User struct {
	gorm.Model
	Name  string
	Email string
}

//used for populating DB
// var users []User = []User{
// 	User{Name: "Ricky", Email: "Sydney"},
// 	User{Name: "Adam", Email: "Brisbane"},
// 	User{Name: "Justin", Email: "California"},
// }

//InitialMigration connects to database
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
	// for _, user := range users {
	// 	db.Create(&user)
	// }
}

//AllUsers ...
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

//NewUser ...
func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Panic(err, " New user error")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New user created")
}

//DeletUser ...
func DeletUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Panic(err, " delete user error")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, " user deleted")
}

//UpdateUser Updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Panic(err, " Update user error")
	}
	defer db.Close()

}
