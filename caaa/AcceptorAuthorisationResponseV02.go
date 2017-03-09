package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.02 Document"`
	Message *AcceptorAuthorisationResponseV02 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200102) AddMessage() *AcceptorAuthorisationResponseV02 {
	d.Message = new(AcceptorAuthorisationResponseV02)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV02 struct {

	// Authorisation response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *iso20022.AcceptorAuthorisationResponse2 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationResponseV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV02) AddAuthorisationResponse() *iso20022.AcceptorAuthorisationResponse2 {
	a.AuthorisationResponse = new(iso20022.AcceptorAuthorisationResponse2)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
