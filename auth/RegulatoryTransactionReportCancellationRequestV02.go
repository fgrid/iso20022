package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900102 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Document"`
	Message *RegulatoryTransactionReportCancellationRequestV02 `xml:"RgltryTxRptCxlReq"`
}

func (d *Document00900102) AddMessage() *RegulatoryTransactionReportCancellationRequestV02 {
	d.Message = new(RegulatoryTransactionReportCancellationRequestV02)
	return d.Message
}

// Scope
// A reporting institution, eg, an investment bank, sends the RegulatoryTransactionReportCancellationRequest to a regulator or to an intermediary (eg a reporting agent), to request a cancellation of a previously sent RegulatoryTransactionReport.
// Usage
// The message definition can be used to cancel an entire RegulatoryTransactionReport or to cancel one or more individual transactions in a previously sent RegulatoryTransactionReport.
type RegulatoryTransactionReportCancellationRequestV02 struct {

	// Identification of the RegulatoryTransactionReportCancellationRequest document.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the firm that executed the transaction.
	//
	ReportingInstitution *iso20022.PartyIdentification23Choice `xml:"RptgInstn"`

	// Identifies the intermediary which is reporting on behalf on the reporting institution.
	ReportingAgent *iso20022.PartyIdentification24Choice `xml:"RptgAgt,omitempty"`

	// Provides all the details of the transaction report that needs to be cancelled. More than one set of details can be provided.
	//
	//
	CancellationByTransactionDetails []*iso20022.TransactionDetails3 `xml:"CxlByTxDtls"`

	// Provides the reference of the RegulatoryTransactionReport document that was previously sent and that needs to be cancelled in its entirety.
	PreviousReference *iso20022.DocumentIdentification9 `xml:"PrvsRef"`

	// Provides the trade reference of the transaction report that needs to be cancelled. More than one reference may be provided.
	CancellationByTradeReference []*iso20022.TransactionDetails2 `xml:"CxlByTradRef"`
}

func (r *RegulatoryTransactionReportCancellationRequestV02) AddIdentification() *iso20022.DocumentIdentification8 {
	r.Identification = new(iso20022.DocumentIdentification8)
	return r.Identification
}

func (r *RegulatoryTransactionReportCancellationRequestV02) AddReportingInstitution() *iso20022.PartyIdentification23Choice {
	r.ReportingInstitution = new(iso20022.PartyIdentification23Choice)
	return r.ReportingInstitution
}

func (r *RegulatoryTransactionReportCancellationRequestV02) AddReportingAgent() *iso20022.PartyIdentification24Choice {
	r.ReportingAgent = new(iso20022.PartyIdentification24Choice)
	return r.ReportingAgent
}

func (r *RegulatoryTransactionReportCancellationRequestV02) AddCancellationByTransactionDetails() *iso20022.TransactionDetails3 {
	newValue := new(iso20022.TransactionDetails3)
	r.CancellationByTransactionDetails = append(r.CancellationByTransactionDetails, newValue)
	return newValue
}

func (r *RegulatoryTransactionReportCancellationRequestV02) AddPreviousReference() *iso20022.DocumentIdentification9 {
	r.PreviousReference = new(iso20022.DocumentIdentification9)
	return r.PreviousReference
}

func (r *RegulatoryTransactionReportCancellationRequestV02) AddCancellationByTradeReference() *iso20022.TransactionDetails2 {
	newValue := new(iso20022.TransactionDetails2)
	r.CancellationByTradeReference = append(r.CancellationByTradeReference, newValue)
	return newValue
}
