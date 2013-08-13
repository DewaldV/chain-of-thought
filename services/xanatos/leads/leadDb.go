package leads

import (
	"github.com/DewaldV/crucible/database"
	"labix.org/v2/mgo/bson"
)

func GetLead(email string) (result *Lead) {
	result = &Lead{}
	conn := database.MgoConnection{"test", "leads"}
	conn.FindOne(bson.M{"email": email}, result)
	return
}

func CreateLead(l Lead) {
	conn := database.MgoConnection{"test", "leads"}
	conn.Insert(l)
	return
}
