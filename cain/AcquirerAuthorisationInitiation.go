package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:cain.001.001.01 Document"`
	Message *AcquirerAuthorisationInitiation `xml:"AcqrrAuthstnInitn"`
}

func (d *Document00100101) AddMessage() *AcquirerAuthorisationInitiation {
	d.Message = new(AcquirerAuthorisationInitiation)
	return d.Message
}

// The AcquirerAuthorisationInitiation message is sent by an acquirer or an agent to an issuer or an agent, to request, advice or notify the approval of a card transaction.
type AcquirerAuthorisationInitiation struct {

	// Information related to the protocol management.
	Header *iso20022.Header17 `xml:"Hdr"`

	// Information related to the authorisation initiation.
	AuthorisationInitiation *iso20022.AcquirerAuthorisationInitiation1 `xml:"AuthstnInitn"`

	// Trailer of the message containing a MAC.
	// It corresponds patially to ISO 8583 field number 53, completed by the field number 64 or 128.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (a *AcquirerAuthorisationInitiation) AddHeader() *iso20022.Header17 {
	a.Header = new(iso20022.Header17)
	return a.Header
}

func (a *AcquirerAuthorisationInitiation) AddAuthorisationInitiation() *iso20022.AcquirerAuthorisationInitiation1 {
	a.AuthorisationInitiation = new(iso20022.AcquirerAuthorisationInitiation1)
	return a.AuthorisationInitiation
}

func (a *AcquirerAuthorisationInitiation) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}

