package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.01 Document"`
	Message *AccountOpeningAdditionalInformationRequestV01 `xml:"AcctOpngAddtlInfReq"`
}

func (d *Document00900101) AddMessage() *AccountOpeningAdditionalInformationRequestV01 {
	d.Message = new(AccountOpeningAdditionalInformationRequestV01)
	return d.Message
}

// Scope
// The AccountOpeningAdditionalInformationRequest message is sent from a financial institution to an organisation as part of the account opening process. This message is sent in response to an opening request message from the organisation, if the business content is valid, but additional information is required.
// Usage
// This message should only be sent if additional information is required as part of the account opening process.
type AccountOpeningAdditionalInformationRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References3 `xml:"Refs"`

	// Identification of the organisation.
	OrganisationIdentification []*iso20022.OrganisationIdentification6 `xml:"OrgId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccount1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Account contract established between the organisation or the Group to which the organisation belongs, and the account Servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`

}


func (a *AccountOpeningAdditionalInformationRequestV01) AddReferences() *iso20022.References3 {
	a.References = new(iso20022.References3)
	return a.References
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddOrganisationIdentification() *iso20022.OrganisationIdentification6 {
	newValue := new (iso20022.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddAccount() *iso20022.CustomerAccount1 {
	a.Account = new(iso20022.CustomerAccount1)
	return a.Account
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningAdditionalInformationRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new (iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

