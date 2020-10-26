package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	ThubImage   string        `bson:"thub_image" json:"thub_image"`
	Description string        `bson:"description" json:"description"`
	Active      bool          `bson:"active" json:"active"`
}
