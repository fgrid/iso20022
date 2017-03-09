package pacs

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.02 Document"`
	Message *FinancialInstitutionCreditTransferV02 `xml:"FinInstnCdtTrf"`
}

func (d *Document00900102) AddMessage() *FinancialInstitutionCreditTransferV02 {
	d.Message = new(FinancialInstitutionCreditTransferV02)
	return d.Message
}

// Scope
// The FinancialInstitutionCreditTransfer message is sent by a debtor financial institution to a creditor financial institution, directly or through other agents and/or a payment clearing and settlement system.
// It is used to move funds from a debtor account to a creditor, where both debtor and creditor are financial institutions.
// Usage
// The FinancialInstitutionCreditTransfer message is exchanged between agents and can contain one or more credit transfer instructions where debtor and creditor are both financial institutions.
// The FinancialInstitutionCreditTransfer message does not allow for grouping: a CreditTransferTransactionInformation block must be present for each credit transfer transaction.
// The FinancialInstitutionCreditTransfer message can be used in domestic and cross-border scenarios.
type FinancialInstitutionCreditTransferV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader35 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*iso20022.CreditTransferTransactionInformation13 `xml:"CdtTrfTxInf"`
}

func (f *FinancialInstitutionCreditTransferV02) AddGroupHeader() *iso20022.GroupHeader35 {
	f.GroupHeader = new(iso20022.GroupHeader35)
	return f.GroupHeader
}

func (f *FinancialInstitutionCreditTransferV02) AddCreditTransferTransactionInformation() *iso20022.CreditTransferTransactionInformation13 {
	newValue := new(iso20022.CreditTransferTransactionInformation13)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}
