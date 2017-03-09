package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300105 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.003.001.05 Document"`
	Message *AcceptorCompletionAdviceV05 `xml:"AccptrCmpltnAdvc"`
}

func (d *Document00300105) AddMessage() *AcceptorCompletionAdviceV05 {
	d.Message = new(AcceptorCompletionAdviceV05)
	return d.Message
}

// The AcceptorCompletionAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the outcome of the payment at the acceptor, and to transfer the  financial data of the transaction to the acquirer (capture).
// A AcceptorCompletionAdvice message is also sent to reverse an approved authorisation and any associated financial transfer (capture), if the card payment transaction could not be completed successfully.
type AcceptorCompletionAdviceV05 struct {

	// Completion advice message management information.
	Header *iso20022.Header24 `xml:"Hdr"`

	// Information related to the completion advice.
	CompletionAdvice *iso20022.AcceptorCompletionAdvice5 `xml:"CmpltnAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCompletionAdviceV05) AddHeader() *iso20022.Header24 {
	a.Header = new(iso20022.Header24)
	return a.Header
}

func (a *AcceptorCompletionAdviceV05) AddCompletionAdvice() *iso20022.AcceptorCompletionAdvice5 {
	a.CompletionAdvice = new(iso20022.AcceptorCompletionAdvice5)
	return a.CompletionAdvice
}

func (a *AcceptorCompletionAdviceV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
