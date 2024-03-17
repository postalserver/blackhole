package main

import (
	"log"
	"net"
	"strings"

	"github.com/mhale/smtpd"
)

func runSMTPServer() {
	srv := &smtpd.Server{
		Addr:         "127.0.0.1:2525",
		Handler:      mailHandler,
		HandlerRcpt:  rcptHandler,
		Appname:      "blackhole",
		Hostname:     "blackhole",
		AuthHandler:  nil,
		AuthRequired: false,
	}
	log.Printf("smtp: starting smtp server")
	srv.ListenAndServe()
}

func getUsername(s string) string {
	parts := strings.Split(s, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func rcptHandler(ip net.Addr, from string, to string) bool {
	username := getUsername(to)
	if username == "accept" {
		log.Printf("smtp: accepting email to %s from %s (%s)", to, from, ip)
		return true
	}

	log.Printf("smtp: rejecting email to %s from %s (%s)", to, from, ip)
	return false
}

func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	log.Printf("smtp data: %s", data)
	return nil
}
