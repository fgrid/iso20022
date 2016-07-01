package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.03 Document"`
	Message *AcceptorCancellationAdviceResponseV03 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800103) AddMessage() *AcceptorCancellationAdviceResponseV03 {
	d.Message = new(AcceptorCancellationAdviceResponseV03)
	return d.Message
}

// The AcceptorCancellationAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) about the notification of the payment cancellation.
type AcceptorCancellationAdviceResponseV03 struct {

	// Cancellation advice response message management information.
	Header *iso20022.Header8 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	CancellationAdviceResponse *iso20022.AcceptorCancellationAdviceResponse3 `xml:"CxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationAdviceResponseV03) AddHeader() *iso20022.Header8 {
	a.Header = new(iso20022.Header8)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV03) AddCancellationAdviceResponse() *iso20022.AcceptorCancellationAdviceResponse3 {
	a.CancellationAdviceResponse = new(iso20022.AcceptorCancellationAdviceResponse3)
	return a.CancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}

