package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400102 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.02 Document"`
	Message *AccountReportV02 `xml:"AcctRpt"`
}

func (d *Document01400102) AddMessage() *AccountReportV02 {
	d.Message = new(AccountReportV02)
	return d.Message
}

// The AccountReport message is sent from a financial institution to an organisation for reporting purposes. It can be sent unsolicited as part of opening, maintenance, or closing process, or it can be sent as response to an AccountReportRequest message.
type AccountReportV02 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References5 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *iso20022.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *iso20022.Organisation12 `xml:"Org"`

	// Account report.
	Report []*iso20022.AccountReport15 `xml:"Rpt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountReportV02) AddReferences() *iso20022.References5 {
	a.References = new(iso20022.References5)
	return a.References
}

func (a *AccountReportV02) AddFrom() *iso20022.OrganisationIdentification8 {
	a.From = new(iso20022.OrganisationIdentification8)
	return a.From
}

func (a *AccountReportV02) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountReportV02) AddOrganisation() *iso20022.Organisation12 {
	a.Organisation = new(iso20022.Organisation12)
	return a.Organisation
}

func (a *AccountReportV02) AddReport() *iso20022.AccountReport15 {
	newValue := new(iso20022.AccountReport15)
	a.Report = append(a.Report, newValue)
	return newValue
}

func (a *AccountReportV02) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountReportV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
