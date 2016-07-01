package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800107 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.008.001.07 Document"`
	Message *ReversalOfTransferInConfirmationV07 `xml:"RvslOfTrfInConf"`
}

func (d *Document00800107) AddMessage() *ReversalOfTransferInConfirmationV07 {
	d.Message = new(ReversalOfTransferInConfirmationV07)
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
type ReversalOfTransferInConfirmationV07 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References20 `xml:"Refs,omitempty"`

	// Function of the transfer-in, that is, whether the message is used as a reversal of a previously sent confirmation or as a reversal of a previously sent advice. The absence of Function indicates the message is a reversal of a previously sent confirmation.
	Function *iso20022.TransferInFunction2Code `xml:"Fctn,omitempty"`

	// Choice between reversal by reference or by reversal details.
	Reversal *iso20022.Reversal7Choice `xml:"Rvsl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

}


func (r *ReversalOfTransferInConfirmationV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferInConfirmationV07) AddReferences() *iso20022.References20 {
	newValue := new (iso20022.References20)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *ReversalOfTransferInConfirmationV07) SetFunction(value string) {
	r.Function = (*iso20022.TransferInFunction2Code)(&value)
}

func (r *ReversalOfTransferInConfirmationV07) AddReversal() *iso20022.Reversal7Choice {
	r.Reversal = new(iso20022.Reversal7Choice)
	return r.Reversal
}

func (r *ReversalOfTransferInConfirmationV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	r.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return r.MarketPracticeVersion
}

func (r *ReversalOfTransferInConfirmationV07) AddCopyDetails() *iso20022.CopyInformation4 {
	r.CopyDetails = new(iso20022.CopyInformation4)
	return r.CopyDetails
}

