package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.001.001.03 Document"`
	Message *AccountOpeningInstructionV03 `xml:"AcctOpngInstr"`
}

func (d *Document00100103) AddMessage() *AccountOpeningInstructionV03 {
	d.Message = new(AccountOpeningInstructionV03)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent sends the AccountOpeningInstruction message to the account servicer, for example, a registrar, transfer agent or custodian to instruct the opening of an account or the opening of an account and establishing an investment plan.
// Usage
// The AccountOpeningInstruction is used to open an account directly or indirectly with the account servicer or an intermediary.
// In some markets, for example, Australia, and for some products in the United Kingdom, a first order (also known as a deposit instruction) is placed at the same time as the account opening. To cater for this scenario, an order message can be linked (via references in the message) to the AccountOpeningInstruction message when needed.
// Execution of the AccountOpeningInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountOpeningInstructionV03 struct {

	// Identifies the message.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order.
	OrderReference *iso20022.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Provide detailed information about the opening instruction.
	InstructionDetails *iso20022.InvestmentAccountOpeningDetails `xml:"InstrDtls"`

	// Detailed information about the investment account to be opened
	InvestmentAccount *iso20022.InvestmentAccount34 `xml:"InvstmtAcct"`

	// Information related to parties who are related to an investment account; eg. primary account owner.
	AccountParties *iso20022.AccountParties6 `xml:"AcctPties"`

	// Information related to an intermediary.
	Intermediaries []*iso20022.Intermediary12 `xml:"Intrmies,omitempty"`

	// Placement agent for the hedge fund industry.
	Placement *iso20022.ReferredAgent1 `xml:"Plcmnt,omitempty"`

	// Eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *iso20022.NewIssueAllocation1 `xml:"NewIsseAllcn,omitempty"`

	// Plan that allows individuals to set aside a fixed amount of money at specified intervals, usually for a special purpose, eg, retirement.
	SavingsInvestmentPlan []*iso20022.InvestmentPlan6 `xml:"SvgsInvstmtPlan,omitempty"`

	// Plan through which an investment fund investor's holdings are depleted through regular withdrawals at specified intervals.
	WithdrawalInvestmentPlan []*iso20022.InvestmentPlan6 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction associated to the investment fund transaction.
	CashSettlement []*iso20022.InvestmentFundCashSettlementInformation5 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*iso20022.DocumentToSend1 `xml:"SvcLvlAgrmt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (a *AccountOpeningInstructionV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountOpeningInstructionV03) AddOrderReference() *iso20022.InvestmentFundOrder4 {
	a.OrderReference = new(iso20022.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountOpeningInstructionV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	a.PreviousReference = new(iso20022.AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountOpeningInstructionV03) AddInstructionDetails() *iso20022.InvestmentAccountOpeningDetails {
	a.InstructionDetails = new(iso20022.InvestmentAccountOpeningDetails)
	return a.InstructionDetails
}

func (a *AccountOpeningInstructionV03) AddInvestmentAccount() *iso20022.InvestmentAccount34 {
	a.InvestmentAccount = new(iso20022.InvestmentAccount34)
	return a.InvestmentAccount
}

func (a *AccountOpeningInstructionV03) AddAccountParties() *iso20022.AccountParties6 {
	a.AccountParties = new(iso20022.AccountParties6)
	return a.AccountParties
}

func (a *AccountOpeningInstructionV03) AddIntermediaries() *iso20022.Intermediary12 {
	newValue := new (iso20022.Intermediary12)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV03) AddPlacement() *iso20022.ReferredAgent1 {
	a.Placement = new(iso20022.ReferredAgent1)
	return a.Placement
}

func (a *AccountOpeningInstructionV03) AddNewIssueAllocation() *iso20022.NewIssueAllocation1 {
	a.NewIssueAllocation = new(iso20022.NewIssueAllocation1)
	return a.NewIssueAllocation
}

func (a *AccountOpeningInstructionV03) AddSavingsInvestmentPlan() *iso20022.InvestmentPlan6 {
	newValue := new (iso20022.InvestmentPlan6)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV03) AddWithdrawalInvestmentPlan() *iso20022.InvestmentPlan6 {
	newValue := new (iso20022.InvestmentPlan6)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV03) AddCashSettlement() *iso20022.InvestmentFundCashSettlementInformation5 {
	newValue := new (iso20022.InvestmentFundCashSettlementInformation5)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV03) AddServiceLevelAgreement() *iso20022.DocumentToSend1 {
	newValue := new (iso20022.DocumentToSend1)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

