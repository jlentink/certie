package render

import (
	"certie/internal"
	"encoding/json"
	"fmt"
	"os"
)

func RendererJSON(data []*internal.TlsCertificate) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Could not render JSON data: %s", err)
		os.Exit(1)
	}
	return string(jsonData)
}
