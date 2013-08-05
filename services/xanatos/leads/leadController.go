package leads

import (
    "time"
   	"code.google.com/p/gorest"
)

func Register() bool {
	gorest.RegisterService(new(LeadService))
	return true
}

func(serv LeadService) GetLead(id int) Lead {
	l := GetLead(id)
	return l
}

func(serv LeadService) CreateLead(l Lead) {
	l.RegistrationDate = time.Now()
	CreateLead(l)
	return
}

func(serv LeadService) UpdateLead(l Lead, id int) {
	l.RegistrationDate = time.Now()
	l.Id = 1
	return
}

func(serv LeadService) DeleteLead(id int) {
	return
}