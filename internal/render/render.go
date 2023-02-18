package render

import "certie/internal"

func Render(renderType string, data []*internal.TlsCertificate) string {
	switch renderType {
	case "json":
		return RendererJSON(data)
	case "yaml":
		return RendererYAML(data)
	case "xml":
		return RendererXML(data)
	default:
		return RendererText(data)
	}
}
