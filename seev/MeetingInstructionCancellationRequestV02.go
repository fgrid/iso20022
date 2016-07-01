package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.005.001.02 Document"`
	Message *MeetingInstructionCancellationRequestV02 `xml:"MtgInstrCxlReq"`
}

func (d *Document00500102) AddMessage() *MeetingInstructionCancellationRequestV02 {
	d.Message = new(MeetingInstructionCancellationRequestV02)
	return d.Message
}

// Scope
// The MeetingInstructionCancellationRequest message is sent by the same party that sent the MeetingInstruction message. It is sent to request the cancellation of all instructions included in the original the MeetingInstruction message.
// Usage
// This message must be answered by a MeetingInstructionStatus message. Some instructions in the previously sent MeetingInstruction message may have already been executed and cannot be cancelled. This information should appear clearly in the status message.
type MeetingInstructionCancellationRequestV02 struct {

	// Uniquely identifies the cancellation request.
	InstructionCancellationIdentification *iso20022.MessageIdentification1 `xml:"InstrCxlId"`

	// Identifies the instruction to be cancelled.
	PreviousReference *iso20022.MessageIdentification `xml:"PrvsRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference3 `xml:"MtgRef,omitempty"`

	// Party requesting the cancellation.
	RequestingParty *iso20022.PartyIdentification9Choice `xml:"RqstngPty,omitempty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *iso20022.SecurityIdentification3 `xml:"SctyId,omitempty"`

	// Identifies the account and instructed positions for which the instruction cancellation request applies.
	InstructedPosition []*iso20022.SafekeepingAccount3 `xml:"InstdPos,omitempty"`

}


func (m *MeetingInstructionCancellationRequestV02) AddInstructionCancellationIdentification() *iso20022.MessageIdentification1 {
	m.InstructionCancellationIdentification = new(iso20022.MessageIdentification1)
	return m.InstructionCancellationIdentification
}

func (m *MeetingInstructionCancellationRequestV02) AddPreviousReference() *iso20022.MessageIdentification {
	m.PreviousReference = new(iso20022.MessageIdentification)
	return m.PreviousReference
}

func (m *MeetingInstructionCancellationRequestV02) AddMeetingReference() *iso20022.MeetingReference3 {
	m.MeetingReference = new(iso20022.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingInstructionCancellationRequestV02) AddRequestingParty() *iso20022.PartyIdentification9Choice {
	m.RequestingParty = new(iso20022.PartyIdentification9Choice)
	return m.RequestingParty
}

func (m *MeetingInstructionCancellationRequestV02) AddSecurityIdentification() *iso20022.SecurityIdentification3 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingInstructionCancellationRequestV02) AddInstructedPosition() *iso20022.SafekeepingAccount3 {
	newValue := new (iso20022.SafekeepingAccount3)
	m.InstructedPosition = append(m.InstructedPosition, newValue)
	return newValue
}

