package notify

import "github.com/NawafSwe/gomailer"

type Notify interface {
	Send(message Message) error
}
type Router struct {
	*gomailer.Mailer
}
