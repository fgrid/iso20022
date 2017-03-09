package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.003.001.03 Document"`
	Message *AcceptorCompletionAdviceV03 `xml:"AccptrCmpltnAdvc"`
}

func (d *Document00300103) AddMessage() *AcceptorCompletionAdviceV03 {
	d.Message = new(AcceptorCompletionAdviceV03)
	return d.Message
}

// The AcceptorCompletionAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the outcome of the payment at the acceptor, and to transfer the  financial data of the transaction to the acquirer (capture).
// A AcceptorCompletionAdvice message is also sent to reverse an approved authorisation and any associated financial transfer (capture), if the card payment transaction could not be completed successfully.
type AcceptorCompletionAdviceV03 struct {

	// Completion advice message management information.
	Header *iso20022.Header8 `xml:"Hdr"`

	// Information related to the completion advice.
	CompletionAdvice *iso20022.AcceptorCompletionAdvice3 `xml:"CmpltnAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceV03) AddHeader() *iso20022.Header8 {
	a.Header = new(iso20022.Header8)
	return a.Header
}

func (a *AcceptorCompletionAdviceV03) AddCompletionAdvice() *iso20022.AcceptorCompletionAdvice3 {
	a.CompletionAdvice = new(iso20022.AcceptorCompletionAdvice3)
	return a.CompletionAdvice
}

func (a *AcceptorCompletionAdviceV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}
