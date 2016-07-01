package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.006.001.02 Document"`
	Message *MeetingInstructionStatusV02 `xml:"MtgInstrSts"`
}

func (d *Document00600102) AddMessage() *MeetingInstructionStatusV02 {
	d.Message = new(MeetingInstructionStatusV02)
	return d.Message
}

// Scope
// The Receiver of the MeetingInstruction or MeetingInstructionCancellationRequest sends the MeetingInstructionStatus message to the Sender of these messages.
// The message gives the status of a complete message or of one or more specific instructions within the message.
// Usage
// The MeetingInstructionStatus message is used for four purposes.
// First, it provides the status on the processing of a MeetingInstructionCancellationRequest message, ie, whether the request message is rejected or accepted.
// Second, it is used to provide a global processing or rejection status of a MeetingInstruction message.
// Third, it is used to provide a detailed processing or rejection status of a MeetingInstruction message, ie, for each instruction in the MeetingInstruction message the processing or rejection status is individually reported by using the InstructionIdentification element. This identification allows the receiver of the status message to link the status confirmation to its original instruction.
// The blocking of securities should be confirmed via an MT 508 (Intra-Position Advice).
// Fourth, it is used as a reminder to request voting instructions. This is done by indicating NONREF in the Identification element of the InstructionIdentification component and by using the status code NotReceived in the ProcessingStatus.
type MeetingInstructionStatusV02 struct {

	// Identifies the meeting instruction status message.
	MeetingInstructionStatusIdentification *iso20022.MessageIdentification1 `xml:"MtgInstrStsId"`

	// Identifies the meeting instruction message for which the status is provided.
	InstructionIdentification *iso20022.MessageIdentification `xml:"InstrId"`

	// Identifies the meeting instruction cancellation request message for which the status is provided.
	InstructionCancellationIdentification *iso20022.MessageIdentification `xml:"InstrCxlId"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference3 `xml:"MtgRef"`

	// Party reporting the status.
	ReportingParty *iso20022.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised. 
	SecurityIdentification *iso20022.SecurityIdentification3 `xml:"SctyId"`

	// Status applying to the instruction request received. The instruction is identified by the InstructionIdentification.
	InstructionStatus *iso20022.InstructionStatus1Choice `xml:"InstrSts"`

	// Status applying to the instruction cancellation request received. The instruction cancellation is identified by the InstructionCancellationIdentification.
	CancellationStatus *iso20022.CancellationStatus1Choice `xml:"CxlSts"`

}


func (m *MeetingInstructionStatusV02) AddMeetingInstructionStatusIdentification() *iso20022.MessageIdentification1 {
	m.MeetingInstructionStatusIdentification = new(iso20022.MessageIdentification1)
	return m.MeetingInstructionStatusIdentification
}

func (m *MeetingInstructionStatusV02) AddInstructionIdentification() *iso20022.MessageIdentification {
	m.InstructionIdentification = new(iso20022.MessageIdentification)
	return m.InstructionIdentification
}

func (m *MeetingInstructionStatusV02) AddInstructionCancellationIdentification() *iso20022.MessageIdentification {
	m.InstructionCancellationIdentification = new(iso20022.MessageIdentification)
	return m.InstructionCancellationIdentification
}

func (m *MeetingInstructionStatusV02) AddMeetingReference() *iso20022.MeetingReference3 {
	m.MeetingReference = new(iso20022.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingInstructionStatusV02) AddReportingParty() *iso20022.PartyIdentification9Choice {
	m.ReportingParty = new(iso20022.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingInstructionStatusV02) AddSecurityIdentification() *iso20022.SecurityIdentification3 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingInstructionStatusV02) AddInstructionStatus() *iso20022.InstructionStatus1Choice {
	m.InstructionStatus = new(iso20022.InstructionStatus1Choice)
	return m.InstructionStatus
}

func (m *MeetingInstructionStatusV02) AddCancellationStatus() *iso20022.CancellationStatus1Choice {
	m.CancellationStatus = new(iso20022.CancellationStatus1Choice)
	return m.CancellationStatus
}

