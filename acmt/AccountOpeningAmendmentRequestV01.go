package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.01 Document"`
	Message *AccountOpeningAmendmentRequestV01 `xml:"AcctOpngAmdmntReq"`
}

func (d *Document00800101) AddMessage() *AccountOpeningAmendmentRequestV01 {
	d.Message = new(AccountOpeningAmendmentRequestV01)
	return d.Message
}

// Scope
// The AccountOpeningAmendmentRequest message is sent from an organisation to a financial institution as part of the account opening process. It is sent in response to a request from the financial institution to provide additional information.
// Usage
// This message should only be sent if additional information is requested as part of the account opening process.
// It could be sent together with other related documents.
type AccountOpeningAmendmentRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the Group to which the organisation belongs, and the account Servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccount1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation []*iso20022.Organisation6 `xml:"Org"`

	// Information specifying the account mandate.
	Mandate []*iso20022.OperationMandate1 `xml:"Mndt,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *iso20022.CashAccount16 `xml:"RefAcct,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountOpeningAmendmentRequestV01) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountOpeningAmendmentRequestV01) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountOpeningAmendmentRequestV01) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningAmendmentRequestV01) AddAccount() *iso20022.CustomerAccount1 {
	a.Account = new(iso20022.CustomerAccount1)
	return a.Account
}

func (a *AccountOpeningAmendmentRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningAmendmentRequestV01) AddOrganisation() *iso20022.Organisation6 {
	newValue := new(iso20022.Organisation6)
	a.Organisation = append(a.Organisation, newValue)
	return newValue
}

func (a *AccountOpeningAmendmentRequestV01) AddMandate() *iso20022.OperationMandate1 {
	newValue := new(iso20022.OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountOpeningAmendmentRequestV01) AddReferenceAccount() *iso20022.CashAccount16 {
	a.ReferenceAccount = new(iso20022.CashAccount16)
	return a.ReferenceAccount
}

func (a *AccountOpeningAmendmentRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new(iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
