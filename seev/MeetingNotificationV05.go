package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.05 Document"`
	Message *MeetingNotificationV05 `xml:"MtgNtfctn"`
}

func (d *Document00100105) AddMessage() *MeetingNotificationV05 {
	d.Message = new(MeetingNotificationV05)
	return d.Message
}

// Scope
// A notifying party, for example, an issuer, its agent or an intermediary, sends the MeetingNotification message to a party holding the right to vote, to announce a shareholders meeting.
// Usage
// The MeetingNotification message is used to announce a shareholders meeting, for example, it provides information on the participation details and requirements for the meeting, the vote parameters and the resolutions. The MeetingNotification message may also be used to announce an update.
// To notify an update, the Amendment building block must be filled in. Any building block that is modified must be included in the amendment message. The information previously notified and not repeated in the amendment message remains valid.
// To update the resolutions of the agenda, the complete list of resolutions must be repeated in the amendment message. The resolutions that are deleted should be assigned the status Withdrawn.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingNotificationV05 struct {

	// Information specific to an amendment.
	Amendment *iso20022.AmendInformation1 `xml:"Amdmnt,omitempty"`

	// Defines the global status of the event contained in the notification.
	NotificationStatus *iso20022.NotificationStatus2 `xml:"NtfctnSts"`

	// Specifies information about the meeting. This component contains meeting identifications, various deadlines, contact persons, electronic and postal locations for accessing information and proxy assignment parameters.
	Meeting *iso20022.MeetingNotice4 `xml:"Mtg"`

	// Dates and details of the shareholders meeting.
	MeetingDetails []*iso20022.Meeting4 `xml:"MtgDtls"`

	// Specifies the institution that is the issuer of the security to which the meeting applies.
	Issuer *iso20022.IssuerInformation2 `xml:"Issr"`

	// Agents of the issuer.
	IssuerAgent []*iso20022.IssuerAgent2 `xml:"IssrAgt,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	Security []*iso20022.SecurityPosition8 `xml:"Scty"`

	// Detailed information of a resolution proposed to the vote.
	Resolution []*iso20022.Resolution3 `xml:"Rsltn,omitempty"`

	// Specifies the conditions to be allowed to vote, the different voting methods and options, the voting deadlines and the parameters of the incentive premium.
	Vote *iso20022.VoteParameters4 `xml:"Vote,omitempty"`

	// Specifies the entitlement ratio and the different deadlines for calculating the entitlement.
	EntitlementSpecification *iso20022.EntitlementAssessment3 `xml:"EntitlmntSpcfctn,omitempty"`

	// Specifies requirements relative to the use of Power of Attorney.
	PowerOfAttorneyRequirements *iso20022.PowerOfAttorneyRequirements3 `xml:"PwrOfAttnyRqrmnts,omitempty"`

	// Provides additional narrative information about the corporate event.
	AdditionalInformation *iso20022.CorporateEventNarrative2 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MeetingNotificationV05) AddAmendment() *iso20022.AmendInformation1 {
	m.Amendment = new(iso20022.AmendInformation1)
	return m.Amendment
}

func (m *MeetingNotificationV05) AddNotificationStatus() *iso20022.NotificationStatus2 {
	m.NotificationStatus = new(iso20022.NotificationStatus2)
	return m.NotificationStatus
}

func (m *MeetingNotificationV05) AddMeeting() *iso20022.MeetingNotice4 {
	m.Meeting = new(iso20022.MeetingNotice4)
	return m.Meeting
}

func (m *MeetingNotificationV05) AddMeetingDetails() *iso20022.Meeting4 {
	newValue := new (iso20022.Meeting4)
	m.MeetingDetails = append(m.MeetingDetails, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddIssuer() *iso20022.IssuerInformation2 {
	m.Issuer = new(iso20022.IssuerInformation2)
	return m.Issuer
}

func (m *MeetingNotificationV05) AddIssuerAgent() *iso20022.IssuerAgent2 {
	newValue := new (iso20022.IssuerAgent2)
	m.IssuerAgent = append(m.IssuerAgent, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddSecurity() *iso20022.SecurityPosition8 {
	newValue := new (iso20022.SecurityPosition8)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddResolution() *iso20022.Resolution3 {
	newValue := new (iso20022.Resolution3)
	m.Resolution = append(m.Resolution, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddVote() *iso20022.VoteParameters4 {
	m.Vote = new(iso20022.VoteParameters4)
	return m.Vote
}

func (m *MeetingNotificationV05) AddEntitlementSpecification() *iso20022.EntitlementAssessment3 {
	m.EntitlementSpecification = new(iso20022.EntitlementAssessment3)
	return m.EntitlementSpecification
}

func (m *MeetingNotificationV05) AddPowerOfAttorneyRequirements() *iso20022.PowerOfAttorneyRequirements3 {
	m.PowerOfAttorneyRequirements = new(iso20022.PowerOfAttorneyRequirements3)
	return m.PowerOfAttorneyRequirements
}

func (m *MeetingNotificationV05) AddAdditionalInformation() *iso20022.CorporateEventNarrative2 {
	m.AdditionalInformation = new(iso20022.CorporateEventNarrative2)
	return m.AdditionalInformation
}

func (m *MeetingNotificationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

