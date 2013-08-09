package leads

import (
	"code.google.com/p/gorest"
	"fmt"
	"time"
)

func Register() bool {
	gorest.RegisterService(new(LeadService))
	return true
}

func (serv LeadService) DoOptions() {
	rb := serv.ResponseBuilder()
	rb.AddHeader("Access-Control-Allow-Origin", "http://localhost:8080")
	rb.AddHeader("Access-Control-Allow-Method", "GET, PUT, POST, DELETE")
	rb.AddHeader("Access-Control-Allow-Headers", "accept, origin, x-requested-with, content-type")

	fmt.Printf("Received an Options request, responding with some options.")
}

func (serv LeadService) GetLead(email string) (l Lead) {
	l = *GetLead(email)
	return
}

func (serv LeadService) CreateLead(l Lead) {
	l.RegistrationDate = time.Now()
	serv.ResponseBuilder().AddHeader("Access-Control-Allow-Origin", "http://localhost:8080")
	fmt.Printf("Received a lead with email: %s /n", l.Email)
	CreateLead(l)
	serv.ResponseBuilder().SetResponseCode(200)
	return
}

func (serv LeadService) UpdateLead(l Lead, email string) {
	l.RegistrationDate = time.Now()
	return
}

func (serv LeadService) DeleteLead(email string) {
	return
}
