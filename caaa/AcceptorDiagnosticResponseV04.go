package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.04 Document"`
	Message *AcceptorDiagnosticResponseV04 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400104) AddMessage() *AcceptorDiagnosticResponseV04 {
	d.Message = new(AcceptorDiagnosticResponseV04)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV04 struct {

	// Diagnostic response message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *iso20022.AcceptorDiagnosticResponse4 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticResponseV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV04) AddDiagnosticResponse() *iso20022.AcceptorDiagnosticResponse4 {
	a.DiagnosticResponse = new(iso20022.AcceptorDiagnosticResponse4)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}
