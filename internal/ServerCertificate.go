package internal

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func GetCertificates(host TargetHost) *TlsCertificates {
	d := &net.Dialer{
		Timeout: time.Duration(host.Timeout) * time.Second,
	}
	conn, err := tls.DialWithDialer(d, "tcp", host.GetDailUrl(), &tls.Config{
		InsecureSkipVerify: true,
		CipherSuites:       getCipherSuites(),
		MaxVersion:         tls.VersionTLS12,
	})

	if err != nil {
		fmt.Printf("Cannot connect to host (%s): %s\n", host.GetDailUrl(), err)
		os.Exit(1)
	}

	x509certs := conn.ConnectionState().PeerCertificates
	host.RemoteAddr = conn.RemoteAddr().String()
	err = conn.Close()
	if err != nil {
		fmt.Printf("Error closing connection: %s\n", err)
		os.Exit(1)
	}
	certs := TlsCertificates{}
	certs.addCertificates(x509certs, host)

	return &certs
}
