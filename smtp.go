package main

import (
	"io"
	"log"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
)

var ValidUsernames = []string{"accept", "softfail", "hardfail", "later"}

type Backend struct{}
type Session struct {
	MailFrom string
	Username string
}

func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

func (s *Session) AuthPlain(username, password string) error {
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("smtp: mail from:", from)
	s.MailFrom = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("smtp: rcpt to:", to)

	username := getUsername(to)
	if !contains(ValidUsernames, username) {
		log.Println("smtp: 550 invalid recipient address")
		return &smtp.SMTPError{
			Code:    550,
			Message: "Invalid recipient address",
		}
	}
	s.Username = getUsername(to)
	return nil
}

func (s *Session) Data(r io.Reader) error {

	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("smtp: data:", string(b))
	}

	switch s.Username {
	case "accept":
		log.Println("smtp: accepting message")
		return nil
	case "softfail":
		log.Println("smtp: 450 mailbox is available")
		return &smtp.SMTPError{
			Code:    450,
			Message: "Mailbox is unavailable at the moment",
		}
	case "later":
		log.Println("smtp: 450 try again later")
		return &smtp.SMTPError{
			Code:    450,
			Message: "Try again in 250 seconds",
		}
	default:
		log.Println("smtp: 550 invalid recipient address")
		return &smtp.SMTPError{
			Code:    550,
			Message: "Invalid recipient address",
		}
	}
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func runSMTPServer() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":2525"
	s.Domain = "blackhole"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("smtp: starting smtp server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getUsername(s string) string {
	parts := strings.Split(s, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
