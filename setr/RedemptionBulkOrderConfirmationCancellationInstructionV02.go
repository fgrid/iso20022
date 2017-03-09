package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05300102 struct {
	XMLName xml.Name                                                   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.053.001.02 Document"`
	Message *RedemptionBulkOrderConfirmationCancellationInstructionV02 `xml:"RedBlkOrdrConfCxlInstr"`
}

func (d *Document05300102) AddMessage() *RedemptionBulkOrderConfirmationCancellationInstructionV02 {
	d.Message = new(RedemptionBulkOrderConfirmationCancellationInstructionV02)
	return d.Message
}

// Scope
// The RedemptionBulkOrderConfirmationCancellationInstruction message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to cancel a previously sent RedemptionBulkOrderConfirmation.
// Usage
// To request the cancellation of one or more individual order executions, the order reference and deal reference of each individual order execution in the original RedemptionBulkOrderConfirmation are specified in the order reference and deal reference elements respectively. The message identification of the RedemptionBulkOrderConfirmation message in which the individual order execution was conveyed may also be quoted in PreviousReference but this is not recommended. The AmendmentIndicator is used to specify whether the redemption order confirmation cancellation is to be followed by an amendment An amendment of a redemption order confirmation is carried out by sending a RedemptionBulkOrderConfirmation message in which the AmendmentIndicator contains the value ‘true’.
type RedemptionBulkOrderConfirmationCancellationInstructionV02 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *iso20022.YesNoIndicator `xml:"AmdmntInd"`

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the individual order confirmation to be cancelled.
	OrderReferences []*iso20022.InvestmentFundOrder11 `xml:"OrdrRefs"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) AddPoolReference() *iso20022.AdditionalReference9 {
	r.PoolReference = new(iso20022.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) AddRelatedReference() *iso20022.AdditionalReference8 {
	r.RelatedReference = new(iso20022.AdditionalReference8)
	return r.RelatedReference
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) SetAmendmentIndicator(value string) {
	r.AmendmentIndicator = (*iso20022.YesNoIndicator)(&value)
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) SetMasterReference(value string) {
	r.MasterReference = (*iso20022.Max35Text)(&value)
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) AddOrderReferences() *iso20022.InvestmentFundOrder11 {
	newValue := new(iso20022.InvestmentFundOrder11)
	r.OrderReferences = append(r.OrderReferences, newValue)
	return newValue
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV02) AddCopyDetails() *iso20022.CopyInformation4 {
	r.CopyDetails = new(iso20022.CopyInformation4)
	return r.CopyDetails
}
