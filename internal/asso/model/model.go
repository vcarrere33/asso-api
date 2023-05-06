package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Association struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name            string             `bson:"name,omitempty" json:"name"`
	IDRna           string             `bson:"idRna,omitempty" json:"idRna"`
	DateCreation    string             `bson:"dateCreation,omitempty" json:"dateCreation"`
	DatePublication string             `bson:"datePublication,omitempty" json:"datePublication"`
	Object          string             `bson:"object,omitempty" json:"object"`
	ObjetSocial1    string             `bson:"objetSocial1,omitempty" json:"objetSocial1"`
	ObjetSocial2    string             `bson:"objetSocial2,omitempty" json:"objetSocial2"`
	Address         string             `bson:"address,omitempty" json:"address"`
	ZipCode         string             `bson:"zipCode,omitempty" json:"zipCode"`
	City            string             `bson:"city,omitempty" json:"city"`
}
