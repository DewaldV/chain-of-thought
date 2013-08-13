package leads

import (
	"github.com/DewaldV/crucible"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func GetLead(email string) (result *Lead) {
	result = &Lead{}
	crucible.ExecuteWithCollection("test", "leads", func(c *mgo.Collection) error { return c.Find(bson.M{"email": email}).One(result) })
	return
}

func CreateLead(l Lead) {
	crucible.ExecuteWithCollection("test", "leads", func(c *mgo.Collection) error { return c.Insert(l) })
}
