package leads

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func GetMgoLead(email string) (result *Lead) {
	result = &Lead{}
	executeMgo(func(c *mgo.Collection) error { return c.Find(bson.M{"email": email}).One(result) }, "test", "leads")
	return
}

func CreateMgoLead(l Lead) {
	executeMgo(func(c *mgo.Collection) error { return c.Insert(l) }, "test", "leads")
}

func executeMgo(f func(c *mgo.Collection) error, database string, collection string) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)

	return f(c)
}
