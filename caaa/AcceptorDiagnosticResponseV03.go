package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Document"`
	Message *AcceptorDiagnosticResponseV03 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400103) AddMessage() *AcceptorDiagnosticResponseV03 {
	d.Message = new(AcceptorDiagnosticResponseV03)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV03 struct {

	// Diagnostic response message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *iso20022.AcceptorDiagnosticResponse3 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`

}


func (a *AcceptorDiagnosticResponseV03) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV03) AddDiagnosticResponse() *iso20022.AcceptorDiagnosticResponse3 {
	a.DiagnosticResponse = new(iso20022.AcceptorDiagnosticResponse3)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}

