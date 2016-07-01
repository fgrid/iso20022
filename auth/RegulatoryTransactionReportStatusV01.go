package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:auth.010.001.01 Document"`
	Message *RegulatoryTransactionReportStatusV01 `xml:"RgltryTxRptStsV01"`
}

func (d *Document01000101) AddMessage() *RegulatoryTransactionReportStatusV01 {
	d.Message = new(RegulatoryTransactionReportStatusV01)
	return d.Message
}

// Scope
// A regulator or an intermediary sends the RegulatoryTransactionReportStatus to a reporting institution to provide the status of a RegulatoryTransactionReport previously sent by the reporting institution.
// Usage
// The message definition may be used to provide a status for the entire report or to provide a status at the level of individual transactions within the report. One of the following statuses can be reported:
// - Completed, or,
// - Pending, or,
// - Rejected.
// If the status is rejected, then reason for the rejection must be specified.
type RegulatoryTransactionReportStatusV01 struct {

	// Identification of the RegulatoryTransactionReportStatus document.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the firm that is legally responsible for sending the transaction report.
	ReportingInstitution *iso20022.PartyIdentification23Choice `xml:"RptgInstn"`

	// Provides the status of the entire RegulatoryTransactionReport that was previously sent by the reporting institution.
	ReportStatus *iso20022.ReportStatusAndReason1 `xml:"RptSts"`

	// Provides the status of one or more transactions that were previously sent within a RegulatoryTransactionReport by the reporting institution.
	IndividualTransactionStatus []*iso20022.TradeTransactionStatusAndReason1 `xml:"IndvTxSts"`

}


func (r *RegulatoryTransactionReportStatusV01) AddIdentification() *iso20022.DocumentIdentification8 {
	r.Identification = new(iso20022.DocumentIdentification8)
	return r.Identification
}

func (r *RegulatoryTransactionReportStatusV01) AddReportingInstitution() *iso20022.PartyIdentification23Choice {
	r.ReportingInstitution = new(iso20022.PartyIdentification23Choice)
	return r.ReportingInstitution
}

func (r *RegulatoryTransactionReportStatusV01) AddReportStatus() *iso20022.ReportStatusAndReason1 {
	r.ReportStatus = new(iso20022.ReportStatusAndReason1)
	return r.ReportStatus
}

func (r *RegulatoryTransactionReportStatusV01) AddIndividualTransactionStatus() *iso20022.TradeTransactionStatusAndReason1 {
	newValue := new (iso20022.TradeTransactionStatusAndReason1)
	r.IndividualTransactionStatus = append(r.IndividualTransactionStatus, newValue)
	return newValue
}

