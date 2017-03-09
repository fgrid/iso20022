package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300107 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.003.001.07 Document"`
	Message *AccountModificationInstructionV07 `xml:"AcctModInstr"`
}

func (d *Document00300107) AddMessage() *AccountModificationInstructionV07 {
	d.Message = new(AccountModificationInstructionV07)
	return d.Message
}

// Scope
// The AccountModificationInstruction message is sent by an account owner, for example, an investor or its designated agent to the account servicer, for example, a registrar, transfer agent, custodian bank or securities depository to modify, that is, create, update or delete specific details of an existing account.
// Usage
// The AccountModificationInstruction message is used to modify the details of an existing account.
// The AccountModificationInstruction message can be used to:
// - maintain/update any of the existing account details, for example, to update the address of the beneficiary or modify the preference to income from distribution to capitalisation, or,
// - add/create specific details to the existing account when these details were not yet recorded at the time of account creation, for example, to add a second address or to establish new cash settlement standing instructions, or,
// - delete specific account details, for example, delete cash standing instructions, or,
// - instruct a change to the status of the account, for example, instruct the closure of the account.
// The usage of this message may be subject to service level agreement (SLA) between the counterparties.
// Execution of the AccountModificationInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountModificationInstructionV07 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Information about the modification instruction.
	InstructionDetails *iso20022.InvestmentAccountModification3 `xml:"InstrDtls,omitempty"`

	// Identifies the account for which the information is modified.
	InvestmentAccountSelection *iso20022.AccountSelection2Choice `xml:"InvstmtAcctSelctn"`

	// Information related to general characteristics of the account to be inserted, updated or deleted.
	ModifiedInvestmentAccount *iso20022.InvestmentAccount63 `xml:"ModfdInvstmtAcct,omitempty"`

	// Information related to an account  party to be inserted, updated or deleted.
	ModifiedAccountParties []*iso20022.AccountParties16 `xml:"ModfdAcctPties,omitempty"`

	// Information related to intermediaries to be inserted, updated or deleted.
	ModifiedIntermediaries []*iso20022.ModificationScope38 `xml:"ModfdIntrmies,omitempty"`

	// Information related to referral information to be inserted, updated or deleted.
	ModifiedPlacement *iso20022.ModificationScope33 `xml:"ModfdPlcmnt,omitempty"`

	// Eligibility conditions related to allocation of new issues to be inserted, updated or deleted.
	ModifiedIssueAllocation *iso20022.ModificationScope21 `xml:"ModfdIsseAllcn,omitempty"`

	// Information related to a savings plan to be either inserted, updated or deleted.
	ModifiedSavingsInvestmentPlan []*iso20022.ModificationScope37 `xml:"ModfdSvgsInvstmtPlan,omitempty"`

	// Information related to a withdrawal plan to be either inserted, updated or deleted.
	ModifiedWithdrawalInvestmentPlan []*iso20022.ModificationScope37 `xml:"ModfdWdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction to be either inserted or deleted.
	ModifiedCashSettlement []*iso20022.CashSettlement2 `xml:"ModfdCshSttlm,omitempty"`

	// Information related to documents to be added, deleted or updated.
	//
	ModifiedServiceLevelAgreement []*iso20022.ModificationScope31 `xml:"ModfdSvcLvlAgrmt,omitempty"`

	// Information related to additional information to be added, deleted or updated.
	ModifiedAdditionalInformation []*iso20022.ModificationScope35 `xml:"ModfdAddtlInf,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountModificationInstructionV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountModificationInstructionV07) AddPreviousReference() *iso20022.AdditionalReference6 {
	a.PreviousReference = new(iso20022.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountModificationInstructionV07) AddInstructionDetails() *iso20022.InvestmentAccountModification3 {
	a.InstructionDetails = new(iso20022.InvestmentAccountModification3)
	return a.InstructionDetails
}

func (a *AccountModificationInstructionV07) AddInvestmentAccountSelection() *iso20022.AccountSelection2Choice {
	a.InvestmentAccountSelection = new(iso20022.AccountSelection2Choice)
	return a.InvestmentAccountSelection
}

func (a *AccountModificationInstructionV07) AddModifiedInvestmentAccount() *iso20022.InvestmentAccount63 {
	a.ModifiedInvestmentAccount = new(iso20022.InvestmentAccount63)
	return a.ModifiedInvestmentAccount
}

func (a *AccountModificationInstructionV07) AddModifiedAccountParties() *iso20022.AccountParties16 {
	newValue := new(iso20022.AccountParties16)
	a.ModifiedAccountParties = append(a.ModifiedAccountParties, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddModifiedIntermediaries() *iso20022.ModificationScope38 {
	newValue := new(iso20022.ModificationScope38)
	a.ModifiedIntermediaries = append(a.ModifiedIntermediaries, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddModifiedPlacement() *iso20022.ModificationScope33 {
	a.ModifiedPlacement = new(iso20022.ModificationScope33)
	return a.ModifiedPlacement
}

func (a *AccountModificationInstructionV07) AddModifiedIssueAllocation() *iso20022.ModificationScope21 {
	a.ModifiedIssueAllocation = new(iso20022.ModificationScope21)
	return a.ModifiedIssueAllocation
}

func (a *AccountModificationInstructionV07) AddModifiedSavingsInvestmentPlan() *iso20022.ModificationScope37 {
	newValue := new(iso20022.ModificationScope37)
	a.ModifiedSavingsInvestmentPlan = append(a.ModifiedSavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddModifiedWithdrawalInvestmentPlan() *iso20022.ModificationScope37 {
	newValue := new(iso20022.ModificationScope37)
	a.ModifiedWithdrawalInvestmentPlan = append(a.ModifiedWithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddModifiedCashSettlement() *iso20022.CashSettlement2 {
	newValue := new(iso20022.CashSettlement2)
	a.ModifiedCashSettlement = append(a.ModifiedCashSettlement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddModifiedServiceLevelAgreement() *iso20022.ModificationScope31 {
	newValue := new(iso20022.ModificationScope31)
	a.ModifiedServiceLevelAgreement = append(a.ModifiedServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddModifiedAdditionalInformation() *iso20022.ModificationScope35 {
	newValue := new(iso20022.ModificationScope35)
	a.ModifiedAdditionalInformation = append(a.ModifiedAdditionalInformation, newValue)
	return newValue
}

func (a *AccountModificationInstructionV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountModificationInstructionV07) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
