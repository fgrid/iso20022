package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04200106 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.06 Document"`
	Message *CorporateActionInstructionStatementReportV06 `xml:"CorpActnInstrStmtRpt"`
}

func (d *Document04200106) AddMessage() *CorporateActionInstructionStatementReportV06 {
	d.Message = new(CorporateActionInstructionStatementReportV06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionStatementReport message to an account owner or its designated agent to report balances at the safekeeping account level for one or more corporate action events or at the corporate action event level for one or more safekeeping accounts.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionStatementReportV06 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General characteristics related to a statement which reports information.
	StatementGeneralDetails *iso20022.Statement47 `xml:"StmtGnlDtls"`

	// Account information and detailed account holdings information report for corporate action events.
	AccountAndStatementDetails []*iso20022.AccountIdentification40 `xml:"AcctAndStmtDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionStatementReportV06) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CorporateActionInstructionStatementReportV06) AddStatementGeneralDetails() *iso20022.Statement47 {
	c.StatementGeneralDetails = new(iso20022.Statement47)
	return c.StatementGeneralDetails
}

func (c *CorporateActionInstructionStatementReportV06) AddAccountAndStatementDetails() *iso20022.AccountIdentification40 {
	newValue := new(iso20022.AccountIdentification40)
	c.AccountAndStatementDetails = append(c.AccountAndStatementDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatementReportV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
