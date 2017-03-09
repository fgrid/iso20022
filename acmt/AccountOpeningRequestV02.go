package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.007.001.02 Document"`
	Message *AccountOpeningRequestV02 `xml:"AcctOpngReq"`
}

func (d *Document00700102) AddMessage() *AccountOpeningRequestV02 {
	d.Message = new(AccountOpeningRequestV02)
	return d.Message
}

// The AccountOpeningRequest message is sent from an organisation to a financial institution as part of the account opening process. It is the initial request message to open an account.
type AccountOpeningRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccount4 `xml:"Acct"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *iso20022.Organisation12 `xml:"Org"`

	// Information specifying the account mandate.
	Mandate []*iso20022.OperationMandate2 `xml:"Mndt,omitempty"`

	// Definition of a group of parties.
	Group []*iso20022.Group1 `xml:"Grp,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *iso20022.CashAccount24 `xml:"RefAcct,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountOpeningRequestV02) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountOpeningRequestV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountOpeningRequestV02) AddAccount() *iso20022.CustomerAccount4 {
	a.Account = new(iso20022.CustomerAccount4)
	return a.Account
}

func (a *AccountOpeningRequestV02) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountOpeningRequestV02) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningRequestV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningRequestV02) AddOrganisation() *iso20022.Organisation12 {
	a.Organisation = new(iso20022.Organisation12)
	return a.Organisation
}

func (a *AccountOpeningRequestV02) AddMandate() *iso20022.OperationMandate2 {
	newValue := new(iso20022.OperationMandate2)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountOpeningRequestV02) AddGroup() *iso20022.Group1 {
	newValue := new(iso20022.Group1)
	a.Group = append(a.Group, newValue)
	return newValue
}

func (a *AccountOpeningRequestV02) AddReferenceAccount() *iso20022.CashAccount24 {
	a.ReferenceAccount = new(iso20022.CashAccount24)
	return a.ReferenceAccount
}

func (a *AccountOpeningRequestV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountOpeningRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
