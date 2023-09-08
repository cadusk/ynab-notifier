package mail

import (
	"time"

	strftime "github.com/itchyny/timefmt-go"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailMessage struct {
	tos        []*mail.Email
	from       *mail.Email
	templateID string
	data       map[string]interface{}
}

func NewMailMessage() *MailMessage {
	m := MailMessage{}

	m.data = make(map[string]interface{})
	m.data["date"] = strftime.Format(time.Now(), "%b %d, %Y")

	return &m
}

func (mm *MailMessage) AddTo(email string) {
	mm.tos = append(mm.tos, &mail.Email{Name: "", Address: email})
}

func (mm *MailMessage) SetFrom(email string) {
	mm.from = &mail.Email{Name: "", Address: email}
}

func (mm *MailMessage) SetTemplateID(templateID string) {
	mm.templateID = templateID
}

func (mm *MailMessage) SetDynamicTemplateData(key string, data interface{}) {
	mm.data[key] = data
}

func (mm *MailMessage) prepareForSending() *mail.SGMailV3 {
	p := mail.NewPersonalization()
	p.AddTos(mm.tos...)

	for key, value := range mm.data {
		p.SetDynamicTemplateData(key, value)
	}

	return mail.NewV3Mail().
		SetFrom(mm.from).
		SetTemplateID(mm.templateID).
		AddPersonalizations(p)
}

func Send(token string, message *MailMessage) {
	m := message.prepareForSending()
	sendgrid.NewSendClient(token).Send(m)
}
