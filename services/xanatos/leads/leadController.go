package leads

import (
	"code.google.com/p/gorest"
	"fmt"
	"github.com/DewaldV/chain-of-thought/services/xanatos/crucible"
	"time"
)

func Register() bool {
	serv := &LeadService{serviceConfig: crucible.Conf.Services["Leads"]}
	gorest.RegisterService(serv)
	return true
}

func (serv LeadService) DoOptions(varArgs ...string) {
	fmt.Printf("Method: %s > Request: %s\n", serv.Context.Request().Method, serv.Context.Request().URL.String())
	crucible.AddOptionsCors(serv.ResponseBuilder(), serv.Context.Request(), serv.serviceConfig)
	fmt.Printf("Received an Options request, responding with some options.\n")
}

func (serv LeadService) CreateLead(l Lead) {
	fmt.Printf("Method: %s > Request: %s\n", serv.Context.Request().Method, serv.Context.Request().URL.String())
	crucible.AddAllowOriginsCors(serv.ResponseBuilder(), serv.Context.Request(), serv.serviceConfig)
	l.RegistrationDate = time.Now()
	CreateLead(l)
	fmt.Printf("Received a create request for a lead with email: %s\n", l.Email)
	return
}

func (serv LeadService) GetLead(email string) (l Lead) {
	fmt.Printf("Method: %s > Request: %s\n", serv.Context.Request().Method, serv.Context.Request().URL.String())
	crucible.AddAllowOriginsCors(serv.ResponseBuilder(), serv.Context.Request(), serv.serviceConfig)
	l = *GetLead(email)
	fmt.Printf("Received a get request for a lead with email: %s\n", email)
	return
}

func (serv LeadService) UpdateLead(l Lead, email string) {
	fmt.Printf("Method: %s > Request: %s\n", serv.Context.Request().Method, serv.Context.Request().URL.String())
	crucible.AddAllowOriginsCors(serv.ResponseBuilder(), serv.Context.Request(), serv.serviceConfig)
	l.RegistrationDate = time.Now()
	fmt.Printf("Received a update request for a lead with email: %s\n", email)
	return
}

func (serv LeadService) DeleteLead(email string) {
	fmt.Printf("Method: %s > Request: %s\n", serv.Context.Request().Method, serv.Context.Request().URL.String())
	crucible.AddAllowOriginsCors(serv.ResponseBuilder(), serv.Context.Request(), serv.serviceConfig)
	fmt.Printf("Received a delete request for a lead with email: %s\n", email)
	return
}
