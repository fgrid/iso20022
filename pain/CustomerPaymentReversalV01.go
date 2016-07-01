package pain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.01 Document"`
	Message *CustomerPaymentReversalV01 `xml:"pain.007.001.01"`
}

func (d *Document00700101) AddMessage() *CustomerPaymentReversalV01 {
	d.Message = new(CustomerPaymentReversalV01)
	return d.Message
}

// Scope
// The CustomerPaymentReversal message is sent by the initiating party to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The CustomerPaymentReversal message is exchanged between a non-financial institution customer and an agent to reverse a CustomerDirectDebitInitiation message that has been settled. The result will be a credit on the debtor account.
// The CustomerPaymentReversal message refers to the original CustomerDirectDebitInitiation message by means of references only or by means of references and a set of elements from the original instruction.
// The CustomerPaymentReversal message can be used in domestic and cross-border scenarios.
type CustomerPaymentReversalV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader8 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupInformation5 `xml:"OrgnlGrpInf"`

	// Information concerning the original transactions, to which the reversal message refers.
	TransactionInformation []*iso20022.PaymentTransactionInformation4 `xml:"TxInf,omitempty"`

}


func (c *CustomerPaymentReversalV01) AddGroupHeader() *iso20022.GroupHeader8 {
	c.GroupHeader = new(iso20022.GroupHeader8)
	return c.GroupHeader
}

func (c *CustomerPaymentReversalV01) AddOriginalGroupInformation() *iso20022.OriginalGroupInformation5 {
	c.OriginalGroupInformation = new(iso20022.OriginalGroupInformation5)
	return c.OriginalGroupInformation
}

func (c *CustomerPaymentReversalV01) AddTransactionInformation() *iso20022.PaymentTransactionInformation4 {
	newValue := new (iso20022.PaymentTransactionInformation4)
	c.TransactionInformation = append(c.TransactionInformation, newValue)
	return newValue
}

