package pain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700106 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.06 Document"`
	Message *CustomerPaymentReversalV06 `xml:"CstmrPmtRvsl"`
}

func (d *Document00700106) AddMessage() *CustomerPaymentReversalV06 {
	d.Message = new(CustomerPaymentReversalV06)
	return d.Message
}

// Scope
// The CustomerPaymentReversal message is sent by the initiating party to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The CustomerPaymentReversal message is exchanged between a non-financial institution customer and an agent to reverse a CustomerDirectDebitInitiation message that has been settled. The result will be a credit on the debtor account.
// The CustomerPaymentReversal message refers to the original CustomerDirectDebitInitiation message by means of references only or by means of references and a set of elements from the original instruction.
// The CustomerPaymentReversal message can be used in domestic and cross-border scenarios.
type CustomerPaymentReversalV06 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader56 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *iso20022.OriginalGroupHeader3 `xml:"OrgnlGrpInf"`

	// Information concerning the original payment information, to which the reversal message refers.
	OriginalPaymentInformationAndReversal []*iso20022.OriginalPaymentInstruction16 `xml:"OrgnlPmtInfAndRvsl,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *CustomerPaymentReversalV06) AddGroupHeader() *iso20022.GroupHeader56 {
	c.GroupHeader = new(iso20022.GroupHeader56)
	return c.GroupHeader
}

func (c *CustomerPaymentReversalV06) AddOriginalGroupInformation() *iso20022.OriginalGroupHeader3 {
	c.OriginalGroupInformation = new(iso20022.OriginalGroupHeader3)
	return c.OriginalGroupInformation
}

func (c *CustomerPaymentReversalV06) AddOriginalPaymentInformationAndReversal() *iso20022.OriginalPaymentInstruction16 {
	newValue := new (iso20022.OriginalPaymentInstruction16)
	c.OriginalPaymentInformationAndReversal = append(c.OriginalPaymentInformationAndReversal, newValue)
	return newValue
}

func (c *CustomerPaymentReversalV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

