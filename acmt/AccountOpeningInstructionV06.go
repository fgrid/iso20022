package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100106 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.001.001.06 Document"`
	Message *AccountOpeningInstructionV06 `xml:"AcctOpngInstr"`
}

func (d *Document00100106) AddMessage() *AccountOpeningInstructionV06 {
	d.Message = new(AccountOpeningInstructionV06)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent sends the AccountOpeningInstruction message to the account servicer, for example, a registrar, transfer agent, custodian or securities depository to instruct the opening of an account or the opening of an account and the establishment of an investment plan.
// Usage
// The AccountOpeningInstruction is used to open an account directly or indirectly with the account servicer or an intermediary.
// In some markets, for example, Australia, and for some products in the United Kingdom, a first order (also known as a deposit instruction) is placed at the same time as the account opening. To cater for this scenario, an order message can be linked (via references in the message) to the AccountOpeningInstruction message when needed.
// Execution of the AccountOpeningInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountOpeningInstructionV06 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order or settlement transaction.
	OrderReference *iso20022.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Information about the opening instruction.
	InstructionDetails *iso20022.InvestmentAccountOpening2 `xml:"InstrDtls"`

	// Detailed information about the account to be opened.
	InvestmentAccount *iso20022.InvestmentAccount49 `xml:"InvstmtAcct"`

	// Information related to parties that are related to the account, for example, primary account owner.
	AccountParties *iso20022.AccountParties13 `xml:"AcctPties"`

	// Intermediary or other party related to the management of the account. In some markets, when this intermediary is a party acting on behalf of the investor for which it has opened an account at, for example, a central securities depository or international central securities depository, this party is known by the investor as the 'account controller'.
	Intermediaries []*iso20022.Intermediary36 `xml:"Intrmies,omitempty"`

	// Referral information.
	Placement *iso20022.ReferredAgent2 `xml:"Plcmnt,omitempty"`

	// Eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *iso20022.NewIssueAllocation2 `xml:"NewIsseAllcn,omitempty"`

	// Plan that allows individuals to set aside a fixed amount of money at specified intervals, usually for a special purpose, for example, retirement.
	SavingsInvestmentPlan []*iso20022.InvestmentPlan12 `xml:"SvgsInvstmtPlan,omitempty"`

	// Plan through which holdings are depleted through regular withdrawals at specified intervals.
	WithdrawalInvestmentPlan []*iso20022.InvestmentPlan12 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction associated to  transactions on the account.
	CashSettlement []*iso20022.CashSettlement1 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*iso20022.DocumentToSend3 `xml:"SvcLvlAgrmt,omitempty"`

	// Additional information concerning limitations and restrictions on the account.
	AdditionalInformation []*iso20022.AccountRestrictions1 `xml:"AddtlInf,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountOpeningInstructionV06) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountOpeningInstructionV06) AddOrderReference() *iso20022.InvestmentFundOrder4 {
	a.OrderReference = new(iso20022.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountOpeningInstructionV06) AddPreviousReference() *iso20022.AdditionalReference6 {
	a.PreviousReference = new(iso20022.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountOpeningInstructionV06) AddInstructionDetails() *iso20022.InvestmentAccountOpening2 {
	a.InstructionDetails = new(iso20022.InvestmentAccountOpening2)
	return a.InstructionDetails
}

func (a *AccountOpeningInstructionV06) AddInvestmentAccount() *iso20022.InvestmentAccount49 {
	a.InvestmentAccount = new(iso20022.InvestmentAccount49)
	return a.InvestmentAccount
}

func (a *AccountOpeningInstructionV06) AddAccountParties() *iso20022.AccountParties13 {
	a.AccountParties = new(iso20022.AccountParties13)
	return a.AccountParties
}

func (a *AccountOpeningInstructionV06) AddIntermediaries() *iso20022.Intermediary36 {
	newValue := new(iso20022.Intermediary36)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV06) AddPlacement() *iso20022.ReferredAgent2 {
	a.Placement = new(iso20022.ReferredAgent2)
	return a.Placement
}

func (a *AccountOpeningInstructionV06) AddNewIssueAllocation() *iso20022.NewIssueAllocation2 {
	a.NewIssueAllocation = new(iso20022.NewIssueAllocation2)
	return a.NewIssueAllocation
}

func (a *AccountOpeningInstructionV06) AddSavingsInvestmentPlan() *iso20022.InvestmentPlan12 {
	newValue := new(iso20022.InvestmentPlan12)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV06) AddWithdrawalInvestmentPlan() *iso20022.InvestmentPlan12 {
	newValue := new(iso20022.InvestmentPlan12)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV06) AddCashSettlement() *iso20022.CashSettlement1 {
	newValue := new(iso20022.CashSettlement1)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV06) AddServiceLevelAgreement() *iso20022.DocumentToSend3 {
	newValue := new(iso20022.DocumentToSend3)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV06) AddAdditionalInformation() *iso20022.AccountRestrictions1 {
	newValue := new(iso20022.AccountRestrictions1)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV06) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountOpeningInstructionV06) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
