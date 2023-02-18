package render

import (
	"certie/internal"
	"encoding/xml"
	"fmt"
	"os"
)

func RendererXML(data []*internal.TlsCertificate) string {
	xmlData, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Could not render XL data: %s", err)
		os.Exit(1)
	}
	return string(xmlData)
}
