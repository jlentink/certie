package internal

import (
	"crypto/x509"
	"fmt"
	"time"
)

type TlsCertificate struct {
	certificate *x509.Certificate
	CommonName  string    `json:"commonName"`
	Sans        []string  `json:"sans"`
	Issuer      string    `json:"issuer"`
	NotBefore   time.Time `json:"notBefore"`
	NotAfter    time.Time `json:"notAfter"`
	IsExpired   bool      `json:"isExpired"`
	ExpiredText string    `json:"expiredText"`
	host        TargetHost
}

func (t *TlsCertificate) getLocaleTime() *time.Location {
	return time.Local
}

func (t *TlsCertificate) addCertificate(certificate *x509.Certificate) {
	t.certificate = certificate
	t.SetSans()
	t.SetNotBefore()
	t.SetNotAfter()
	t.SetIssuer()
	t.SetCommonName()
	t.SetIsExpired()
	t.SetIsExpiredString()
}

func (t *TlsCertificate) SetSans() {
	t.Sans = t.certificate.DNSNames
}

func (t *TlsCertificate) SetNotBefore() {
	t.NotBefore = t.certificate.NotBefore.In(t.getLocaleTime())
}

func (t *TlsCertificate) SetNotAfter() {
	t.NotAfter = t.certificate.NotAfter.In(t.getLocaleTime())
}

func (t *TlsCertificate) SetIssuer() {
	t.Issuer = t.certificate.Issuer.CommonName
}

func (t *TlsCertificate) SetCommonName() {
	t.CommonName = t.certificate.Subject.CommonName
}

func (t *TlsCertificate) SetIsExpired() {
	diff := t.NotAfter.Sub(time.Now())
	t.IsExpired = diff < 0
}

func (t *TlsCertificate) SetIsExpiredString() {
	diff := t.NotAfter.Sub(time.Now())
	days := diff / (24 * time.Hour)
	if days < 0 {
		t.ExpiredText = fmt.Sprintf("Expired %d days ago", days*-1)
	}
	t.ExpiredText = fmt.Sprintf("Valid for %d days", days)
}

func (t *TlsCertificate) GetHost() TargetHost {
	return t.host
}

func (t *TlsCertificate) Summary(prefix bool) string {
	output := ""
	if prefix {
		output += fmt.Sprintf("-------------------------------------------------------")
	}
	output += fmt.Sprintf("Remote Addr: %s\n", t.host.RemoteAddr)
	output += fmt.Sprintf("Common:      %s\n", t.CommonName)
	output += fmt.Sprintf("SANs :       %s\n", t.Sans)
	output += fmt.Sprintf("Issuer:      %s\n", t.Issuer)
	output += fmt.Sprintf("Not before:  %s\n", t.NotBefore.String())
	output += fmt.Sprintf("Not After:   %s\n", t.NotAfter.String())
	output += fmt.Sprintf("Expired:     %s\n", t.ExpiredText)
	return output
}
