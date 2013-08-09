package leads

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func GetMgoLead(id int) Lead {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("leads")
	result := Lead{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func CreateMgoLead(l Lead) Lead {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("leads")
	err = c.Insert(l)
	if err != nil {
		panic(err)
	}

	return l
}
