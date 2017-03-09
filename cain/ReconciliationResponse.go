package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:cain.008.001.01 Document"`
	Message *ReconciliationResponse `xml:"RcncltnRspn"`
}

func (d *Document00800101) AddMessage() *ReconciliationResponse {
	d.Message = new(ReconciliationResponse)
	return d.Message
}

// The ReconciliationResponse message is sent by an issuer or an agent to return the reconciled totals for debits, credits, chargebacks and other transactions.
type ReconciliationResponse struct {

	// Information related to the protocol management.
	Header *iso20022.Header17 `xml:"Hdr"`

	// Information related to the response to a reconciliation.
	ReconciliationResponse *iso20022.AcquirerReconciliationResponse1 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (r *ReconciliationResponse) AddHeader() *iso20022.Header17 {
	r.Header = new(iso20022.Header17)
	return r.Header
}

func (r *ReconciliationResponse) AddReconciliationResponse() *iso20022.AcquirerReconciliationResponse1 {
	r.ReconciliationResponse = new(iso20022.AcquirerReconciliationResponse1)
	return r.ReconciliationResponse
}

func (r *ReconciliationResponse) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	r.SecurityTrailer = new(iso20022.ContentInformationType15)
	return r.SecurityTrailer
}
