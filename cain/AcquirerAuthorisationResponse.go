package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:cain.002.001.01 Document"`
	Message *AcquirerAuthorisationResponse `xml:"AcqrrAuthstnRspn"`
}

func (d *Document00200101) AddMessage() *AcquirerAuthorisationResponse {
	d.Message = new(AcquirerAuthorisationResponse)
	return d.Message
}

// The AcquirerAuthorisationResponse message is sent by an issuer or an agent to answer to an AcquirerAuthorisationInitiation message.
type AcquirerAuthorisationResponse struct {

	// Information related to the protocol management.
	Header *iso20022.Header17 `xml:"Hdr"`

	// Information related to the response of an authorisation.
	AuthorisationResponse *iso20022.AcquirerAuthorisationResponse1 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcquirerAuthorisationResponse) AddHeader() *iso20022.Header17 {
	a.Header = new(iso20022.Header17)
	return a.Header
}

func (a *AcquirerAuthorisationResponse) AddAuthorisationResponse() *iso20022.AcquirerAuthorisationResponse1 {
	a.AuthorisationResponse = new(iso20022.AcquirerAuthorisationResponse1)
	return a.AuthorisationResponse
}

func (a *AcquirerAuthorisationResponse) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
