package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.03 Document"`
	Message *AcceptorAuthorisationResponseV03 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200103) AddMessage() *AcceptorAuthorisationResponseV03 {
	d.Message = new(AcceptorAuthorisationResponseV03)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV03 struct {

	// Authorisation response message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *iso20022.AcceptorAuthorisationResponse3 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`

}


func (a *AcceptorAuthorisationResponseV03) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV03) AddAuthorisationResponse() *iso20022.AcceptorAuthorisationResponse3 {
	a.AuthorisationResponse = new(iso20022.AcceptorAuthorisationResponse3)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}

