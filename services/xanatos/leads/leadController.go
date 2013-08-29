package leads

import (
	"code.google.com/p/gorest"
	"fmt"
	"github.com/DewaldV/crucible"
	"github.com/DewaldV/crucible/logging"
	"time"
)

func init() {
	Register()
}

func Register() bool {
	serv := &LeadService{serviceConfig: crucible.Conf.Services["Leads"]}
	gorest.RegisterService(serv)
	return true
}

func (serv LeadService) CreateLead(l Lead) {
	fmt.Printf(logging.LogRequest(serv.Context.Request()))
	l.RegistrationDate = time.Now()
	CreateLead(l)
	fmt.Printf("Received a create request for a lead with email: %s\n", l.Email)
	return
}

func (serv LeadService) GetLead(email string) (l Lead) {
	fmt.Printf(logging.LogRequest(serv.Context.Request()))
	l = *GetLead(email)
	fmt.Printf("Received a get request for a lead with email: %s\n", email)
	return
}

func (serv LeadService) UpdateLead(l Lead, email string) {
	fmt.Printf(logging.LogRequest(serv.Context.Request()))
	l.RegistrationDate = time.Now()
	fmt.Printf("Received a update request for a lead with email: %s\n", email)
	return
}

func (serv LeadService) DeleteLead(email string) {
	fmt.Printf(logging.LogRequest(serv.Context.Request()))
	fmt.Printf("Received a delete request for a lead with email: %s\n", email)
	return
}
