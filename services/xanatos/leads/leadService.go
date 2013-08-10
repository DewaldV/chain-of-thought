package leads

import (
	"code.google.com/p/gorest"
	"github.com/DewaldV/chain-of-thought/services/xanatos/crucible"
	"time"
)

type Lead struct {
	Email            string
	RegistrationDate time.Time
}

type LeadService struct {
	gorest.RestService `root:"/leads" consumes:"application/json" produces:"application/json"`

	doOptions  gorest.EndPoint `method:"OPTIONS"	path:"/{...:string}"`
	getLead    gorest.EndPoint `method:"GET" 		path:"/{Email:string}"	output:"Lead"`
	createLead gorest.EndPoint `method:"POST"		path:"/" 				postdata:"Lead"`
	updateLead gorest.EndPoint `method:"PUT"		path:"/{Email:string}"	postdata:"Lead"`
	deleteLead gorest.EndPoint `method:"DELETE"		path:"/{Email:string}"`

	serviceConfig *crucible.ServiceConfigStruct
}
