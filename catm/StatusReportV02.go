package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100102 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.02 Document"`
	Message *StatusReportV02 `xml:"StsRpt"`
}

func (d *Document00100102) AddMessage() *StatusReportV02 {
	d.Message = new(StatusReportV02)
	return d.Message
}

// Informs the master terminal manager (MTM) or the terminal manager (TM) about the status of the acceptor system including the identification of the POI, its components and their installed versions.
type StatusReportV02 struct {

	// Set of characteristics related to the transfer of the status report.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Status of the point of interaction (POI), its components and their installed versions.
	StatusReport *iso20022.StatusReport2 `xml:"StsRpt"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType4 `xml:"SctyTrlr"`
}

func (s *StatusReportV02) AddHeader() *iso20022.Header4 {
	s.Header = new(iso20022.Header4)
	return s.Header
}

func (s *StatusReportV02) AddStatusReport() *iso20022.StatusReport2 {
	s.StatusReport = new(iso20022.StatusReport2)
	return s.StatusReport
}

func (s *StatusReportV02) AddSecurityTrailer() *iso20022.ContentInformationType4 {
	s.SecurityTrailer = new(iso20022.ContentInformationType4)
	return s.SecurityTrailer
}
