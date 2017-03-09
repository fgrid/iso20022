package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600106 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.006.001.06 Document"`
	Message *AccountManagementStatusReportV06 `xml:"AcctMgmtStsRpt"`
}

func (d *Document00600106) AddMessage() *AccountManagementStatusReportV06 {
	d.Message = new(AccountManagementStatusReportV06)
	return d.Message
}

// Scope
// The AccountManagementStatusReport message is sent by an account servicer, for example, a registrar, transfer agent, custodian bank or securities depository to the account owner or its designated agent, for example, an investor to report on the receipt or the processing status of a previously received account management message.
// Usage
// The AccountManagementStatusReport message is used to provide the status of a previously received AccountOpeningInstruction, an AccountModificationInstruction or a GetAccountDetails message. It may also be used to report the status of the account.
// The AccountManagementStatusReport message is also used to reject an AccountOpeningInstruction, AccountModificationInstruction or GetAccountDetails message when the message is not compliant with the agreed SLA or when the account cannot be uniquely identified.
type AccountManagementStatusReportV06 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference6 `xml:"RltdRef"`

	// Status report details of the account management instruction that was previously received.
	StatusReport *iso20022.AccountManagementStatusAndReason5 `xml:"StsRpt"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountManagementStatusReportV06) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountManagementStatusReportV06) AddRelatedReference() *iso20022.AdditionalReference6 {
	newValue := new(iso20022.AdditionalReference6)
	a.RelatedReference = append(a.RelatedReference, newValue)
	return newValue
}

func (a *AccountManagementStatusReportV06) AddStatusReport() *iso20022.AccountManagementStatusAndReason5 {
	a.StatusReport = new(iso20022.AccountManagementStatusAndReason5)
	return a.StatusReport
}

func (a *AccountManagementStatusReportV06) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountManagementStatusReportV06) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
