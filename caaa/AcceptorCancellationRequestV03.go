package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.03 Document"`
	Message *AcceptorCancellationRequestV03 `xml:"AccptrCxlReq"`
}

func (d *Document00500103) AddMessage() *AcceptorCancellationRequestV03 {
	d.Message = new(AcceptorCancellationRequestV03)
	return d.Message
}

// The AcceptorCancellationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to request the cancellation of a successfully completed transaction. Cancellation should only occur before the transaction has been cleared.
// 
// 
type AcceptorCancellationRequestV03 struct {

	// Cancellation request message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *iso20022.AcceptorCancellationRequest3 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationRequestV03) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorCancellationRequestV03) AddCancellationRequest() *iso20022.AcceptorCancellationRequest3 {
	a.CancellationRequest = new(iso20022.AcceptorCancellationRequest3)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}

