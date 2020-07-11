package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sundowndev/castle"
	)

type Profile struct {
	Role castle.RoleInterface
}

func (p *Profile) HasRole(role castle.RoleInterface) bool {
	return p.Role.String() == role.String()
}

func (p *Profile) SetRole(role castle.RoleInterface) {
	p.Role = role
}

func (p *Profile) String() string {
	return p.Role.String()
}
// --------------------

type User struct {
	gorm.Model
	Profile `gorm:"type:varchar(100)"`
	Username string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Author User
	Title string
}

var app *castle.Application

var create *castle.Permission
var edit *castle.Permission
var delete *castle.Permission
var read *castle.Permission

var admin castle.RoleInterface
var guest castle.RoleInterface

func main() {
	app, err := castle.NewApplication("myapp")
	if err != nil {
		panic(err) // Validation error
	}

	create = app.NewPermission()
	edit = app.NewPermission()
	delete = app.NewPermission()
	read = app.NewPermission()

	guest = app.NewRole("guest", read)
	admin = app.NewRole("admin", create, edit, delete).InheritFrom(guest)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})

	user := User{Username: "sundown"}
	user.SetRole(guest)

	if user.Role.HasPermission(create) {
		panic(errors.New("guest can create!!!"))
	}

	post := Post{Title: "L1212", Author: user}

	// Create
	db.Create(&user)
	db.Create(&post)
}