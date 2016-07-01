package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.01 Document"`
	Message *AcceptorCancellationAdviceResponseV01 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800101) AddMessage() *AcceptorCancellationAdviceResponseV01 {
	d.Message = new(AcceptorCancellationAdviceResponseV01)
	return d.Message
}

// Scope
// The AcceptorCancellationAdviceResponse message is sent by the acquirer to acknowledge the proper reception of the AcceptorCancellationAdvice. The message can be sent directly to the card acceptor or through an agent.
// Usage
// The AcceptorCancellationAdviceResponse message should be accepted by the card acceptor unless the message received was invalid.
type AcceptorCancellationAdviceResponseV01 struct {

	// Cancellation advice response message management information.
	Header *iso20022.Header2 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	AcceptorCancellationAdviceResponse *iso20022.AcceptorCancellationAdviceResponse1 `xml:"AccptrCxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationAdviceResponseV01) AddHeader() *iso20022.Header2 {
	a.Header = new(iso20022.Header2)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV01) AddAcceptorCancellationAdviceResponse() *iso20022.AcceptorCancellationAdviceResponse1 {
	a.AcceptorCancellationAdviceResponse = new(iso20022.AcceptorCancellationAdviceResponse1)
	return a.AcceptorCancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}

