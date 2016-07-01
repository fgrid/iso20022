package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.015.001.02 Document"`
	Message *SwitchOrderConfirmationV02 `xml:"setr.015.001.02"`
}

func (d *Document01500102) AddMessage() *SwitchOrderConfirmationV02 {
	d.Message = new(SwitchOrderConfirmationV02)
	return d.Message
}

// Scope
// The SwitchOrderConfirmation message is sent by an executing party, eg, a transfer agent, to an instruction party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to confirm the details of the execution of a SwitchOrder message.
// Usage
// The SwitchOrderConfirmation message is sent to confirm that all the legs of the switch have been executed.
type SwitchOrderConfirmationV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef"`

	// Information related to a switch execution.
	SwitchExecutionDetails *iso20022.SwitchExecution3 `xml:"SwtchExctnDtls"`

	// Confirmation of the information related to an intermediary.
	IntermediaryDetails []*iso20022.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (s *SwitchOrderConfirmationV02) AddMasterReference() *iso20022.AdditionalReference3 {
	s.MasterReference = new(iso20022.AdditionalReference3)
	return s.MasterReference
}

func (s *SwitchOrderConfirmationV02) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	s.RelatedReference = new(iso20022.AdditionalReference3)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationV02) AddSwitchExecutionDetails() *iso20022.SwitchExecution3 {
	s.SwitchExecutionDetails = new(iso20022.SwitchExecution3)
	return s.SwitchExecutionDetails
}

func (s *SwitchOrderConfirmationV02) AddIntermediaryDetails() *iso20022.Intermediary4 {
	newValue := new (iso20022.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV02) AddCopyDetails() *iso20022.CopyInformation1 {
	s.CopyDetails = new(iso20022.CopyInformation1)
	return s.CopyDetails
}

func (s *SwitchOrderConfirmationV02) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

