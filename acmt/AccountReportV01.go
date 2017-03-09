package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400101 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.01 Document"`
	Message *AccountReportV01 `xml:"AcctRpt"`
}

func (d *Document01400101) AddMessage() *AccountReportV01 {
	d.Message = new(AccountReportV01)
	return d.Message
}

// Scope
// The AccountReport message is sent from a financial institution to an organisation for reporting purposes.
// Usage
// It can be sent unsolicited as part of opening, maintenance, or closing process, or it can be sent as response to an AccountReportRequest message.
type AccountReportV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References5 `xml:"Refs"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation []*iso20022.Organisation6 `xml:"Org"`

	// Account report.
	Report []*iso20022.AccountReport1 `xml:"Rpt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountReportV01) AddReferences() *iso20022.References5 {
	a.References = new(iso20022.References5)
	return a.References
}

func (a *AccountReportV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountReportV01) AddOrganisation() *iso20022.Organisation6 {
	newValue := new(iso20022.Organisation6)
	a.Organisation = append(a.Organisation, newValue)
	return newValue
}

func (a *AccountReportV01) AddReport() *iso20022.AccountReport1 {
	newValue := new(iso20022.AccountReport1)
	a.Report = append(a.Report, newValue)
	return newValue
}

func (a *AccountReportV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new(iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
