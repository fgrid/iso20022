package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.05 Document"`
	Message *MeetingResultDisseminationV05 `xml:"MtgRsltDssmntn"`
}

func (d *Document00800105) AddMessage() *MeetingResultDisseminationV05 {
	d.Message = new(MeetingResultDisseminationV05)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingResultDissemination message to another intermediary, to a party holding the right to vote, to a registered security holder or to a beneficial holder to provide information on the voting results of a shareholders meeting.
// Usage
// The MeetingResultDissemination message is used to provide the vote results per resolution. It may also provide information on the level of participation.
// This message is also used to notify an update or amendment to a previously sent MeetingResultDissemination message.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingResultDisseminationV05 struct {

	// Information specific to an amendment.
	Amendment *iso20022.AmendInformation3 `xml:"Amdmnt,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference7 `xml:"MtgRef"`

	// Identifies the securities for which the meeting is organised.
	Security []*iso20022.SecurityPosition8 `xml:"Scty"`

	// Results per resolution.
	VoteResult []*iso20022.Vote7 `xml:"VoteRslt"`

	// Information about the participation to the voting process.
	Participation *iso20022.Participation4 `xml:"Prtcptn,omitempty"`

	// Information on where additional information can be received.
	AdditionalInformation *iso20022.CommunicationAddress4 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MeetingResultDisseminationV05) AddAmendment() *iso20022.AmendInformation3 {
	m.Amendment = new(iso20022.AmendInformation3)
	return m.Amendment
}

func (m *MeetingResultDisseminationV05) AddMeetingReference() *iso20022.MeetingReference7 {
	m.MeetingReference = new(iso20022.MeetingReference7)
	return m.MeetingReference
}

func (m *MeetingResultDisseminationV05) AddSecurity() *iso20022.SecurityPosition8 {
	newValue := new (iso20022.SecurityPosition8)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV05) AddVoteResult() *iso20022.Vote7 {
	newValue := new (iso20022.Vote7)
	m.VoteResult = append(m.VoteResult, newValue)
	return newValue
}

func (m *MeetingResultDisseminationV05) AddParticipation() *iso20022.Participation4 {
	m.Participation = new(iso20022.Participation4)
	return m.Participation
}

func (m *MeetingResultDisseminationV05) AddAdditionalInformation() *iso20022.CommunicationAddress4 {
	m.AdditionalInformation = new(iso20022.CommunicationAddress4)
	return m.AdditionalInformation
}

func (m *MeetingResultDisseminationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

