package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.02 Document"`
	Message *AcceptorDiagnosticResponseV02 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400102) AddMessage() *AcceptorDiagnosticResponseV02 {
	d.Message = new(AcceptorDiagnosticResponseV02)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV02 struct {

	// Diagnostic response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *iso20022.AcceptorDiagnosticResponse2 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticResponseV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV02) AddDiagnosticResponse() *iso20022.AcceptorDiagnosticResponse2 {
	a.DiagnosticResponse = new(iso20022.AcceptorDiagnosticResponse2)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
