package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800104 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.04 Document"`
	Message *AcceptorCancellationAdviceResponseV04 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800104) AddMessage() *AcceptorCancellationAdviceResponseV04 {
	d.Message = new(AcceptorCancellationAdviceResponseV04)
	return d.Message
}

// The AcceptorCancellationAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) about the notification of the payment cancellation.
type AcceptorCancellationAdviceResponseV04 struct {

	// Cancellation advice response message management information.
	Header *iso20022.Header11 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	CancellationAdviceResponse *iso20022.AcceptorCancellationAdviceResponse4 `xml:"CxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceResponseV04) AddHeader() *iso20022.Header11 {
	a.Header = new(iso20022.Header11)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV04) AddCancellationAdviceResponse() *iso20022.AcceptorCancellationAdviceResponse4 {
	a.CancellationAdviceResponse = new(iso20022.AcceptorCancellationAdviceResponse4)
	return a.CancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}
