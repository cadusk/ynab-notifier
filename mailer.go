package main

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func sendMail(s *Summary) {
	tos := []*mail.Email{}
	for _, email := range config.Sendgrid.MailTo {
		tos = append(tos, mail.NewEmail("", email))
	}

	p := mail.NewPersonalization()
	p.AddTos(tos...)
	p.SetDynamicTemplateData("date", s.Date)
	p.SetDynamicTemplateData("favorites", s.Favorites)
	p.SetDynamicTemplateData("red_flags", s.RedFlags)
	p.SetDynamicTemplateData("goals", s.Goals)

	m := mail.NewV3Mail().
		SetFrom(mail.NewEmail(config.Sendgrid.MailFromName, config.Sendgrid.MailFrom)).
		SetTemplateID(config.Sendgrid.TemplateID).
		AddPersonalizations(p)

	sendgrid.NewSendClient(config.Sendgrid.AccessToken).Send(m)
}
