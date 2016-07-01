package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.02 Document"`
	Message *AccountOpeningAdditionalInformationRequestV02 `xml:"AcctOpngAddtlInfReq"`
}

func (d *Document00900102) AddMessage() *AccountOpeningAdditionalInformationRequestV02 {
	d.Message = new(AccountOpeningAdditionalInformationRequestV02)
	return d.Message
}

// The AccountOpeningAdditionalInformationRequest message is sent from a financial institution to an organisation as part of the account opening process. This message is sent in response to an opening request message from the organisation, if the business content is valid, but additional information is required.
type AccountOpeningAdditionalInformationRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References3 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution. OrganisationIdentification6
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Identification of the organisation.
	OrganisationIdentification *iso20022.OrganisationIdentification8 `xml:"OrgId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *iso20022.CustomerAccount4 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (a *AccountOpeningAdditionalInformationRequestV02) AddReferences() *iso20022.References3 {
	a.References = new(iso20022.References3)
	return a.References
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddOrganisationIdentification() *iso20022.OrganisationIdentification8 {
	a.OrganisationIdentification = new(iso20022.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddAccount() *iso20022.CustomerAccount4 {
	a.Account = new(iso20022.CustomerAccount4)
	return a.Account
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new (iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountOpeningAdditionalInformationRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

