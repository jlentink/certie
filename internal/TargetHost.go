package internal

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TargetHost struct {
	Host       string
	Port       int64
	Timeout    int64
	RemoteAddr string
}

func (s *TargetHost) ParseUrl(u string) {
	httpRegex := regexp.MustCompile(`(?i)^http(s?)://`)
	if httpRegex.MatchString(u) {
		s.parseHTTPString(u)
	} else {
		s.parseHostString(u)
	}

	if s.Port == 0 {
		s.SetPort(443)
	}
}

func (s *TargetHost) parseHostString(u string) {
	if strings.Contains(u, ":") {
		parts := strings.Split(u, ":")
		if len(parts) == 2 {
			s.SetHost(parts[0])
			s.SetStringPort(parts[1])
			return
		}
	} else {
		s.SetHost(u)
	}
}

func (s *TargetHost) parseHTTPString(u string) {
	uo, err := url.Parse(u)
	if err != nil {
		fmt.Printf("Error parsing URL: %s", err)
		os.Exit(1)
	}
	if uo.Port() != "" {
		s.SetStringPort(uo.Port())
	}
	s.SetHost(uo.Host)
}

func (s *TargetHost) SetStringPort(p string) {
	port, err := strconv.ParseInt(p, 10, 64)
	if err == nil {
		s.SetPort(port)
	}
}

func (s *TargetHost) SetHost(h string) {
	s.Host = h
}

func (s *TargetHost) SetPort(p int64) {
	if p > 0 && s.Port == 0 {
		s.Port = p
	}
}

func (s *TargetHost) GetHost() string {
	return s.Host
}

func (s *TargetHost) GetPort() int64 {
	return s.Port
}

func (s *TargetHost) GetDailUrl() string {
	return fmt.Sprintf("%s:%d", s.GetHost(), s.GetPort())
}

func (s *TargetHost) SetTimeout(t int64) {
	s.Timeout = t
}

func (s *TargetHost) GetIP() string {
	return s.RemoteAddr
}
