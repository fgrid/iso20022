package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700104 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.007.001.04 Document"`
	Message *SubscriptionBulkOrderV04 `xml:"SbcptBlkOrdr"`
}

func (d *Document00700104) AddMessage() *SubscriptionBulkOrderV04 {
	d.Message = new(SubscriptionBulkOrderV04)
	return d.Message
}

// Scope
// The SubscriptionBulkOrder message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to instruct a subscription to a financial instrument for two or more accounts.
// Usage
// The SubscriptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are for the same financial instrument. This message will typically be used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
type SubscriptionBulkOrderV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// General information related to the orders.
	BulkOrderDetails *iso20022.SubscriptionBulkOrder5 `xml:"BlkOrdrDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderV04) AddPoolReference() *iso20022.AdditionalReference9 {
	s.PoolReference = new(iso20022.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV04) AddBulkOrderDetails() *iso20022.SubscriptionBulkOrder5 {
	s.BulkOrderDetails = new(iso20022.SubscriptionBulkOrder5)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderV04) AddCopyDetails() *iso20022.CopyInformation4 {
	s.CopyDetails = new(iso20022.CopyInformation4)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
