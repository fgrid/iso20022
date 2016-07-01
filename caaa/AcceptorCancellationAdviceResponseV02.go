package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.02 Document"`
	Message *AcceptorCancellationAdviceResponseV02 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800102) AddMessage() *AcceptorCancellationAdviceResponseV02 {
	d.Message = new(AcceptorCancellationAdviceResponseV02)
	return d.Message
}

// The AcceptorCancellationAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) about the notification of the payment cancellation.
type AcceptorCancellationAdviceResponseV02 struct {

	// Cancellation advice response message management information.
	Header *iso20022.Header2 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	CancellationAdviceResponse *iso20022.AcceptorCancellationAdviceResponse2 `xml:"CxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationAdviceResponseV02) AddHeader() *iso20022.Header2 {
	a.Header = new(iso20022.Header2)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV02) AddCancellationAdviceResponse() *iso20022.AcceptorCancellationAdviceResponse2 {
	a.CancellationAdviceResponse = new(iso20022.AcceptorCancellationAdviceResponse2)
	return a.CancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}

