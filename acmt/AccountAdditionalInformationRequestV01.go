package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.012.001.01 Document"`
	Message *AccountAdditionalInformationRequestV01 `xml:"AcctAddtlInfReq"`
}

func (d *Document01200101) AddMessage() *AccountAdditionalInformationRequestV01 {
	d.Message = new(AccountAdditionalInformationRequestV01)
	return d.Message
}

// Scope
// The AccountAdditionalInformationRequest message is sent from a financial institution to an organisation as part of maintenance process. This message is sent in response to a request message from the organisation, if the business content is valid, but additional information is required.
// Usage
// This message should only be sent if additional information is required as part of the account maintenance process.
type AccountAdditionalInformationRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References3 `xml:"Refs"`

	// Identifier for an organisation.
	OrganisationIdentification []*iso20022.OrganisationIdentification6 `xml:"OrgId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification *iso20022.AccountForAction1 `xml:"AcctId"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountAdditionalInformationRequestV01) AddReferences() *iso20022.References3 {
	a.References = new(iso20022.References3)
	return a.References
}

func (a *AccountAdditionalInformationRequestV01) AddOrganisationIdentification() *iso20022.OrganisationIdentification6 {
	newValue := new(iso20022.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountAdditionalInformationRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountAdditionalInformationRequestV01) AddAccountIdentification() *iso20022.AccountForAction1 {
	a.AccountIdentification = new(iso20022.AccountForAction1)
	return a.AccountIdentification
}

func (a *AccountAdditionalInformationRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new(iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
