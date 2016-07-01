package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.03 Document"`
	Message *MeetingVoteExecutionConfirmationV03 `xml:"MtgVoteExctnConf"`
}

func (d *Document00700103) AddMessage() *MeetingVoteExecutionConfirmationV03 {
	d.Message = new(MeetingVoteExecutionConfirmationV03)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingVoteExecutionConfirmation message to confirm to the Sender of the MeetingInstruction message, the execution of their voting instruction.
// Usage
// This message is sent after the shareholders meeting has taken place. The Sender of this message confirms the execution of the vote at the meeting and confirms that the vote has been processed as instructed via the MeetingInstruction message.
// This messages is sent if the Sender of the MeetingInstruction message has requested such a confirmation or if market practice or regulation stipulates the need for a full audit trail.
type MeetingVoteExecutionConfirmationV03 struct {

	// Identifies the vote execution confirmation message.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Identifies the meeting instruction message.
	RelatedReference *iso20022.MessageIdentification `xml:"RltdRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef"`

	// Party confirming the votes.
	ReportingParty *iso20022.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised. 
	SecurityIdentification *iso20022.SecurityIdentification11 `xml:"SctyId"`

	// Specifies how a party has voted for each agenda item.
	VoteInstructions []*iso20022.DetailedInstructionStatus9 `xml:"VoteInstrs"`

}


func (m *MeetingVoteExecutionConfirmationV03) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingVoteExecutionConfirmationV03) AddRelatedReference() *iso20022.MessageIdentification {
	m.RelatedReference = new(iso20022.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingVoteExecutionConfirmationV03) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingVoteExecutionConfirmationV03) AddReportingParty() *iso20022.PartyIdentification9Choice {
	m.ReportingParty = new(iso20022.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingVoteExecutionConfirmationV03) AddSecurityIdentification() *iso20022.SecurityIdentification11 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingVoteExecutionConfirmationV03) AddVoteInstructions() *iso20022.DetailedInstructionStatus9 {
	newValue := new (iso20022.DetailedInstructionStatus9)
	m.VoteInstructions = append(m.VoteInstructions, newValue)
	return newValue
}

