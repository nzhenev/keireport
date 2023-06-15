package exporter

import (
	"github.com/adlindo/keireport/core"
	"github.com/adlindo/keireport/exporter/html"
)

func init() {

	core.RegisterExporter("html", &html.HTMLExporter{})
}
