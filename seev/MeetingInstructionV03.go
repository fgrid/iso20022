package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.004.001.03 Document"`
	Message *MeetingInstructionV03 `xml:"MtgInstr"`
}

func (d *Document00400103) AddMessage() *MeetingInstructionV03 {
	d.Message = new(MeetingInstructionV03)
	return d.Message
}

// Scope
// A party holding the right to vote sends the MeetingInstruction message to an intermediary, the issuer or its agent to request the receiving party to act upon one or several instructions.
// Usage
// The MeetingInstruction message is used to register for a shareholders meeting, request blocking or registration of securities. It is used to assign a proxy, to specify the names of meeting attendees and to relay vote instructions per resolution electronically.
// The MeetingInstruction message may only be sent for one security, though several safekeeping places may be specified.
// Once the message is sent, it cannot be modified. It must be cancelled by a MeetingInstructionCancellationRequest. Only after receipt of a confirmed cancelled status via the MeetingInstructionStatus message, a new MeetingInstruction message can be sent.
type MeetingInstructionV03 struct {

	// Identifies the meeting instruction message. 
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef"`

	// Party notifying the instructions.
	InstructingParty *iso20022.PartyIdentification9Choice `xml:"InstgPty"`

	// Identifies the security for which the meeting is organised.
	SecurityIdentification *iso20022.SecurityIdentification11 `xml:"SctyId"`

	// Identifies the position of the instructing party and the action that they want to take.
	Instruction []*iso20022.Instruction2 `xml:"Instr"`

}


func (m *MeetingInstructionV03) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingInstructionV03) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingInstructionV03) AddInstructingParty() *iso20022.PartyIdentification9Choice {
	m.InstructingParty = new(iso20022.PartyIdentification9Choice)
	return m.InstructingParty
}

func (m *MeetingInstructionV03) AddSecurityIdentification() *iso20022.SecurityIdentification11 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingInstructionV03) AddInstruction() *iso20022.Instruction2 {
	newValue := new (iso20022.Instruction2)
	m.Instruction = append(m.Instruction, newValue)
	return newValue
}

