package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.010.001.03 Document"`
	Message *SubscriptionOrderV03 `xml:"SbcptOrdrV03"`
}

func (d *Document01000103) AddMessage() *SubscriptionOrderV03 {
	d.Message = new(SubscriptionOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the SubscriptionOrder message to the executing party, for example, a transfer agent, to instruct the subscription of one or more financial instruments for one investment fund account.
// Usage
// The SubscriptionOrder message is used to instruct single subscription orders, that is, a message containing one order for one financial instrument and related to one investment account. The SubscriptionOrder message may also be used for multiple orders, that is, a message containing several orders related to the same investment account for different financial instruments.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for the same financial instrument but for different accounts, then the SubscriptionBulkOrder message must be used.
type SubscriptionOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	MultipleOrderDetails *iso20022.SubscriptionMultipleOrder4 `xml:"MltplOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (s *SubscriptionOrderV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderV03) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionOrderV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderV03) AddMultipleOrderDetails() *iso20022.SubscriptionMultipleOrder4 {
	s.MultipleOrderDetails = new(iso20022.SubscriptionMultipleOrder4)
	return s.MultipleOrderDetails
}

func (s *SubscriptionOrderV03) AddRelatedPartyDetails() *iso20022.Intermediary8 {
	newValue := new (iso20022.Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrderV03) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionOrderV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

