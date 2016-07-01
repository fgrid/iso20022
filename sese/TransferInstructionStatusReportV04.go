package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.04 Document"`
	Message *TransferInstructionStatusReportV04 `xml:"TrfInstrStsRpt"`
}

func (d *Document01100104) AddMessage() *TransferInstructionStatusReportV04 {
	d.Message = new(TransferInstructionStatusReportV04)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferInstructionStatusReport message to the instructing party, for example, an investment manager or one of its authorised representatives to provide the status of a previously received transfer instruction.
// Usage
// The TransferInstructionStatusReport message is used to report on the status of a transfer in or transfer out instruction. The reference of the transfer instruction for which the status is reported is identified in TransferReference.
// The message identification of the transfer instruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// One of the following statuses can be reported:
// - an accepted status, or,
// - an already executed status, or,
// - a sent to next party status, or,
// - a matched status, or,
// - a settled status, or,
// - a pending settlement status and the reason for the status, or,
// - an unmatched status and the reason for the status, or,
// - an in-repair status and the reason for the status, or,
// - a rejected status and the reason for the status, or,
// - a failed settlement status and the reason for the status, or,
// - a cancelled status and the reason for the status, or,
// - a cancelled status and the reason for the status, or,
// - a cancellation pending status and the reason for the status.
type TransferInstructionStatusReportV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *iso20022.AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Reference to the message or communication that was previously received.
	Reference *iso20022.References40Choice `xml:"Ref,omitempty"`

	// Status of the transfer instruction.
	StatusReport *iso20022.TransferStatusAndReason3 `xml:"StsRpt"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (t *TransferInstructionStatusReportV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInstructionStatusReportV04) AddCounterpartyReference() *iso20022.AdditionalReference2 {
	t.CounterpartyReference = new(iso20022.AdditionalReference2)
	return t.CounterpartyReference
}

func (t *TransferInstructionStatusReportV04) AddReference() *iso20022.References40Choice {
	t.Reference = new(iso20022.References40Choice)
	return t.Reference
}

func (t *TransferInstructionStatusReportV04) AddStatusReport() *iso20022.TransferStatusAndReason3 {
	t.StatusReport = new(iso20022.TransferStatusAndReason3)
	return t.StatusReport
}

func (t *TransferInstructionStatusReportV04) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInstructionStatusReportV04) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

