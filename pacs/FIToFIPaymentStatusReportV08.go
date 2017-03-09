package pacs

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200108 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Document"`
	Message *FIToFIPaymentStatusReportV08 `xml:"FIToFIPmtStsRpt"`
}

func (d *Document00200108) AddMessage() *FIToFIPaymentStatusReportV08 {
	d.Message = new(FIToFIPaymentStatusReportV08)
	return d.Message
}

// Scope
// The FIToFIPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusReport message can be used in domestic and cross-border scenarios.
// The FIToFIPaymentStatusReport may also be sent to the receiver of the payment in a real time payment scenario, as both sides of the transactions must be informed of the status of the transaction (e.g. either the beneficiary is credited, or the transaction is rejected).
type FIToFIPaymentStatusReportV08 struct {

	// Set of characteristics shared by all individual transactions included in the status report message.
	GroupHeader *iso20022.GroupHeader53 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus []*iso20022.OriginalGroupHeader7 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Information concerning the original transactions, to which the status report message refers.
	TransactionInformationAndStatus []*iso20022.PaymentTransaction80 `xml:"TxInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFIPaymentStatusReportV08) AddGroupHeader() *iso20022.GroupHeader53 {
	f.GroupHeader = new(iso20022.GroupHeader53)
	return f.GroupHeader
}

func (f *FIToFIPaymentStatusReportV08) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupHeader7 {
	newValue := new(iso20022.OriginalGroupHeader7)
	f.OriginalGroupInformationAndStatus = append(f.OriginalGroupInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV08) AddTransactionInformationAndStatus() *iso20022.PaymentTransaction80 {
	newValue := new(iso20022.PaymentTransaction80)
	f.TransactionInformationAndStatus = append(f.TransactionInformationAndStatus, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
