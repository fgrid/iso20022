package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.01 Document"`
	Message *AcceptorDiagnosticResponseV01 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400101) AddMessage() *AcceptorDiagnosticResponseV01 {
	d.Message = new(AcceptorDiagnosticResponseV01)
	return d.Message
}

// Scope
// The AcceptorDiagnosticResponse message is sent by the acquirer to the card acceptor to confirm the availability of the acquirer. An agent never forwards the message.
// Usage
// The AcceptorDiagnosticResponse message is used to:
// - confirm the availability of the acquirer;
// - validate the security of the exchanges with the acquirer;
// - validate the version of the configuration parameters.
type AcceptorDiagnosticResponseV01 struct {

	// Diagnostic response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *iso20022.AcceptorDiagnosticResponse1 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticResponseV01) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV01) AddDiagnosticResponse() *iso20022.AcceptorDiagnosticResponse1 {
	a.DiagnosticResponse = new(iso20022.AcceptorDiagnosticResponse1)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}
