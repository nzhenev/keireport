package exporter

import (
	"github.com/adlindo/keireport/core"
	"github.com/adlindo/keireport/exporter/pdf"
)

func init() {

	core.RegisterExporter("pdf", &pdf.PDFExporter{})
}
