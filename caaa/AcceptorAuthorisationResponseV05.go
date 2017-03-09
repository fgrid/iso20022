package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200105 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.05 Document"`
	Message *AcceptorAuthorisationResponseV05 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200105) AddMessage() *AcceptorAuthorisationResponseV05 {
	d.Message = new(AcceptorAuthorisationResponseV05)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV05 struct {

	// Authorisation response message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *iso20022.AcceptorAuthorisationResponse5 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorAuthorisationResponseV05) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV05) AddAuthorisationResponse() *iso20022.AcceptorAuthorisationResponse5 {
	a.AuthorisationResponse = new(iso20022.AcceptorAuthorisationResponse5)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
