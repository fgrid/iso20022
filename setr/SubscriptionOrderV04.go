package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000104 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.010.001.04 Document"`
	Message *SubscriptionOrderV04 `xml:"SbcptOrdr"`
}

func (d *Document01000104) AddMessage() *SubscriptionOrderV04 {
	d.Message = new(SubscriptionOrderV04)
	return d.Message
}

// Scope
// The SubscriptionOrder message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to instruct the subscription of one or more financial instruments for one investment fund account.
// Usage
// The SubscriptionOrder message is used to instruct single subscription orders, that is, a message containing one order for one financial instrument for one investment account. The SubscriptionOrder message may also be used for multiple orders, that is, a message containing several orders for the same investment account for different financial instruments.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for the same financial instrument but for different accounts that are to be communicated in a single message, then the SubscriptionBulkOrder message must be used.
type SubscriptionOrderV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// General information related to the orders.
	MultipleOrderDetails *iso20022.SubscriptionMultipleOrder6 `xml:"MltplOrdrDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionOrderV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderV04) AddPoolReference() *iso20022.AdditionalReference9 {
	s.PoolReference = new(iso20022.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionOrderV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderV04) AddMultipleOrderDetails() *iso20022.SubscriptionMultipleOrder6 {
	s.MultipleOrderDetails = new(iso20022.SubscriptionMultipleOrder6)
	return s.MultipleOrderDetails
}

func (s *SubscriptionOrderV04) AddCopyDetails() *iso20022.CopyInformation4 {
	s.CopyDetails = new(iso20022.CopyInformation4)
	return s.CopyDetails
}

func (s *SubscriptionOrderV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
