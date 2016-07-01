package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.002.001.04 Document"`
	Message *TransferOutCancellationRequestV04 `xml:"TrfOutCxlReq"`
}

func (d *Document00200104) AddMessage() *TransferOutCancellationRequestV04 {
	d.Message = new(TransferOutCancellationRequestV04)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferOutCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent TransferOutInstruction.
// Usage
// The TransferOutCancellationRequest message is used to request cancellation of a previously sent TransferOutInstruction. There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferOutInstruction message in which the original transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferOutInstruction message by quoting its message identification in PreviousReference.
type TransferOutCancellationRequestV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References11 `xml:"Refs"`

	// Choice between cancellation by reference or by transfer details.
	Cancellation *iso20022.Cancellation1Choice `xml:"Cxl"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

}


func (t *TransferOutCancellationRequestV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutCancellationRequestV04) AddReferences() *iso20022.References11 {
	newValue := new (iso20022.References11)
	t.References = append(t.References, newValue)
	return newValue
}

func (t *TransferOutCancellationRequestV04) AddCancellation() *iso20022.Cancellation1Choice {
	t.Cancellation = new(iso20022.Cancellation1Choice)
	return t.Cancellation
}

func (t *TransferOutCancellationRequestV04) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}

