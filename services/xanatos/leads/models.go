package leads

import (
	"time"
	"code.google.com/p/gorest"
)

type Lead struct {
	Id int
	Email string
	RegistrationDate time.Time
}

type LeadService struct {
	gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`

	getLead 	gorest.EndPoint `method:"GET" 		path:"/leads/{Id:int}" 	output:"Lead"`
	createLead	gorest.EndPoint `method:"POST"		path:"/leads" 			postdata:"Lead"`
	updateLead	gorest.EndPoint `method:"PUT"		path:"/leads/{Id:int}"	postdata:"Lead"`
	deleteLead	gorest.EndPoint	`method:"DELETE"	path:"/leads/{Id:int}"`
}