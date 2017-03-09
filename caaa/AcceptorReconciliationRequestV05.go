package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900105 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.05 Document"`
	Message *AcceptorReconciliationRequestV05 `xml:"AccptrRcncltnReq"`
}

func (d *Document00900105) AddMessage() *AcceptorReconciliationRequestV05 {
	d.Message = new(AcceptorReconciliationRequestV05)
	return d.Message
}

// The AcceptorReconciliationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationRequestV05 struct {

	// Reconciliation request message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the reconciliation request.
	ReconciliationRequest *iso20022.AcceptorReconciliationRequest5 `xml:"RcncltnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorReconciliationRequestV05) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorReconciliationRequestV05) AddReconciliationRequest() *iso20022.AcceptorReconciliationRequest5 {
	a.ReconciliationRequest = new(iso20022.AcceptorReconciliationRequest5)
	return a.ReconciliationRequest
}

func (a *AcceptorReconciliationRequestV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
