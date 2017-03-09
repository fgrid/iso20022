package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800103 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.03 Document"`
	Message *MeetingResultDisseminationV03 `xml:"MtgRsltDssmntn"`
}

func (d *Document00800103) AddMessage() *MeetingResultDisseminationV03 {
	d.Message = new(MeetingResultDisseminationV03)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingResultDissemination message to another intermediary, to a party holding the right to vote, to a registered security holder or to a beneficial holder to provide information on the voting results of a shareholders meeting.
// Usage
// The MeetingResultDissemination message is used to provide the vote results per resolution. It may also provide information on the level of participation.
// This message is also used to notify an update or amendment to a previously sent MeetingResultDissemination message.
type MeetingResultDisseminationV03 struct {

	// Identifies the meeting dissemination notification message.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Information specific to an amemdment.
	Amendment *iso20022.AmendInformation2 `xml:"Amdmnt,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef"`

	// Party reporting the meeting results.
	ReportingParty *iso20022.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	Security []*iso20022.SecurityPosition6 `xml:"Scty"`

	// Results per resolution.
	VoteResult []*iso20022.Vote5 `xml:"VoteRslt"`

	// Information about the participation to the voting process.
	Participation *iso20022.Participation3 `xml:"Prtcptn,omitempty"`

	// Information on where additionnal information can be received.
	AdditionalInformation *iso20022.CommunicationAddress4 `xml:"AddtlInf,omitempty"`
}

func (m *MeetingResultDisseminationV03) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingResultDisseminationV03) AddAmendment() *iso20022.AmendInformation2 {
	m.Amendment = new(iso20022.AmendInformation2)
	return m.Amendment
}

func (m *MeetingResultDisseminationV03) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingResultDisseminationV03) AddReportingParty() *iso20022.PartyIdentification9Choice {
	m.ReportingParty = new(iso20022.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingResultDisseminationV03) AddSecurity() *iso20022.SecurityPosition6 {
	newValue := new(iso20022.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV03) AddVoteResult() *iso20022.Vote5 {
	newValue := new(iso20022.Vote5)
	m.VoteResult = append(m.VoteResult, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV03) AddParticipation() *iso20022.Participation3 {
	m.Participation = new(iso20022.Participation3)
	return m.Participation
}

func (m *MeetingResultDisseminationV03) AddAdditionalInformation() *iso20022.CommunicationAddress4 {
	m.AdditionalInformation = new(iso20022.CommunicationAddress4)
	return m.AdditionalInformation
}
