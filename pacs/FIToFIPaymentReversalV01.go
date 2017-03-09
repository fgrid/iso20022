package pacs

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.01 Document"`
	Message *FIToFIPaymentReversalV01 `xml:"pacs.007.001.01"`
}

func (d *Document00700101) AddMessage() *FIToFIPaymentReversalV01 {
	d.Message = new(FIToFIPaymentReversalV01)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentReversal message is sent by an agent to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The FIToFIPaymentReversal message is exchanged between agents to reverse a FIToFICustomerDirectDebit message that has been settled. The result will be a credit on the debtor account.
// The FIToFIPaymentReversal message may or may not be the follow-up of a CustomerDirectDebitInitiation message.
// The FIToFIPaymentReversal message refers to the original FIToFICustomerDirectDebit message by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentReversal message can be used in domestic and cross-border scenarios.
//
type FIToFIPaymentReversalV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader9 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupInformation5 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*iso20022.PaymentTransactionInformation5 `xml:"TxInf,omitempty"`
}

func (f *FIToFIPaymentReversalV01) AddGroupHeader() *iso20022.GroupHeader9 {
	f.GroupHeader = new(iso20022.GroupHeader9)
	return f.GroupHeader
}

func (f *FIToFIPaymentReversalV01) AddOriginalGroupInformation() *iso20022.OriginalGroupInformation5 {
	f.OriginalGroupInformation = new(iso20022.OriginalGroupInformation5)
	return f.OriginalGroupInformation
}

func (f *FIToFIPaymentReversalV01) AddTransactionInformation() *iso20022.PaymentTransactionInformation5 {
	newValue := new(iso20022.PaymentTransactionInformation5)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}
