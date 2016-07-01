package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.02 Document"`
	Message *AcceptorCancellationRequestV02 `xml:"AccptrCxlReq"`
}

func (d *Document00500102) AddMessage() *AcceptorCancellationRequestV02 {
	d.Message = new(AcceptorCancellationRequestV02)
	return d.Message
}

// The AcceptorCancellationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to request the cancellation of a successfully completed transaction. Cancellation should only occur before the transaction has been cleared.
// 
// 
type AcceptorCancellationRequestV02 struct {

	// Cancellation request message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *iso20022.AcceptorCancellationRequest2 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationRequestV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorCancellationRequestV02) AddCancellationRequest() *iso20022.AcceptorCancellationRequest2 {
	a.CancellationRequest = new(iso20022.AcceptorCancellationRequest2)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}

