package render

import (
	"certie/internal"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func RendererYAML(data []*internal.TlsCertificate) string {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("Could not render YAML data: %s", err)
		os.Exit(1)
	}
	return string(yamlData)
}
