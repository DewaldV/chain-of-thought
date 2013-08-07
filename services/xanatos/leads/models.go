package leads

import (
	"code.google.com/p/gorest"
	"time"
)

type Lead struct {
	Id               int
	Email            string
	RegistrationDate time.Time
}

type LeadService struct {
	gorest.RestService `root:"/leads" consumes:"application/json" produces:"application/json"`

	getLead    gorest.EndPoint `method:"GET" 	path:"/{Id:int}"	output:"Lead"`
	createLead gorest.EndPoint `method:"POST"	path:"/" 			postdata:"Lead"`
	updateLead gorest.EndPoint `method:"PUT"	path:"/{Id:int}"	postdata:"Lead"`
	deleteLead gorest.EndPoint `method:"DELETE"	path:"/{Id:int}"`
}
