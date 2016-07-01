package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.01 Document"`
	Message *AcceptorReconciliationResponseV01 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000101) AddMessage() *AcceptorReconciliationResponseV01 {
	d.Message = new(AcceptorReconciliationResponseV01)
	return d.Message
}

// Scope
// The AcceptorReconciliationResponse message is sent by the acquirer to communicate to the card acceptor the totals of the card payment transaction performed for the reconciliation period. An agent never forwards the message.
// Usage
// The AcceptorReconciliationResponse message is used to compare the totals between a card acceptor and an acquirer for the reconciliation period.
type AcceptorReconciliationResponseV01 struct {

	// Reconciliation response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to thereconciliation response.
	ReconciliationResponse *iso20022.AcceptorReconciliationResponse1 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`

}


func (a *AcceptorReconciliationResponseV01) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorReconciliationResponseV01) AddReconciliationResponse() *iso20022.AcceptorReconciliationResponse1 {
	a.ReconciliationResponse = new(iso20022.AcceptorReconciliationResponse1)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}

