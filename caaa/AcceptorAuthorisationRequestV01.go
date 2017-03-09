package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.01 Document"`
	Message *AcceptorAuthorisationRequestV01 `xml:"AccptrAuthstnReq"`
}

func (d *Document00100101) AddMessage() *AcceptorAuthorisationRequestV01 {
	d.Message = new(AcceptorAuthorisationRequestV01)
	return d.Message
}

// Scope
// The AcceptorAuthorisationRequest message is sent by the card acceptor to the acquirer or its agent when an online authorisation is required for the card payment transaction.
// Usage
// The AcceptorAuthorisationRequest message may embed the information required for transferring to the acquirer the data needed to perform the financial settlement of the transaction (capture). An intermediary agent may act on behalf of the card acceptor or of the acquirer in providing additional functionality such as: switching to different acquirers, loyalty processing or management of the acceptor system.
type AcceptorAuthorisationRequestV01 struct {

	// Authorisation request message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the authorisation request.
	AuthorisationRequest *iso20022.AcceptorAuthorisationRequest1 `xml:"AuthstnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationRequestV01) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorAuthorisationRequestV01) AddAuthorisationRequest() *iso20022.AcceptorAuthorisationRequest1 {
	a.AuthorisationRequest = new(iso20022.AcceptorAuthorisationRequest1)
	return a.AuthorisationRequest
}

func (a *AcceptorAuthorisationRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}
