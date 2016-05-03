package db

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	"time"
)

//User Object
type User struct {
	ID bson.ObjectId	`bson:"_id,omitempty" json:"_id,omitempty" `
	Name string			`json:"name"`
	Username string		`json:"username"`
	Email string		`json:"email"`
	Password string 	`json:"password"`
	Created time.Time	`json:"created"`
	Updated time.Time	`json:"updated"`
}

// DB Interface
type Users struct {}

func (U Users) Create(u User) (*User, error) {
	if u.Name == "" {
		return nil, errors.New("Name is required")
	}
	if u.Password == "" {
		return nil, errors.New("Password is required")
	}

	u.Created = time.Now()
	u.Updated = u.Created

	insertError := GetC("users").Insert(u)

	if insertError != nil {
		return nil, insertError
	}

	return &u, nil
}

func (U Users) Get(id bson.ObjectId) (user *User, err error) {
	err = GetC("users").FindId(id).One(&user)
	return
}

func (U Users) GetAll() (users []User, err error) {
	err = GetC("users").Find(nil).All(&users)
	if users == nil {
		users = []User{}
	}
	return
}

func (U Users) Find(query interface{}) (results []User, err error) {
	err = GetC("users").Find(query).All(&results)
	if results == nil {
		results = []User{}
	}
	return
}

func (U Users) FindOne(query interface{}) (result User, err error) {
	err = GetC("users").Find(query).One(&result)
	return
}


func (U Users) Update(u *User) error {
	u.Updated = time.Now()
	return GetC("users").UpdateId(u.ID, u)

}

func (U Users) RemoveUser(id bson.ObjectId) error {
	return GetC("users").RemoveId(id)
}

