package render

import (
	"certie/internal"
	"fmt"
)

func RendererText(data []*internal.TlsCertificate) string {
	output := ""
	for c, cert := range data {
		if c > 0 {
			output += "-----------------------------------------------------------------\n"
		}
		output += fmt.Sprintf("Remote Addr: %s\n", cert.GetHost().RemoteAddr)
		output += fmt.Sprintf("Common:      %s\n", cert.CommonName)
		output += fmt.Sprintf("SANs :       %s\n", cert.Sans)
		output += fmt.Sprintf("Issuer:      %s\n", cert.Issuer)
		output += fmt.Sprintf("Not before:  %s\n", cert.NotBefore.String())
		output += fmt.Sprintf("Not After:   %s\n", cert.NotAfter.String())
		output += fmt.Sprintf("Expired:     %s\n", cert.ExpiredText)

	}
	return output
}
