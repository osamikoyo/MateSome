package notify

import (
	"crypto/tls"
	"github.com/NawafSwe/gomailer"
	"time"
)

type Notify interface {
	SendNotify(message Message) error
}
type Router struct {
	*gomailer.Mailer
}

func New(host, username, password string) *Router {
	email := gomailer.NewMailer(
		host,
		587,
		username,
		password,
		gomailer.WithLocalName("localhost"),
		gomailer.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		gomailer.WithDialTimeout(2*time.Second),
	)

	return &Router{
		email,
	}
}

func (r *Router) SendNotify(msg Message) error {
	return r.Mailer.Send(msg.message)
}
