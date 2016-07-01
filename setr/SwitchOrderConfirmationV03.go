package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.015.001.03 Document"`
	Message *SwitchOrderConfirmationV03 `xml:"SwtchOrdrConfV03"`
}

func (d *Document01500103) AddMessage() *SwitchOrderConfirmationV03 {
	d.Message = new(SwitchOrderConfirmationV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SwitchOrderConfirmation message to the instructing party, for example, an investment manager or its authorised representative to confirm the details of the execution of a previously received SwitchOrder instruction.
// Usage
// The SwitchOrderConfirmation message is used to confirm that all the legs of the previously instructed switch order have been executed. The reference of the switch order confirmation is identified in DealReference.
// The reference of the original switch order is specified in OrderReference. The message identification of the SwitchOrder message in which the switch order was conveyed may also be quoted in RelatedReference.
// If the SwitchOrderConfirmation Details sequence is present more than once, then the redemption leg details and subscription leg details sequences may only be present once.
// If Switch Confirmation Order Details\Investment Account is used, then the Investment Account Details sequences in Subscription Leg Details and Redemption Leg Details are not allowed.
// If the SwitchOrderConfirmation Details sequence is present more than once, then the redemption leg details and subscription leg details sequences may only be present once.
// If SwitchOrderConfirmationDetails\InvestmentAccount is used, then the InvestmentAccountDetails sequences in SubscriptionLegDetails and RedemptionLegDetails are not allowed. This functionality is to be used by institutions that set up two accounts per investor, rather than one investment account.
// 
type SwitchOrderConfirmationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to a switch execution.
	SwitchExecutionDetails []*iso20022.SwitchExecution4 `xml:"SwtchExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (s *SwitchOrderConfirmationV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderConfirmationV03) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV03) AddRelatedReference() *iso20022.AdditionalReference3 {
	s.RelatedReference = new(iso20022.AdditionalReference3)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationV03) AddSwitchExecutionDetails() *iso20022.SwitchExecution4 {
	newValue := new (iso20022.SwitchExecution4)
	s.SwitchExecutionDetails = append(s.SwitchExecutionDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV03) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

func (s *SwitchOrderConfirmationV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

