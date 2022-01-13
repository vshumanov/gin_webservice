package main

import (
	u_models "gin_webservice/models"
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var users = []u_models.User{{
	ID:     1,
	F_name: "user01",
	L_name: "test",
	Email:  "test_1@test.com",
}, {
	ID:     2,
	F_name: "user02",
	L_name: "test",
	Email:  "test_2@test.com",
}}

var gdb *gorm.DB

func setup() *u_models.Database {
	return &u_models.Database{
		DB: gdb,
	}
}

func TestMain(m *testing.M) {

	var err error

	db_name := "test.db"

	os.Remove(db_name)
	gdb, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	gdb.AutoMigrate(&u_models.User{})

	gdb.Create(&users)

	os.Exit(m.Run())
}

func TestGetUsers_Success(t *testing.T) {
	db := setup()
	res_users, err := db.GetAllUsers()
	if err != nil {
		t.Fatalf("Error on getting users %s", err)
	}
	if len(res_users) != len(users) {
		t.Fail()
	}
}
