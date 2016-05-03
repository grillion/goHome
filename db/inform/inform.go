package inform

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/grillion/goHome/db"
	"github.com/grillion/mFi"
)

//User Object
type DbInformPacket struct {
	ID bson.ObjectId	`bson:"_id,omitempty" json:"_id,omitempty"`
	mFi.InformPacket 	`json:"payload"`
	Created time.Time	`json:"created"`
}

func Save(packet *mFi.InformPacket) error {

	dbPacket := DbInformPacket{}
	dbPacket.InformPacket = packet
	dbPacket.Created = time.Now()

	insertError := db.GetC("inform").Insert(dbPacket)

	if insertError != nil {
		return insertError
	}

	return nil
}
