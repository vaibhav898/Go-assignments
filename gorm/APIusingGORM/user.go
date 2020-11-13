package main

import{
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
}

var db *gorm.DB

    gorm.Model
    Name  string
    Email string
}

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request){
	 db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    var users []User
    db.Find(&users)
    fmt.Println("{}", users)

    json.NewEncoder(w).Encode(users)
}
}

func NewUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(W, "New User Endpoint Hit")

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]
    email := vars["email"]

    db.Create(&User{Name: name, Email: email})
    fmt.Fprintf(w, "New User Successfully Created")

}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	 db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]

    var user User
    db.Where("name = ?", name).Find(&user)
    db.Delete(&user)

    fmt.Fprintf(w, "Successfully Deleted User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	func updateUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    name := vars["name"]
    email := vars["email"]

    var user User
    db.Where("name = ?", name).Find(&user)

    user.Email = email

    db.Save(&user)
    fmt.Fprintf(w, "Successfully Updated User")
}
}