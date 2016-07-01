package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.02 Document"`
	Message *MeetingVoteExecutionConfirmationV02 `xml:"MtgVoteExctnConf"`
}

func (d *Document00700102) AddMessage() *MeetingVoteExecutionConfirmationV02 {
	d.Message = new(MeetingVoteExecutionConfirmationV02)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingVoteExecutionConfirmation message to confirm to the Sender of the MeetingInstruction message, the execution of their voting instruction.
// Usage
// This message is sent after the shareholders meeting has taken place. The Sender of this message confirms the execution of the vote at the meeting and confirms that the vote has been processed as instructed via the MeetingInstruction message.
// This messages is sent if the Sender of the MeetingInstruction message has requested such a confirmation or if market practice or regulation stipulates the need for a full audit trail.
type MeetingVoteExecutionConfirmationV02 struct {

	// Identifies the vote execution confirmation message.
	VoteExecutionConfirmationIdentification *iso20022.MessageIdentification1 `xml:"VoteExctnConfId"`

	// Identifies the meeting instruction message.
	RelatedReference *iso20022.MessageIdentification `xml:"RltdRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference3 `xml:"MtgRef"`

	// Party confirming the votes.
	ReportingParty *iso20022.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised. 
	SecurityIdentification *iso20022.SecurityIdentification3 `xml:"SctyId"`

	// Specifies how a party has voted for each agenda item.
	VoteInstruction []*iso20022.DetailedInstructionStatus2 `xml:"VoteInstr"`

}


func (m *MeetingVoteExecutionConfirmationV02) AddVoteExecutionConfirmationIdentification() *iso20022.MessageIdentification1 {
	m.VoteExecutionConfirmationIdentification = new(iso20022.MessageIdentification1)
	return m.VoteExecutionConfirmationIdentification
}

func (m *MeetingVoteExecutionConfirmationV02) AddRelatedReference() *iso20022.MessageIdentification {
	m.RelatedReference = new(iso20022.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingVoteExecutionConfirmationV02) AddMeetingReference() *iso20022.MeetingReference3 {
	m.MeetingReference = new(iso20022.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingVoteExecutionConfirmationV02) AddReportingParty() *iso20022.PartyIdentification9Choice {
	m.ReportingParty = new(iso20022.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingVoteExecutionConfirmationV02) AddSecurityIdentification() *iso20022.SecurityIdentification3 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification3)
	return m.SecurityIdentification
}

func (m *MeetingVoteExecutionConfirmationV02) AddVoteInstruction() *iso20022.DetailedInstructionStatus2 {
	newValue := new (iso20022.DetailedInstructionStatus2)
	m.VoteInstruction = append(m.VoteInstruction, newValue)
	return newValue
}

