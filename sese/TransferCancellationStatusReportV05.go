package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000105 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.05 Document"`
	Message *TransferCancellationStatusReportV05 `xml:"TrfCxlStsRpt"`
}

func (d *Document01000105) AddMessage() *TransferCancellationStatusReportV05 {
	d.Message = new(TransferCancellationStatusReportV05)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferCancellationStatusReport message to the instructing party, for example, an investment manager or one of its authorised representatives to provide the status of a previously received transfer cancellation instruction.
// Usage
// The TransferCancellationStatusReport message is used to report on the status of a transfer in or transfer out cancellation request.
// The reference of the transfer instruction for which the cancellation status is reported is identified in TransferReference. The message identification of the transfer cancellation request message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// The message identification of the transfer instruction request message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// One of the following statuses can be reported:
// - the transfer cancellation is accepted, or,
// - the transfer cancellation has been sent to the next party, or,
// - the transfer cancellation is complete and the reason for the status,
// - the transfer cancellation pending and the reason for the status,
// - the transfer cancellation is rejected and the reason for the status.
type TransferCancellationStatusReportV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *iso20022.AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Reference to the message or communication that was previously received.
	Reference *iso20022.References49Choice `xml:"Ref,omitempty"`

	// Status of the transfer cancellation instruction.
	StatusReport *iso20022.CancellationStatusAndReason3 `xml:"StsRpt"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferCancellationStatusReportV05) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferCancellationStatusReportV05) AddCounterpartyReference() *iso20022.AdditionalReference7 {
	t.CounterpartyReference = new(iso20022.AdditionalReference7)
	return t.CounterpartyReference
}

func (t *TransferCancellationStatusReportV05) AddReference() *iso20022.References49Choice {
	t.Reference = new(iso20022.References49Choice)
	return t.Reference
}

func (t *TransferCancellationStatusReportV05) AddStatusReport() *iso20022.CancellationStatusAndReason3 {
	t.StatusReport = new(iso20022.CancellationStatusAndReason3)
	return t.StatusReport
}

func (t *TransferCancellationStatusReportV05) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferCancellationStatusReportV05) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
