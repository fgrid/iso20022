package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05100101 struct {
	XMLName xml.Name                                               `xml:"urn:iso:std:iso:20022:tech:xsd:setr.051.001.01 Document"`
	Message *RedemptionOrderConfirmationCancellationInstructionV01 `xml:"RedOrdrConfCxlInstrV01"`
}

func (d *Document05100101) AddMessage() *RedemptionOrderConfirmationCancellationInstructionV01 {
	d.Message = new(RedemptionOrderConfirmationCancellationInstructionV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the RedemptionOrderConfirmationCancellationInstruction message to the instructing party, for example, an investment manager or its authorised representative to cancel a previously sent RedemptionOrderConfirmation message.
// Usage
// The RedemptionOrderConfirmationCancellationInstruction message is used to cancel one or more previously sent redemption order confirmations. The amendment indicator element is used to specify whether the redemption order confirmation cancellation is to be followed by a RedemptionOrderConfirmationAmendment.
// The RedemptionOrderConfirmationCancellationInstruction message is used to either:
// - cancel an entire RedemptionOrderConfirmation messae, that is, all the individual order confirmations that it contained, or,
// - request the cancellation of one or more individual confirmations.
// There are two ways to use the message.
// (1) When the RedemptionOrderConfirmationCancellationInstruction message is used to cancel an entire message, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of all the individual order confirmations listed in the RedemptionOrderConfirmation message, or,
// - quoting the details of all the individual order confirmations (this includes the OrderReference and DealReference) listed in RedemptionOrderConfirmation message but this is not recommended.
// The message identification of the RedemptionOrderConfirmation message may also be quoted in PreviousReference.
// It is also possible to instruct the cancellation of an entire confirmation message by quoting its message identification in PreviousReference, but this is not recommended.
// (2) When the RedemptionOrderConfirmationCancellationInstruction message is used to cancel one or more individual order confirmations, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of each individual order confirmation listed in the RedemptionOrderConfirmation message, or,
// - quoting the details of each individual order execution (this includes the OrderReference and DealReference) listed in RedemptionOrderConfirmation message but this is not recommended.
// The message identification of the RedemptionOrderConfirmation message in which the individual order confirmation was conveyed may also be quoted in PreviousReference.
// The rejection or acceptance of a RedemptionOrderConfirmationCancellationInstruction is made using an OrderConfirmationStatusReport message.
type RedemptionOrderConfirmationCancellationInstructionV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// References of the orders confirmations to be cancelled.
	CancellationByReference *iso20022.InvestmentFundOrderExecution1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the orders confirmations to be cancelled.
	CancellationByOrderConfirmationDetails *iso20022.RedemptionOrderConfirmation1 `xml:"CxlByOrdrConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	r.RelatedReference = new(iso20022.AdditionalReference3)
	return r.RelatedReference
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddCancellationByReference() *iso20022.InvestmentFundOrderExecution1 {
	r.CancellationByReference = new(iso20022.InvestmentFundOrderExecution1)
	return r.CancellationByReference
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddCancellationByOrderConfirmationDetails() *iso20022.RedemptionOrderConfirmation1 {
	r.CancellationByOrderConfirmationDetails = new(iso20022.RedemptionOrderConfirmation1)
	return r.CancellationByOrderConfirmationDetails
}

func (r *RedemptionOrderConfirmationCancellationInstructionV01) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}
