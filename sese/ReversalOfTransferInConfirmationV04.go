package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.008.001.04 Document"`
	Message *ReversalOfTransferInConfirmationV04 `xml:"RvslOfTrfInConf"`
}

func (d *Document00800104) AddMessage() *ReversalOfTransferInConfirmationV04 {
	d.Message = new(ReversalOfTransferInConfirmationV04)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the ReversalOfTransferInConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to cancel a previously sent TransferInConfirmation message.
// Usage
// The ReversalOfTransferInConfirmation message is used to reverse a previously sent TransferInConfirmation.
// There are two ways to specify the reversal of the transfer in confirmation. Either:
// - the business references, for example, TransferReference, TransferConfirmationIdentification, of the transfer confirmation are quoted, or,
// - all the details of the transfer confirmation (this includes TransferReference and TransferConfirmationIdentification) are quoted but this is not recommended.
// The message identification of the TransferInConfirmation message in which the transfer confirmation was conveyed may also be quoted in PreviousReference.
// The message reference (MessageIdentification) of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// 
type ReversalOfTransferInConfirmationV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References11 `xml:"Refs"`

	// Choice between reversal by reference or by reversal details.
	Reversal *iso20022.Reversal2Choice `xml:"Rvsl"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

}


func (r *ReversalOfTransferInConfirmationV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferInConfirmationV04) AddReferences() *iso20022.References11 {
	newValue := new (iso20022.References11)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *ReversalOfTransferInConfirmationV04) AddReversal() *iso20022.Reversal2Choice {
	r.Reversal = new(iso20022.Reversal2Choice)
	return r.Reversal
}

func (r *ReversalOfTransferInConfirmationV04) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}

