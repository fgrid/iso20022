package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.004.001.01 Document"`
	Message *AcquirerFinancialResponse `xml:"AcqrrFinRspn"`
}

func (d *Document00400101) AddMessage() *AcquirerFinancialResponse {
	d.Message = new(AcquirerFinancialResponse)
	return d.Message
}

// The AcquirerFinancialResponse message is sent by an issuer or an agent to answer to an AcquirerFinancialInitiation message.
type AcquirerFinancialResponse struct {

	// Information related to the protocol management.
	Header *iso20022.Header17 `xml:"Hdr"`

	// Information related to the response of a financial authorisation.
	FinancialResponse *iso20022.AcquirerFinancialResponse1 `xml:"FinRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcquirerFinancialResponse) AddHeader() *iso20022.Header17 {
	a.Header = new(iso20022.Header17)
	return a.Header
}

func (a *AcquirerFinancialResponse) AddFinancialResponse() *iso20022.AcquirerFinancialResponse1 {
	a.FinancialResponse = new(iso20022.AcquirerFinancialResponse1)
	return a.FinancialResponse
}

func (a *AcquirerFinancialResponse) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
