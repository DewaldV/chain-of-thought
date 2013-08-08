package leads

import (
	"code.google.com/p/gorest"
	"time"
)

func Register() bool {
	gorest.RegisterService(new(LeadService))
	return true
}

func (serv LeadService) GetLead(id int) (l Lead) {
	l = GetLead(id)
	return
}

func (serv LeadService) CreateLead(l Lead) {
	l.RegistrationDate = time.Now()
	CreateLead(l)
	return
}

func (serv LeadService) UpdateLead(l Lead, id int) {
	l.RegistrationDate = time.Now()
	l.Id = 1
	return
}

func (serv LeadService) DeleteLead(id int) {
	return
}
