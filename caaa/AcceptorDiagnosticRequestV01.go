package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.01 Document"`
	Message *AcceptorDiagnosticRequestV01 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300101) AddMessage() *AcceptorDiagnosticRequestV01 {
	d.Message = new(AcceptorDiagnosticRequestV01)
	return d.Message
}

// Scope
// The AcceptorDiagnosticRequest message is sent by the card acceptor to the acquirer to ensure the availability of the acquirer. An agent never forwards the message.
// Usage
// The AcceptorDiagnosticRequest message is used to:
// - test the availability of the acquirer;
// - validate the security of the exchanges with the acquirer;
// - validate the version of the configuration parameters.
type AcceptorDiagnosticRequestV01 struct {

	// Diagnostic request message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *iso20022.AcceptorDiagnosticRequest1 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`

}


func (a *AcceptorDiagnosticRequestV01) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV01) AddDiagnosticRequest() *iso20022.AcceptorDiagnosticRequest1 {
	a.DiagnosticRequest = new(iso20022.AcceptorDiagnosticRequest1)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}

