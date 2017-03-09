package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.02 Document"`
	Message *AcceptorReconciliationResponseV02 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000102) AddMessage() *AcceptorReconciliationResponseV02 {
	d.Message = new(AcceptorReconciliationResponseV02)
	return d.Message
}

// The AcceptorReconciliationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationResponseV02 struct {

	// Reconciliation response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to thereconciliation response.
	ReconciliationResponse *iso20022.AcceptorReconciliationResponse2 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationResponseV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorReconciliationResponseV02) AddReconciliationResponse() *iso20022.AcceptorReconciliationResponse2 {
	a.ReconciliationResponse = new(iso20022.AcceptorReconciliationResponse2)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
