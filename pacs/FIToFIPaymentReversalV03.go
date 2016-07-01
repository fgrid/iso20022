package pacs

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.03 Document"`
	Message *FIToFIPaymentReversalV03 `xml:"FIToFIPmtRvsl"`
}

func (d *Document00700103) AddMessage() *FIToFIPaymentReversalV03 {
	d.Message = new(FIToFIPaymentReversalV03)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionPaymentReversal message is sent by an agent to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The FIToFIPaymentReversal message is exchanged between agents to reverse a FIToFICustomerDirectDebit message that has been settled. The result will be a credit on the debtor account.
// The FIToFIPaymentReversal message may or may not be the follow-up of a CustomerDirectDebitInitiation message.
// The FIToFIPaymentReversal message refers to the original FIToFICustomerDirectDebit message by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentReversal message can be used in domestic and cross-border scenarios.
type FIToFIPaymentReversalV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader57 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupHeader3 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*iso20022.PaymentTransaction36 `xml:"TxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (f *FIToFIPaymentReversalV03) AddGroupHeader() *iso20022.GroupHeader57 {
	f.GroupHeader = new(iso20022.GroupHeader57)
	return f.GroupHeader
}

func (f *FIToFIPaymentReversalV03) AddOriginalGroupInformation() *iso20022.OriginalGroupHeader3 {
	f.OriginalGroupInformation = new(iso20022.OriginalGroupHeader3)
	return f.OriginalGroupInformation
}

func (f *FIToFIPaymentReversalV03) AddTransactionInformation() *iso20022.PaymentTransaction36 {
	newValue := new (iso20022.PaymentTransaction36)
	f.TransactionInformation = append(f.TransactionInformation, newValue)
	return newValue
}

func (f *FIToFIPaymentReversalV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}

