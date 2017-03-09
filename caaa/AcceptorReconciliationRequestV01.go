package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Document"`
	Message *AcceptorReconciliationRequestV01 `xml:"AccptrRcncltnReq"`
}

func (d *Document00900101) AddMessage() *AcceptorReconciliationRequestV01 {
	d.Message = new(AcceptorReconciliationRequestV01)
	return d.Message
}

// Scope
// The AcceptorReconciliationRequest message is sent by the card acceptor to the acquirer or an agent to communicate the totals of the card payment transaction for a reconciliation period. An agent never forwards the message.
// Usage
// The AcceptorReconciliationRequest message is used to ensure that the debits and credits correspond to the balances computed by the acquirer or its agent for the same reconciliation period.
// The AcceptorReconciliationRequest message is also used to close a reconciliation period.
type AcceptorReconciliationRequestV01 struct {

	// Reconciliation request message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the reconcilliation request.
	ReconciliationRequest *iso20022.AcceptorReconciliationRequest1 `xml:"RcncltnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorReconciliationRequestV01) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorReconciliationRequestV01) AddReconciliationRequest() *iso20022.AcceptorReconciliationRequest1 {
	a.ReconciliationRequest = new(iso20022.AcceptorReconciliationRequest1)
	return a.ReconciliationRequest
}

func (a *AcceptorReconciliationRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType3 {
	a.SecurityTrailer = new(iso20022.ContentInformationType3)
	return a.SecurityTrailer
}
