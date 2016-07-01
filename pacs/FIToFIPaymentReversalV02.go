package pacs

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.02 Document"`
	Message *FIToFIPaymentReversalV02 `xml:"FIToFIPmtRvsl"`
}

func (d *Document00700102) AddMessage() *FIToFIPaymentReversalV02 {
	d.Message = new(FIToFIPaymentReversalV02)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentReversal message is sent by an agent to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The FIToFIPaymentReversal message is exchanged between agents to reverse a FIToFICustomerDirectDebit message that has been settled. The result will be a credit on the debtor account.
// The FIToFIPaymentReversal message may or may not be the follow-up of a CustomerDirectDebitInitiation message.
// The FIToFIPaymentReversal message refers to the original FIToFICustomerDirectDebit message by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentReversal message can be used in domestic and cross-border scenarios.
type FIToFIPaymentReversalV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader41 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupInformation22 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*iso20022.PaymentTransactionInformation29 `xml:"TxInf,omitempty"`

}


func (f *FIToFIPaymentReversalV02) AddGroupHeader() *iso20022.GroupHeader41 {
	f.GroupHeader = new(iso20022.GroupHeader41)
	return f.GroupHeader
}

func (f *FIToFIPaymentReversalV02) AddOriginalGroupInformation() *iso20022.OriginalGroupInformation22 {
	f.OriginalGroupInformation = new(iso20022.OriginalGroupInformation22)
	return f.OriginalGroupInformation
}

func (f *FIToFIPaymentReversalV02) AddTransactionInformation() *iso20022.PaymentTransactionInformation29 {
	newValue := new (iso20022.PaymentTransactionInformation29)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}

