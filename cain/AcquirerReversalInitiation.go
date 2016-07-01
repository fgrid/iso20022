package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:cain.005.001.01 Document"`
	Message *AcquirerReversalInitiation `xml:"AcqrrRvslInitn"`
}

func (d *Document00500101) AddMessage() *AcquirerReversalInitiation {
	d.Message = new(AcquirerReversalInitiation)
	return d.Message
}

// The AcquirerReversalInitiation message is sent by an acquirer or an agent to an issuer or an agent, to request, advice or notify the reversal of a card transaction.
type AcquirerReversalInitiation struct {

	// Information related to the protocol management
	Header *iso20022.Header18 `xml:"Hdr"`

	// Information related to the reversal.
	ReversalInitiation *iso20022.AcquirerReversalInitiation1 `xml:"RvslInitn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (a *AcquirerReversalInitiation) AddHeader() *iso20022.Header18 {
	a.Header = new(iso20022.Header18)
	return a.Header
}

func (a *AcquirerReversalInitiation) AddReversalInitiation() *iso20022.AcquirerReversalInitiation1 {
	a.ReversalInitiation = new(iso20022.AcquirerReversalInitiation1)
	return a.ReversalInitiation
}

func (a *AcquirerReversalInitiation) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}

