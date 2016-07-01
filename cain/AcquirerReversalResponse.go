package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:cain.006.001.01 Document"`
	Message *AcquirerReversalResponse `xml:"AcqrrRvslRspn"`
}

func (d *Document00600101) AddMessage() *AcquirerReversalResponse {
	d.Message = new(AcquirerReversalResponse)
	return d.Message
}

// The AcquirerReversalResponse message is sent by an issuer or an agent to answer to an AcquirerReversalInitiation message.
type AcquirerReversalResponse struct {

	// Information related to the protocol management.
	Header *iso20022.Header18 `xml:"Hdr"`

	// Information related to the response of a reversal.
	ReversalResponse *iso20022.AcquirerReversalResponse1 `xml:"RvslRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (a *AcquirerReversalResponse) AddHeader() *iso20022.Header18 {
	a.Header = new(iso20022.Header18)
	return a.Header
}

func (a *AcquirerReversalResponse) AddReversalResponse() *iso20022.AcquirerReversalResponse1 {
	a.ReversalResponse = new(iso20022.AcquirerReversalResponse1)
	return a.ReversalResponse
}

func (a *AcquirerReversalResponse) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}

