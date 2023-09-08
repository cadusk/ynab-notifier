package mail

import (
	"time"

	strftime "github.com/itchyny/timefmt-go"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type message struct {
	tos        []*mail.Email
	from       *mail.Email
	templateID string
	data       map[string]interface{}
}

func NewMessage() *message {
	m := message{}

	m.data = make(map[string]interface{})
	m.data["date"] = strftime.Format(time.Now(), "%b %d, %Y")

	return &m
}

func (m *message) AddTo(email string) {
	m.tos = append(m.tos, &mail.Email{Name: "", Address: email})
}

func (m *message) SetFrom(email string) {
	m.from = &mail.Email{Name: "", Address: email}
}

func (m *message) SetTemplateID(templateID string) {
	m.templateID = templateID
}

func (m *message) SetDynamicTemplateData(key string, data interface{}) {
	m.data[key] = data
}

func (m *message) prepareForSending() *mail.SGMailV3 {
	p := mail.NewPersonalization()
	p.AddTos(m.tos...)

	for key, value := range m.data {
		p.SetDynamicTemplateData(key, value)
	}

	return mail.NewV3Mail().
		SetFrom(m.from).
		SetTemplateID(m.templateID).
		AddPersonalizations(p)
}

func Send(token string, message *message) {
	m := message.prepareForSending()
	sendgrid.NewSendClient(token).Send(m)
}
