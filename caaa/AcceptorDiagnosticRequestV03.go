package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.03 Document"`
	Message *AcceptorDiagnosticRequestV03 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300103) AddMessage() *AcceptorDiagnosticRequestV03 {
	d.Message = new(AcceptorDiagnosticRequestV03)
	return d.Message
}

// The AcceptorDiagnosticRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check the end-to-end communication, to test the availability of this acquirer, or to validate the security environment.
type AcceptorDiagnosticRequestV03 struct {

	// Diagnostic request message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *iso20022.AcceptorDiagnosticRequest3 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticRequestV03) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV03) AddDiagnosticRequest() *iso20022.AcceptorDiagnosticRequest3 {
	a.DiagnosticRequest = new(iso20022.AcceptorDiagnosticRequest3)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}
