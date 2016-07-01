package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.006.001.05 Document"`
	Message *AccountManagementStatusReportV05 `xml:"AcctMgmtStsRpt"`
}

func (d *Document00600105) AddMessage() *AccountManagementStatusReportV05 {
	d.Message = new(AccountManagementStatusReportV05)
	return d.Message
}

// Scope
// An account servicer, for example, a registrar, transfer agent, custodian bank or securities depository  sends the AccountManagementStatusReport message to the account owner or its designated agent, for example, an investor to report on the receipt or the processing status of a previously received AccountOpeningInstruction or AccountModificationInstruction or GetAccountDetails message.
// Usage
// The AccountManagementStatusReport message is used to provide the processing status of a previously received AccountOpeningInstruction or of an AccountModificationInstruction message.
// The AccountManagementStatusReport message is also used by an account servicer to reject an AccountOpeningInstruction or AccountModificationInstruction or GetAccountDetails message when the message is not compliant with the agreed SLA or when the account cannot be uniquely identified.
// The account owner may report that the status of the instruction is either rejected, accepted, that the instruction is being processed or that the instruction has been forwarded to the next intermediary party for further processing.
type AccountManagementStatusReportV05 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference6 `xml:"RltdRef"`

	// Status report details of an account opening instruction or account modification instruction that was previously received.
	StatusReport *iso20022.AccountManagementStatusAndReason4 `xml:"StsRpt"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (a *AccountManagementStatusReportV05) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountManagementStatusReportV05) AddRelatedReference() *iso20022.AdditionalReference6 {
	newValue := new (iso20022.AdditionalReference6)
	a.RelatedReference = append(a.RelatedReference, newValue)
	return newValue
}

func (a *AccountManagementStatusReportV05) AddStatusReport() *iso20022.AccountManagementStatusAndReason4 {
	a.StatusReport = new(iso20022.AccountManagementStatusAndReason4)
	return a.StatusReport
}

func (a *AccountManagementStatusReportV05) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountManagementStatusReportV05) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

