package crucible

import (
	"labix.org/v2/mgo"
)

func ExecuteWithCollection(database, collection string, f func(*mgo.Collection) error) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)

	return f(c)
}
