package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100104 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.04 Document"`
	Message *AcceptorAuthorisationRequestV04 `xml:"AccptrAuthstnReq"`
}

func (d *Document00100104) AddMessage() *AcceptorAuthorisationRequestV04 {
	d.Message = new(AcceptorAuthorisationRequestV04)
	return d.Message
}

// The AcceptorAuthorisationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check with the issuer (or its agent) that the account associated to the card has the resources to fund the payment. This checking will include validation of the card data and any additional transaction data provided.
type AcceptorAuthorisationRequestV04 struct {

	// Authorisation request message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the authorisation request.
	AuthorisationRequest *iso20022.AcceptorAuthorisationRequest4 `xml:"AuthstnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationRequestV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorAuthorisationRequestV04) AddAuthorisationRequest() *iso20022.AcceptorAuthorisationRequest4 {
	a.AuthorisationRequest = new(iso20022.AcceptorAuthorisationRequest4)
	return a.AuthorisationRequest
}

func (a *AcceptorAuthorisationRequestV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}
