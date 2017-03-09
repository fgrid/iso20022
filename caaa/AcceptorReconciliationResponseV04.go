package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.04 Document"`
	Message *AcceptorReconciliationResponseV04 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000104) AddMessage() *AcceptorReconciliationResponseV04 {
	d.Message = new(AcceptorReconciliationResponseV04)
	return d.Message
}

// The AcceptorReconciliationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationResponseV04 struct {

	// Reconciliation response message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the reconciliation response.
	ReconciliationResponse *iso20022.AcceptorReconciliationResponse3 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationResponseV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorReconciliationResponseV04) AddReconciliationResponse() *iso20022.AcceptorReconciliationResponse3 {
	a.ReconciliationResponse = new(iso20022.AcceptorReconciliationResponse3)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}
