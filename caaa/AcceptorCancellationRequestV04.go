package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.04 Document"`
	Message *AcceptorCancellationRequestV04 `xml:"AccptrCxlReq"`
}

func (d *Document00500104) AddMessage() *AcceptorCancellationRequestV04 {
	d.Message = new(AcceptorCancellationRequestV04)
	return d.Message
}

// The AcceptorCancellationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to request the cancellation of a successfully completed transaction. Cancellation should only occur before the transaction has been cleared.
// 
// 
type AcceptorCancellationRequestV04 struct {

	// Cancellation request message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *iso20022.AcceptorCancellationRequest4 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationRequestV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorCancellationRequestV04) AddCancellationRequest() *iso20022.AcceptorCancellationRequest4 {
	a.CancellationRequest = new(iso20022.AcceptorCancellationRequest4)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}

