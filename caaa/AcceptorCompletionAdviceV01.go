package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.003.001.01 Document"`
	Message *AcceptorCompletionAdviceV01 `xml:"AccptrCmpltnAdvc"`
}

func (d *Document00300101) AddMessage() *AcceptorCompletionAdviceV01 {
	d.Message = new(AcceptorCompletionAdviceV01)
	return d.Message
}

// Scope
// The AcceptorCompletionAdvice message is sent by a card acceptor to notify an acquirer about the completion and final outcome of a card payment transaction. The message can be sent directly to the acquirer or through an agent.
// Usage
// The AcceptorCompletionAdvice message is used either to:
// - inform the acquirer about the successful end of a transaction;
// - reverse a transaction which was not successfully completed (for example, cancellation of transaction by the cardholder), but where an authorisation had been previously given.
// The AcceptorCompletionAdvice message may also embed the information required for transferring to the acquirer all data needed to perform the financial settlement of the transaction (capture). Should the acquirer not receive a correct response to an AcceptorCompletionAdvice message; the card acceptor sends back an AcceptorCompletionAdvice message to the acquirer.
type AcceptorCompletionAdviceV01 struct {

	// Completion advice message management information.
	Header *iso20022.Header2 `xml:"Hdr"`

	// Information related to the completion advice.
	CompletionAdvice *iso20022.AcceptorCompletionAdvice1 `xml:"CmpltnAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`

}


func (a *AcceptorCompletionAdviceV01) AddHeader() *iso20022.Header2 {
	a.Header = new(iso20022.Header2)
	return a.Header
}

func (a *AcceptorCompletionAdviceV01) AddCompletionAdvice() *iso20022.AcceptorCompletionAdvice1 {
	a.CompletionAdvice = new(iso20022.AcceptorCompletionAdvice1)
	return a.CompletionAdvice
}

func (a *AcceptorCompletionAdviceV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}

