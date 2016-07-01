package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.03 Document"`
	Message *StatusReportV03 `xml:"StsRpt"`
}

func (d *Document00100103) AddMessage() *StatusReportV03 {
	d.Message = new(StatusReportV03)
	return d.Message
}

// Informs the master terminal manager (MTM) or the terminal manager (TM) about the status of the acceptor system including the identification of the POI, its components and their installed versions.
type StatusReportV03 struct {

	// Set of characteristics related to the transfer of the status report.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Status of the point of interaction (POI), its components and their installed versions.
	StatusReport *iso20022.StatusReport3 `xml:"StsRpt"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType9 `xml:"SctyTrlr"`

}


func (s *StatusReportV03) AddHeader() *iso20022.Header4 {
	s.Header = new(iso20022.Header4)
	return s.Header
}

func (s *StatusReportV03) AddStatusReport() *iso20022.StatusReport3 {
	s.StatusReport = new(iso20022.StatusReport3)
	return s.StatusReport
}

func (s *StatusReportV03) AddSecurityTrailer() *iso20022.ContentInformationType9 {
	s.SecurityTrailer = new(iso20022.ContentInformationType9)
	return s.SecurityTrailer
}

