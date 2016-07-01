package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.014.001.03 Document"`
	Message *SwitchOrderCancellationRequestV03 `xml:"SwtchOrdrCxlReqV03"`
}

func (d *Document01400103) AddMessage() *SwitchOrderCancellationRequestV03 {
	d.Message = new(SwitchOrderCancellationRequestV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the SwitchOrderCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent SwitchOrder instruction.
// Usage
// The SwitchOrderCancellationRequest is used to cancel the entire previously sent SwitchOrder instruction and all the individual legs that it contains. There is no amendment, but a cancellation and re-instruct policy.
// There are two ways to specify the switch cancellation. Either:
// - the order reference of the original switch order is quoted, or,
// - all the details of the original switch order (this includes the OrderReference) are quoted, but this is not recommended.
// The message identification of the SwitchOrder message may also be quoted in PreviousReference.
// It is also possible to request the cancellation of a SwitchOrder message by quoting its message identification in PreviousReference, but this is not recommended.
// The deadline and acceptance of a cancellation request is subject to a service level agreement (SLA). This cancellation message is a cancellation request. There is no automatic acceptance of the cancellation request.
// The rejection or acceptance of a SwitchOrderCancellationRequest is made using an OrderCancellationStatusReport message.
type SwitchOrderCancellationRequestV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// References of the switch orders to be cancelled.
	CancellationByReference *iso20022.InvestmentFundOrder1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the switch orders to be cancelled.
	CancellationByOrderDetails *iso20022.SwitchOrderInstruction2 `xml:"CxlByOrdrDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

}


func (s *SwitchOrderCancellationRequestV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderCancellationRequestV03) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderCancellationRequestV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	s.PreviousReference = new(iso20022.AdditionalReference3)
	return s.PreviousReference
}

func (s *SwitchOrderCancellationRequestV03) AddCancellationByReference() *iso20022.InvestmentFundOrder1 {
	s.CancellationByReference = new(iso20022.InvestmentFundOrder1)
	return s.CancellationByReference
}

func (s *SwitchOrderCancellationRequestV03) AddCancellationByOrderDetails() *iso20022.SwitchOrderInstruction2 {
	s.CancellationByOrderDetails = new(iso20022.SwitchOrderInstruction2)
	return s.CancellationByOrderDetails
}

func (s *SwitchOrderCancellationRequestV03) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

