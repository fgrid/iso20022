package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.03 Document"`
	Message *AcceptorReconciliationRequestV03 `xml:"AccptrRcncltnReq"`
}

func (d *Document00900103) AddMessage() *AcceptorReconciliationRequestV03 {
	d.Message = new(AcceptorReconciliationRequestV03)
	return d.Message
}

// The AcceptorReconciliationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationRequestV03 struct {

	// Reconciliation request message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the reconcilliation request.
	ReconciliationRequest *iso20022.AcceptorReconciliationRequest3 `xml:"RcncltnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationRequestV03) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorReconciliationRequestV03) AddReconciliationRequest() *iso20022.AcceptorReconciliationRequest3 {
	a.ReconciliationRequest = new(iso20022.AcceptorReconciliationRequest3)
	return a.ReconciliationRequest
}

func (a *AcceptorReconciliationRequestV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}
