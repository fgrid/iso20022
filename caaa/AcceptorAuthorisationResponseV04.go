package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Document"`
	Message *AcceptorAuthorisationResponseV04 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200104) AddMessage() *AcceptorAuthorisationResponseV04 {
	d.Message = new(AcceptorAuthorisationResponseV04)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV04 struct {

	// Authorisation response message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *iso20022.AcceptorAuthorisationResponse4 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`

}


func (a *AcceptorAuthorisationResponseV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV04) AddAuthorisationResponse() *iso20022.AcceptorAuthorisationResponse4 {
	a.AuthorisationResponse = new(iso20022.AcceptorAuthorisationResponse4)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}

