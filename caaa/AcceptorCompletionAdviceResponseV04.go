package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.04 Document"`
	Message *AcceptorCompletionAdviceResponseV04 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400104) AddMessage() *AcceptorCompletionAdviceResponseV04 {
	d.Message = new(AcceptorCompletionAdviceResponseV04)
	return d.Message
}

// The AcceptorCompletionAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) of the outcome of the payment transaction, and the transfer the  financial data of the transaction contained in the completion advice.
type AcceptorCompletionAdviceResponseV04 struct {

	// Completion advice response message management information.
	Header *iso20022.Header11 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *iso20022.AcceptorCompletionAdviceResponse4 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceResponseV04) AddHeader() *iso20022.Header11 {
	a.Header = new(iso20022.Header11)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV04) AddCompletionAdviceResponse() *iso20022.AcceptorCompletionAdviceResponse4 {
	a.CompletionAdviceResponse = new(iso20022.AcceptorCompletionAdviceResponse4)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}
