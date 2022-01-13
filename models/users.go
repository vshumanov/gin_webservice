package models

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	F_name string `json:"f_name" gorm:"not null"`
	L_name string `json: "l_name" gorm:"not null"`
	Email  string `json: "email" gorm:"not null"`
}

type UsersInterface interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	CreateUser(u *User) (User, error)
	UpdateUser(u *User, id int)
	DeleteUser(id int)
}

func (db Database) GetAllUsers() (users []User, err error) {
	users = []User{}
	err = db.DB.Find(&users).Error
	return users, err
}

func (db Database) GetUserByID(id string) (user User, err error) {
	user = User{}
	err = db.DB.First(&user, id).Error
	return user, err
}

func (db Database) CreateUser(user *User) (err error) {
	return db.DB.Create(user).Error
}

func (db Database) UpdateUser(user *User, id string) {
	db.DB.Save(user)
}

func (db Database) DeleteUser(id string) {
	user := User{}
	db.DB.Where("id = ?", id).Delete(user)
}
