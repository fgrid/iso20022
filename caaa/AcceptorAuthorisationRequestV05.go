package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100105 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.05 Document"`
	Message *AcceptorAuthorisationRequestV05 `xml:"AccptrAuthstnReq"`
}

func (d *Document00100105) AddMessage() *AcceptorAuthorisationRequestV05 {
	d.Message = new(AcceptorAuthorisationRequestV05)
	return d.Message
}

// The AcceptorAuthorisationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check with the issuer (or its agent) that the account associated to the card has the resources to fund the payment. This checking will include validation of the card data and any additional transaction data provided.
type AcceptorAuthorisationRequestV05 struct {

	// Authorisation request message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the authorisation request.
	AuthorisationRequest *iso20022.AcceptorAuthorisationRequest5 `xml:"AuthstnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorAuthorisationRequestV05) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorAuthorisationRequestV05) AddAuthorisationRequest() *iso20022.AcceptorAuthorisationRequest5 {
	a.AuthorisationRequest = new(iso20022.AcceptorAuthorisationRequest5)
	return a.AuthorisationRequest
}

func (a *AcceptorAuthorisationRequestV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
