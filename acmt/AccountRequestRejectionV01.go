package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.011.001.01 Document"`
	Message *AccountRequestRejectionV01 `xml:"AcctReqRjctn"`
}

func (d *Document01100101) AddMessage() *AccountRequestRejectionV01 {
	d.Message = new(AccountRequestRejectionV01)
	return d.Message
}

// Scope
// The AccountRequestRejection message is sent from a financial institution to an organisation. This message is sent in response to a request message from the organisation, if the business content is not valid.
// Usage
type AccountRequestRejectionV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References6 `xml:"Refs"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*iso20022.AccountForAction1 `xml:"AcctId,omitempty"`

	// Identifier for an organisation.
	OrganisationIdentification []*iso20022.OrganisationIdentification6 `xml:"OrgId"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`

}


func (a *AccountRequestRejectionV01) AddReferences() *iso20022.References6 {
	a.References = new(iso20022.References6)
	return a.References
}

func (a *AccountRequestRejectionV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountRequestRejectionV01) AddAccountIdentification() *iso20022.AccountForAction1 {
	newValue := new (iso20022.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountRequestRejectionV01) AddOrganisationIdentification() *iso20022.OrganisationIdentification6 {
	newValue := new (iso20022.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountRequestRejectionV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new (iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

