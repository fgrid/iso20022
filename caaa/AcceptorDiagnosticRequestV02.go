package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.02 Document"`
	Message *AcceptorDiagnosticRequestV02 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300102) AddMessage() *AcceptorDiagnosticRequestV02 {
	d.Message = new(AcceptorDiagnosticRequestV02)
	return d.Message
}

// The AcceptorDiagnosticRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check the end-to-end communication, to test the availability of this acquirer, or to validate the security environment.
type AcceptorDiagnosticRequestV02 struct {

	// Diagnostic request message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *iso20022.AcceptorDiagnosticRequest2 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticRequestV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV02) AddDiagnosticRequest() *iso20022.AcceptorDiagnosticRequest2 {
	a.DiagnosticRequest = new(iso20022.AcceptorDiagnosticRequest2)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
