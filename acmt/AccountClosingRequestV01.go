package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01900101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.019.001.01 Document"`
	Message *AccountClosingRequestV01 `xml:"AcctClsgReq"`
}

func (d *Document01900101) AddMessage() *AccountClosingRequestV01 {
	d.Message = new(AccountClosingRequestV01)
	return d.Message
}

// Scope
// The AccountClosingRequest message is sent from an organisation to a financial institution as part of the account closing process. It is the initial request message to close an account.
// Usage
// This message is the initial request message to close an account.
// It could be sent together with other related documents.
type AccountClosingRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References4 `xml:"Refs"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification *iso20022.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification []*iso20022.OrganisationIdentification6 `xml:"OrgId"`

	// Specifies target dates.
	ContractDates *iso20022.AccountContract4 `xml:"CtrctDts,omitempty"`

	// Identification of the account to which the remaining positive balance of the account to be closed must be transferred or account from which funds can be moved to the account to be closed and which balance is negative. This account must be held in the same financial institution as the account to be closed if the transfer account is used to compensate a negative balance. For a positive balance to be transferred, an account in another financial institution might be used. In that case the account servicer is mandatory.
	BalanceTransferAccount *iso20022.AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme, that is the servicer of the transfer account.
	TransferAccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"TrfAcctSvcrId,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountClosingRequestV01) AddReferences() *iso20022.References4 {
	a.References = new(iso20022.References4)
	return a.References
}

func (a *AccountClosingRequestV01) AddAccountIdentification() *iso20022.AccountForAction1 {
	a.AccountIdentification = new(iso20022.AccountForAction1)
	return a.AccountIdentification
}

func (a *AccountClosingRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountClosingRequestV01) AddOrganisationIdentification() *iso20022.OrganisationIdentification6 {
	newValue := new(iso20022.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountClosingRequestV01) AddContractDates() *iso20022.AccountContract4 {
	a.ContractDates = new(iso20022.AccountContract4)
	return a.ContractDates
}

func (a *AccountClosingRequestV01) AddBalanceTransferAccount() *iso20022.AccountForAction1 {
	a.BalanceTransferAccount = new(iso20022.AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountClosingRequestV01) AddTransferAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.TransferAccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.TransferAccountServicerIdentification
}

func (a *AccountClosingRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new(iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
