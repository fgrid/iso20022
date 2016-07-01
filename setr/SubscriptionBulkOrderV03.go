package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.007.001.03 Document"`
	Message *SubscriptionBulkOrderV03 `xml:"SbcptBlkOrdrV03"`
}

func (d *Document00700103) AddMessage() *SubscriptionBulkOrderV03 {
	d.Message = new(SubscriptionBulkOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative sends the SubscriptionBulkOrder message to the executing party, for example, a transfer agent, to instruct a subscription to a financial instrument for two or more accounts.
// Usage
// The SubscriptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are related to the same financial instrument. This message will typically be used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for different financial instruments but for the same account, then the SubscriptionOrder must be used.
type SubscriptionBulkOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	BulkOrderDetails *iso20022.SubscriptionBulkOrder4 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (s *SubscriptionBulkOrderV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderV03) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV03) AddBulkOrderDetails() *iso20022.SubscriptionBulkOrder4 {
	s.BulkOrderDetails = new(iso20022.SubscriptionBulkOrder4)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderV03) AddRelatedPartyDetails() *iso20022.Intermediary8 {
	newValue := new (iso20022.Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV03) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

