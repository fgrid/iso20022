package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.04 Document"`
	Message *AcceptorDiagnosticRequestV04 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300104) AddMessage() *AcceptorDiagnosticRequestV04 {
	d.Message = new(AcceptorDiagnosticRequestV04)
	return d.Message
}

// The AcceptorDiagnosticRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check the end-to-end communication, to test the availability of this acquirer, or to validate the security environment.
type AcceptorDiagnosticRequestV04 struct {

	// Diagnostic request message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *iso20022.AcceptorDiagnosticRequest4 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`

}


func (a *AcceptorDiagnosticRequestV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV04) AddDiagnosticRequest() *iso20022.AcceptorDiagnosticRequest4 {
	a.DiagnosticRequest = new(iso20022.AcceptorDiagnosticRequest4)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}

