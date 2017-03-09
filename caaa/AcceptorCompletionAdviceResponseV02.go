package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.02 Document"`
	Message *AcceptorCompletionAdviceResponseV02 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400102) AddMessage() *AcceptorCompletionAdviceResponseV02 {
	d.Message = new(AcceptorCompletionAdviceResponseV02)
	return d.Message
}

// The AcceptorCompletionAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) of the outcome of the payment transaction, and the transfer the  financial data of the transaction contained in the completion advice.
type AcceptorCompletionAdviceResponseV02 struct {

	// Completion advice response message management information.
	Header *iso20022.Header2 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *iso20022.AcceptorCompletionAdviceResponse2 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceResponseV02) AddHeader() *iso20022.Header2 {
	a.Header = new(iso20022.Header2)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV02) AddCompletionAdviceResponse() *iso20022.AcceptorCompletionAdviceResponse2 {
	a.CompletionAdviceResponse = new(iso20022.AcceptorCompletionAdviceResponse2)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
