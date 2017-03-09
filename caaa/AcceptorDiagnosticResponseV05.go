package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400105 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.05 Document"`
	Message *AcceptorDiagnosticResponseV05 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400105) AddMessage() *AcceptorDiagnosticResponseV05 {
	d.Message = new(AcceptorDiagnosticResponseV05)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV05 struct {

	// Diagnostic response message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *iso20022.AcceptorDiagnosticResponse4 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorDiagnosticResponseV05) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV05) AddDiagnosticResponse() *iso20022.AcceptorDiagnosticResponse4 {
	a.DiagnosticResponse = new(iso20022.AcceptorDiagnosticResponse4)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
