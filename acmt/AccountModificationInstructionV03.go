package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.003.001.03 Document"`
	Message *AccountModificationInstructionV03 `xml:"AcctModInstr"`
}

func (d *Document00300103) AddMessage() *AccountModificationInstructionV03 {
	d.Message = new(AccountModificationInstructionV03)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent, sends the AccountModificationInstruction message to the account servicer, for example, a registrar, transfer agent or custodian bank to modify, that is, create, update or delete specific details of an existing investment fund account.
// Usage
// The AccountModificationInstruction message is used to modify the details of an existing account.
// The AccountModificationInstruction message has three specific uses:
// - to maintain/update any of the existing account details, for example, to update the address of the beneficiary or modify the preference to income from distribution to capitalisation, or,
// - to add/create specific details to the existing account when these details were not yet recorded at the time of account creation, for example, to add a second address or to establish new cash settlement standing instructions, or,
// - to delete specific account details, for example, delete cash standing instructions.
// This message cannot be used to delete an entire account, as institution specific and regulatory rules pertaining to account deletion are diverse.
// The usage of this message may be subject to service level agreement (SLA) between the counterparties.
// Execution of the AccountModificationInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountModificationInstructionV03 struct {

	// Identifies the message.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Provide detailed information about the application modification instruction.
	InstructionDetails *iso20022.InvestmentAccountModificationDetails `xml:"InstrDtls,omitempty"`

	// Investment account selection information used to identify the account for which the information is modified..
	InvestmentAccountSelection *iso20022.InvestmentAccountSelection2 `xml:"InvstmtAcctSelctn"`

	// Information related to general characteristics of an investment account to be inserted, updated or deleted.
	ModifiedInvestmentAccount *iso20022.InvestmentAccount36 `xml:"ModfdInvstmtAcct,omitempty"`

	// Information related to the account related parties (eg. account owner) to be inserted, updated or deleted.
	ModifiedAccountParties []*iso20022.AccountParties7 `xml:"ModfdAcctPties,omitempty"`

	// Information related to intermediaries to be inserted, updated or deleted.
	ModifiedIntermediaries []*iso20022.ModificationScope7 `xml:"ModfdIntrmies,omitempty"`

	// Information related to referred placement agent in the hedge fund industry to be inserted, updated or deleted.
	ModifiedPlacement *iso20022.ReferredAgent1 `xml:"ModfdPlcmnt,omitempty"`

	// Eligibility conditions information related to new issues allocation to be inserted, updated or deleted.
	ModifiedIssueAllocation *iso20022.ModificationScope9 `xml:"ModfdIsseAllcn,omitempty"`

	// Information related to a savings plan to be either inserted, updated or deleted.
	ModifiedSavingsInvestmentPlan []*iso20022.ModificationScope16 `xml:"ModfdSvgsInvstmtPlan,omitempty"`

	// Information related to a withdrawal plan to be either inserted, updated or deleted.
	ModifiedWithdrawalInvestmentPlan []*iso20022.ModificationScope16 `xml:"ModfdWdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction associated to the investment fund transaction and to be either inserted or deleted.
	ModifiedCashSettlement []*iso20022.InvestmentFundCashSettlementInformation6 `xml:"ModfdCshSttlm,omitempty"`

	// Information related to documents to be added, deleted or updated.
	//
	ModifiedServiceLevelAgreement []*iso20022.ModificationScope10 `xml:"ModfdSvcLvlAgrmt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountModificationInstructionV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountModificationInstructionV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	a.PreviousReference = new(iso20022.AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountModificationInstructionV03) AddInstructionDetails() *iso20022.InvestmentAccountModificationDetails {
	a.InstructionDetails = new(iso20022.InvestmentAccountModificationDetails)
	return a.InstructionDetails
}

func (a *AccountModificationInstructionV03) AddInvestmentAccountSelection() *iso20022.InvestmentAccountSelection2 {
	a.InvestmentAccountSelection = new(iso20022.InvestmentAccountSelection2)
	return a.InvestmentAccountSelection
}

func (a *AccountModificationInstructionV03) AddModifiedInvestmentAccount() *iso20022.InvestmentAccount36 {
	a.ModifiedInvestmentAccount = new(iso20022.InvestmentAccount36)
	return a.ModifiedInvestmentAccount
}

func (a *AccountModificationInstructionV03) AddModifiedAccountParties() *iso20022.AccountParties7 {
	newValue := new(iso20022.AccountParties7)
	a.ModifiedAccountParties = append(a.ModifiedAccountParties, newValue)
	return newValue
}

func (a *AccountModificationInstructionV03) AddModifiedIntermediaries() *iso20022.ModificationScope7 {
	newValue := new(iso20022.ModificationScope7)
	a.ModifiedIntermediaries = append(a.ModifiedIntermediaries, newValue)
	return newValue
}

func (a *AccountModificationInstructionV03) AddModifiedPlacement() *iso20022.ReferredAgent1 {
	a.ModifiedPlacement = new(iso20022.ReferredAgent1)
	return a.ModifiedPlacement
}

func (a *AccountModificationInstructionV03) AddModifiedIssueAllocation() *iso20022.ModificationScope9 {
	a.ModifiedIssueAllocation = new(iso20022.ModificationScope9)
	return a.ModifiedIssueAllocation
}

func (a *AccountModificationInstructionV03) AddModifiedSavingsInvestmentPlan() *iso20022.ModificationScope16 {
	newValue := new(iso20022.ModificationScope16)
	a.ModifiedSavingsInvestmentPlan = append(a.ModifiedSavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV03) AddModifiedWithdrawalInvestmentPlan() *iso20022.ModificationScope16 {
	newValue := new(iso20022.ModificationScope16)
	a.ModifiedWithdrawalInvestmentPlan = append(a.ModifiedWithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV03) AddModifiedCashSettlement() *iso20022.InvestmentFundCashSettlementInformation6 {
	newValue := new(iso20022.InvestmentFundCashSettlementInformation6)
	a.ModifiedCashSettlement = append(a.ModifiedCashSettlement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV03) AddModifiedServiceLevelAgreement() *iso20022.ModificationScope10 {
	newValue := new(iso20022.ModificationScope10)
	a.ModifiedServiceLevelAgreement = append(a.ModifiedServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV03) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
