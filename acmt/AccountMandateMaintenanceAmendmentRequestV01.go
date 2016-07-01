package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01800101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.018.001.01 Document"`
	Message *AccountMandateMaintenanceAmendmentRequestV01 `xml:"AcctMndtMntncAmdmntReq"`
}

func (d *Document01800101) AddMessage() *AccountMandateMaintenanceAmendmentRequestV01 {
	d.Message = new(AccountMandateMaintenanceAmendmentRequestV01)
	return d.Message
}

// Scope
// The AccountMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information. This update is only about mandate information.
// Usage
// This message should only be sent in response to a request from the financial institution. This update is only about mandate information. The organisation will specify under the Mandate tag the complete information as it should be in the financial institutions records after processing the update request. It is not possible to update the account characteristics or organisation information with this message.
// It is possible to request to update the mandate information of several accounts, if these accounts must have exactly the same mandates.
// This message could be sent together with other related documents.
type AccountMandateMaintenanceAmendmentRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *iso20022.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*iso20022.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme. 
	// 
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification []*iso20022.OrganisationIdentification6 `xml:"OrgId"`

	// Information specifying the account mandate.
	Mandate []*iso20022.OperationMandate1 `xml:"Mndt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`

}


func (a *AccountMandateMaintenanceAmendmentRequestV01) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddContractDates() *iso20022.AccountContract2 {
	a.ContractDates = new(iso20022.AccountContract2)
	return a.ContractDates
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddUnderlyingMasterAgreement() *iso20022.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(iso20022.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddAccountIdentification() *iso20022.AccountForAction1 {
	newValue := new (iso20022.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddOrganisationIdentification() *iso20022.OrganisationIdentification6 {
	newValue := new (iso20022.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddMandate() *iso20022.OperationMandate1 {
	newValue := new (iso20022.OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new (iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

