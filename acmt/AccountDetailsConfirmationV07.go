package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200107 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.002.001.07 Document"`
	Message *AccountDetailsConfirmationV07 `xml:"AcctDtlsConf"`
}

func (d *Document00200107) AddMessage() *AccountDetailsConfirmationV07 {
	d.Message = new(AccountDetailsConfirmationV07)
	return d.Message
}

// Scope
// The AccountDetailsConfirmation message is sent by an account servicer, for example, a registrar, transfer agent, custodian bank or securities depository to the account owner, for example, an investor to confirm the opening of an account, execution of an AccountModificationInstruction or to return information requested in a GetAccountDetails message.
// Usage
// The AccountDetailsConfirmation message is used to confirm the opening of an account, modification of an account or the provision of information requested in a previously sent GetAccountDetails message. The message contains detailed information relevant to the opened account.
// When the AccountDetailsConfirmation is used to confirm execution of an AccountModificationInstruction message, it may:
// - contain the modified subsets of account details that were specified in the AccountModificationInstruction, and/or
// - provide the status of the account.
// When the AccountModificationInstruction message is used to instruct the closure of an account, the AccountDetailsConfirmation message is used to confirm the account has been closed.
// When the AccountDetailsConfirmation is used to reply to a GetAccountDetails message, it returns the selected subsets of account details that were specified in the GetAccountDetails message.
type AccountDetailsConfirmationV07 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order or settlement transaction.
	OrderReference *iso20022.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Information about the request or instruction which triggered this confirmation.
	ConfirmationDetails *iso20022.AccountManagementConfirmation4 `xml:"ConfDtls"`

	// Confirmation of the information related to the investment account.
	InvestmentAccount *iso20022.InvestmentAccount62 `xml:"InvstmtAcct,omitempty"`

	// Confirmation of information related to parties that are related to the account, for example, primary account owner.
	AccountParties *iso20022.AccountParties15 `xml:"AcctPties,omitempty"`

	// Confirmation of an intermediary or other party related to the management of the account.
	Intermediaries []*iso20022.Intermediary36 `xml:"Intrmies,omitempty"`

	// Confirmation of referral information.
	Placement *iso20022.ReferredAgent2 `xml:"Plcmnt,omitempty"`

	// Confirmation of eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *iso20022.NewIssueAllocation2 `xml:"NewIsseAllcn,omitempty"`

	// Confirmation of the information related to a savings plan that is related to the account.
	SavingsInvestmentPlan []*iso20022.InvestmentPlan14 `xml:"SvgsInvstmtPlan,omitempty"`

	// Confirmation of the information related to a withdrawal plan that is related to the account.
	WithdrawalInvestmentPlan []*iso20022.InvestmentPlan14 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Confirmation of a cash settlement standing instruction associated to  transactions on the account.
	CashSettlement []*iso20022.CashSettlement1 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*iso20022.DocumentToSend3 `xml:"SvcLvlAgrmt,omitempty"`

	// Additional information such as remarks or notes that must be conveyed about the party and or  limitations and restrictions.
	AdditionalInformation []*iso20022.AdditiononalInformation12 `xml:"AddtlInf,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountDetailsConfirmationV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountDetailsConfirmationV07) AddOrderReference() *iso20022.InvestmentFundOrder4 {
	a.OrderReference = new(iso20022.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountDetailsConfirmationV07) AddRelatedReference() *iso20022.AdditionalReference6 {
	a.RelatedReference = new(iso20022.AdditionalReference6)
	return a.RelatedReference
}

func (a *AccountDetailsConfirmationV07) AddConfirmationDetails() *iso20022.AccountManagementConfirmation4 {
	a.ConfirmationDetails = new(iso20022.AccountManagementConfirmation4)
	return a.ConfirmationDetails
}

func (a *AccountDetailsConfirmationV07) AddInvestmentAccount() *iso20022.InvestmentAccount62 {
	a.InvestmentAccount = new(iso20022.InvestmentAccount62)
	return a.InvestmentAccount
}

func (a *AccountDetailsConfirmationV07) AddAccountParties() *iso20022.AccountParties15 {
	a.AccountParties = new(iso20022.AccountParties15)
	return a.AccountParties
}

func (a *AccountDetailsConfirmationV07) AddIntermediaries() *iso20022.Intermediary36 {
	newValue := new(iso20022.Intermediary36)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV07) AddPlacement() *iso20022.ReferredAgent2 {
	a.Placement = new(iso20022.ReferredAgent2)
	return a.Placement
}

func (a *AccountDetailsConfirmationV07) AddNewIssueAllocation() *iso20022.NewIssueAllocation2 {
	a.NewIssueAllocation = new(iso20022.NewIssueAllocation2)
	return a.NewIssueAllocation
}

func (a *AccountDetailsConfirmationV07) AddSavingsInvestmentPlan() *iso20022.InvestmentPlan14 {
	newValue := new(iso20022.InvestmentPlan14)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV07) AddWithdrawalInvestmentPlan() *iso20022.InvestmentPlan14 {
	newValue := new(iso20022.InvestmentPlan14)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV07) AddCashSettlement() *iso20022.CashSettlement1 {
	newValue := new(iso20022.CashSettlement1)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV07) AddServiceLevelAgreement() *iso20022.DocumentToSend3 {
	newValue := new(iso20022.DocumentToSend3)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV07) AddAdditionalInformation() *iso20022.AdditiononalInformation12 {
	newValue := new(iso20022.AdditiononalInformation12)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountDetailsConfirmationV07) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
