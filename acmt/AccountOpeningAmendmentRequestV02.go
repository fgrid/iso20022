package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.02 Document"`
	Message *AccountOpeningAmendmentRequestV02 `xml:"AcctOpngAmdmntReq"`
}

func (d *Document00800102) AddMessage() *AccountOpeningAmendmentRequestV02 {
	d.Message = new(AccountOpeningAmendmentRequestV02)
	return d.Message
}

// The AccountOpeningAmendmentRequest message is sent from an organisation to a financial institution as part of the account opening process. It is sent in response to a request from the financial institution to send additional information.
type AccountOpeningAmendmentRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccount4 `xml:"Acct"`

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


func (a *AccountOpeningAmendmentRequestV02) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountOpeningAmendmentRequestV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountOpeningAmendmentRequestV02) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountOpeningAmendmentRequestV02) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningAmendmentRequestV02) AddAccount() *iso20022.CustomerAccount4 {
	a.Account = new(iso20022.CustomerAccount4)
	return a.Account
}

func (a *AccountOpeningAmendmentRequestV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningAmendmentRequestV02) AddOrganisation() *iso20022.Organisation12 {
	a.Organisation = new(iso20022.Organisation12)
	return a.Organisation
}

func (a *AccountOpeningAmendmentRequestV02) AddMandate() *iso20022.OperationMandate2 {
	newValue := new (iso20022.OperationMandate2)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountOpeningAmendmentRequestV02) AddGroup() *iso20022.Group1 {
	newValue := new (iso20022.Group1)
	a.Group = append(a.Group, newValue)
	return newValue
}

func (a *AccountOpeningAmendmentRequestV02) AddReferenceAccount() *iso20022.CashAccount24 {
	a.ReferenceAccount = new(iso20022.CashAccount24)
	return a.ReferenceAccount
}

func (a *AccountOpeningAmendmentRequestV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new (iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountOpeningAmendmentRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

