package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.01 Document"`
	Message *AcceptorAuthorisationResponseV01 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200101) AddMessage() *AcceptorAuthorisationResponseV01 {
	d.Message = new(AcceptorAuthorisationResponseV01)
	return d.Message
}

// Scope
// The AcceptorAuthorisationResponse message is sent by the acquirer to inform the card acceptor of the outcome of the authorisation process. The message can be sent directly to the acceptor or through an agent.
// Usage
// The AcceptorAuthorisationResponse message is used to indicate one of the possible outcomes of an authorisation process:
// - a successful authorisation;
// - a decline from the acquirer for financial reasons;
// - a decline from the acquirer for technical reasons (for instance, a timeout).
type AcceptorAuthorisationResponseV01 struct {

	// Authorisation response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *iso20022.AcceptorAuthorisationResponse1 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`

}


func (a *AcceptorAuthorisationResponseV01) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV01) AddAuthorisationResponse() *iso20022.AcceptorAuthorisationResponse1 {
	a.AuthorisationResponse = new(iso20022.AcceptorAuthorisationResponse1)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}

