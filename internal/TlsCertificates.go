package internal

import "crypto/x509"

type TlsCertificates struct {
	TlsCertificate []*TlsCertificate
}

func (t *TlsCertificates) addCertificates(certs []*x509.Certificate, host TargetHost) {
	for _, cert := range certs {
		tlsCert := &TlsCertificate{host: host}
		tlsCert.addCertificate(cert)
		t.TlsCertificate = append(t.TlsCertificate, tlsCert)
	}
}

func (t *TlsCertificates) GetCertificates(includeChain bool) []*TlsCertificate {
	if includeChain {
		return t.TlsCertificate
	}
	return []*TlsCertificate{t.TlsCertificate[0]}
}

func (t *TlsCertificates) GetHostCertificate() *TlsCertificate {
	return t.TlsCertificate[0]
}

func (t *TlsCertificates) Summary() string {
	return t.GetHostCertificate().Summary(false)
}

func (t *TlsCertificates) SummaryAll() string {
	return t.GetHostCertificate().Summary(false)
}

func (t *TlsCertificates) GetAllCertificates() []*TlsCertificate {
	return t.TlsCertificate
}

func (t *TlsCertificates) GetIntermediateCertificates() []*TlsCertificate {
	return t.TlsCertificate[1:]
}
