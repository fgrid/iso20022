package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600102 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.02 Document"`
	Message *AccountExcludedMandateMaintenanceAmendmentRequestV02 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (d *Document01600102) AddMessage() *AccountExcludedMandateMaintenanceAmendmentRequestV02 {
	d.Message = new(AccountExcludedMandateMaintenanceAmendmentRequestV02)
	return d.Message
}

// The AccountExcludedMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information. Usage: this update is about account details excluding any mandate information.
// If modification codes are not used: the organisation will specify under the “Account” and “Organisation” tags the complete information as it should be in the financial institution’s records after processing the update request.
// If modification codes are used (in that case, they must be used everywhere): the organisation will specify under the “Account” and “Organisation” tags which elements must be added, deleted, modified, or if they are unchanged.
type AccountExcludedMandateMaintenanceAmendmentRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution. OrganisationIdentification6
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccountModification1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *iso20022.OrganisationModification1 `xml:"Org"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddAccount() *iso20022.CustomerAccountModification1 {
	a.Account = new(iso20022.CustomerAccountModification1)
	return a.Account
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddOrganisation() *iso20022.OrganisationModification1 {
	a.Organisation = new(iso20022.OrganisationModification1)
	return a.Organisation
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
