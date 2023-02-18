package internal

import "crypto/tls"

var cipherSuites map[string]uint16

func init() {
	cipherSuites = make(map[string]uint16)
	for _, suite := range tls.CipherSuites() {
		cipherSuites[suite.Name] = suite.ID
	}
}

func getCipherSuites() []uint16 {
	var suites []uint16
	for _, suite := range tls.CipherSuites() {
		suites = append(suites, suite.ID)
	}
	return suites
}
