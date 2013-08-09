package repository

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Registerable interface {
	Register() bool
}

//Create
//Read
//Update
//Delete
