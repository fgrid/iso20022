package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500103 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.005.001.03 Document"`
	Message *MeetingInstructionCancellationRequestV03 `xml:"MtgInstrCxlReq"`
}

func (d *Document00500103) AddMessage() *MeetingInstructionCancellationRequestV03 {
	d.Message = new(MeetingInstructionCancellationRequestV03)
	return d.Message
}

// Scope
// The MeetingInstructionCancellationRequest message is sent by the same party that sent the MeetingInstruction message. It is sent to request the cancellation of all instructions included in the original the MeetingInstruction message.
// Usage
// This message must be answered by a MeetingInstructionStatus message. Some instructions in the previously sent MeetingInstruction message may have already been executed and cannot be cancelled. This information should appear clearly in the status message.
type MeetingInstructionCancellationRequestV03 struct {

	// Uniquely identifies the cancellation request.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Identifies the instruction to be cancelled.
	PreviousReference *iso20022.MessageIdentification `xml:"PrvsRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef,omitempty"`

	// Party requesting the cancellation.
	RequestingParty *iso20022.PartyIdentification9Choice `xml:"RqstngPty,omitempty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *iso20022.SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Identifies the account and instructed positions for which the instruction cancellation request applies.
	InstructedPosition []*iso20022.SafekeepingAccount4 `xml:"InstdPos,omitempty"`
}

func (m *MeetingInstructionCancellationRequestV03) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingInstructionCancellationRequestV03) AddPreviousReference() *iso20022.MessageIdentification {
	m.PreviousReference = new(iso20022.MessageIdentification)
	return m.PreviousReference
}

func (m *MeetingInstructionCancellationRequestV03) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingInstructionCancellationRequestV03) AddRequestingParty() *iso20022.PartyIdentification9Choice {
	m.RequestingParty = new(iso20022.PartyIdentification9Choice)
	return m.RequestingParty
}

func (m *MeetingInstructionCancellationRequestV03) AddSecurityIdentification() *iso20022.SecurityIdentification11 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingInstructionCancellationRequestV03) AddInstructedPosition() *iso20022.SafekeepingAccount4 {
	newValue := new(iso20022.SafekeepingAccount4)
	m.InstructedPosition = append(m.InstructedPosition, newValue)
	return newValue
}
