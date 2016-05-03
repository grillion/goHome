package db

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	"time"
)

//Device Object
type Device struct {
	ID bson.ObjectId	`bson:"_id,omitempty" json:"_id,omitempty" `
	Name string			`json:"name"`
	Type string			`json:"type"`
	Created time.Time	`json:"created"`
	Updated time.Time	`json:"updated"`
}

// DB Interface
type Devices struct {}

func (D Devices) Create(d Device) (*Device, error) {
	if d.Name == "" {
		return nil, errors.New("Name is required")
	}

	d.Created = time.Now()
	d.Updated = d.Created

	insertError := GetC("devices").Insert(d)

	if insertError != nil {
		return nil, insertError
	}

	return &d, nil
}

func (D Devices) Get(id bson.ObjectId) (device *Device, err error) {
	err = GetC("devices").FindId(id).One(&device)
	return
}

func (D Devices) GetAll() (devices []Device, err error) {
	err = GetC("devices").Find(nil).All(&devices)
	if devices == nil {
		devices = []Device{}
	}
	return
}

func (D Devices) Find(query interface{}) (results []Device, err error) {
	err = GetC("devices").Find(query).All(&results)
	if results == nil {
		results = []Device{}
	}
	return
}

func (D Devices) FindOne(query interface{}) (result Device, err error) {
	err = GetC("devices").Find(query).One(&result)
	return
}


func (D Devices) Update(d *Device) error {
	d.Updated = time.Now()
	return GetC("devices").UpdateId(d.ID, d)

}

func (D Devices) RemoveDevice(id bson.ObjectId) error {
	return GetC("devices").RemoveId(id)
}

