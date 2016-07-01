package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.01 Document"`
	Message *AccountExcludedMandateMaintenanceAmendmentRequestV01 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (d *Document01600101) AddMessage() *AccountExcludedMandateMaintenanceAmendmentRequestV01 {
	d.Message = new(AccountExcludedMandateMaintenanceAmendmentRequestV01)
	return d.Message
}

// Scope
// The AccountExcludedMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information.
// Usage
// This update is about account details excluding any mandate information. The organisation will specify under the Account and Organisation tags the complete information as it should be in the financial institutions records after processing the update request.
type AccountExcludedMandateMaintenanceAmendmentRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the Group to which the organisation belongs, and the account Servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccount1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme. 
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation []*iso20022.Organisation6 `xml:"Org"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`

}


func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddAccount() *iso20022.CustomerAccount1 {
	a.Account = new(iso20022.CustomerAccount1)
	return a.Account
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddOrganisation() *iso20022.Organisation6 {
	newValue := new (iso20022.Organisation6)
	a.Organisation = append(a.Organisation, newValue)
	return newValue
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new (iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

