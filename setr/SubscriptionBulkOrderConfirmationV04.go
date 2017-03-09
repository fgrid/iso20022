package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900104 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.009.001.04 Document"`
	Message *SubscriptionBulkOrderConfirmationV04 `xml:"SbcptBlkOrdrConf"`
}

func (d *Document00900104) AddMessage() *SubscriptionBulkOrderConfirmationV04 {
	d.Message = new(SubscriptionBulkOrderConfirmationV04)
	return d.Message
}

// Scope
// The SubscriptionBulkOrderConfirmation message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to confirm the details of the execution of a SubscriptionBulkOrder instruction.
// Usage
// The SubscriptionBulkOrderConfirmation message is used to confirm the execution of all individual orders.
// There is usually one bulk confirmation message for one bulk order message.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the SubscriptionBulkOrder message in which the individual order was conveyed may also be quoted in RelatedReference.
// A SubscriptionBulkOrder must in all cases be responded to by a SubscriptionBulkOrderConfirmation and in no circumstances by a SubscriptionOrderConfirmation.
// If the executing party needs to confirm a SubscriptionOrder instruction, then the SubscriptionOrderConfirmation must be used.
// When the message is used to convey a confirmation amendment/s, the AmendmentIndicator must be present with the value ‘true’ or ‘1’. When this is the case, the message must only contain a confirmation amendment/s and not contain both a confirmation amendment/s and a ‘new’ confirmation/s.
type SubscriptionBulkOrderConfirmationV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// General information related to the execution of the orders.
	BulkExecutionDetails *iso20022.SubscriptionBulkExecution4 `xml:"BlkExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderConfirmationV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderConfirmationV04) AddPoolReference() *iso20022.AdditionalReference9 {
	s.PoolReference = new(iso20022.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderConfirmationV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationV04) AddRelatedReference() *iso20022.AdditionalReference8 {
	s.RelatedReference = new(iso20022.AdditionalReference8)
	return s.RelatedReference
}

func (s *SubscriptionBulkOrderConfirmationV04) AddBulkExecutionDetails() *iso20022.SubscriptionBulkExecution4 {
	s.BulkExecutionDetails = new(iso20022.SubscriptionBulkExecution4)
	return s.BulkExecutionDetails
}

func (s *SubscriptionBulkOrderConfirmationV04) AddCopyDetails() *iso20022.CopyInformation4 {
	s.CopyDetails = new(iso20022.CopyInformation4)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderConfirmationV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
