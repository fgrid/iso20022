package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.011.001.02 Document"`
	Message *AccountRequestRejectionV02 `xml:"AcctReqRjctn"`
}

func (d *Document01100102) AddMessage() *AccountRequestRejectionV02 {
	d.Message = new(AccountRequestRejectionV02)
	return d.Message
}

// The AccountRequestRejection message is sent from a financial institution to an organisation. This message is sent in response to a request message from the organisation, if the business content is not valid.
type AccountRequestRejectionV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References6 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*iso20022.AccountForAction1 `xml:"AcctId,omitempty"`

	// Identifier for an organisation.
	OrganisationIdentification *iso20022.OrganisationIdentification8 `xml:"OrgId"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountRequestRejectionV02) AddReferences() *iso20022.References6 {
	a.References = new(iso20022.References6)
	return a.References
}

func (a *AccountRequestRejectionV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountRequestRejectionV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountRequestRejectionV02) AddAccountIdentification() *iso20022.AccountForAction1 {
	newValue := new(iso20022.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountRequestRejectionV02) AddOrganisationIdentification() *iso20022.OrganisationIdentification8 {
	a.OrganisationIdentification = new(iso20022.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountRequestRejectionV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountRequestRejectionV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
