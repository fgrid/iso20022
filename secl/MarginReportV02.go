package secl

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:secl.005.001.02 Document"`
	Message *MarginReportV02 `xml:"MrgnRpt"`
}

func (d *Document00500102) AddMessage() *MarginReportV02 {
	d.Message = new(MarginReportV02)
	return d.Message
}

// Scope
// The MarginReport message is sent by the central counterparty (CCP) to a clearing member to report on:
// - the exposure resulting from the trade positions
// - the value of the collateral held by the CCP (market value of this collateral) and
// - the resulting difference representing the risk encountered by the CCP.
// 
// The message definition is intended for use with the ISO20022 Business Application Header.
// 
// Usage
// There are four possibilities to report the above information. Indeed, the margin report may be structured as follows:
// - per clearing member: the report will only show the information for the clearing member, or
// - per clearing member and per financial instrument: the report will show the information for the clearing member, structured by security identification, or
// - per clearing member and per non clearing member: the report will show the information for the clearing member (that is for global clearing member only) structured by non clearing member(s), or
// - per clearing member and per non clearing member and per security identification: the report will show the information for the clearing member (global clearing member only) structured by non clearing member(s) and by security identification.
type MarginReportV02 struct {

	// Provides parameters of the margin report such as the creation date and time, the report currency or the calculation date and time.
	ReportParameters *iso20022.ReportParameters3 `xml:"RptParams"`

	// Page number of the message (within a report) and continuation indicator to indicate that the report is to continue or that the message is the last page of the report.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides the identification of the account owner, that is the clearing member (individual clearing member or general clearing member).
	ClearingMember *iso20022.PartyIdentification35Choice `xml:"ClrMmb"`

	// Provides details on the valuation of the collateral on deposit.
	ReportSummary *iso20022.MarginCalculation1 `xml:"RptSummry,omitempty"`

	// Provides the margin report details.
	ReportDetails []*iso20022.MarginReport2 `xml:"RptDtls"`

	// Additional information that can't be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MarginReportV02) AddReportParameters() *iso20022.ReportParameters3 {
	m.ReportParameters = new(iso20022.ReportParameters3)
	return m.ReportParameters
}

func (m *MarginReportV02) AddPagination() *iso20022.Pagination {
	m.Pagination = new(iso20022.Pagination)
	return m.Pagination
}

func (m *MarginReportV02) AddClearingMember() *iso20022.PartyIdentification35Choice {
	m.ClearingMember = new(iso20022.PartyIdentification35Choice)
	return m.ClearingMember
}

func (m *MarginReportV02) AddReportSummary() *iso20022.MarginCalculation1 {
	m.ReportSummary = new(iso20022.MarginCalculation1)
	return m.ReportSummary
}

func (m *MarginReportV02) AddReportDetails() *iso20022.MarginReport2 {
	newValue := new (iso20022.MarginReport2)
	m.ReportDetails = append(m.ReportDetails, newValue)
	return newValue
}

func (m *MarginReportV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

