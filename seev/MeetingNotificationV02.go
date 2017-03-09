package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.02 Document"`
	Message *MeetingNotificationV02 `xml:"MtgNtfctn"`
}

func (d *Document00100102) AddMessage() *MeetingNotificationV02 {
	d.Message = new(MeetingNotificationV02)
	return d.Message
}

// Scope
// A notifying party, eg, an issuer, its agent or an intermediary, sends the MeetingNotification message to a party holding the right to vote, to announce a shareholders meeting.
// Usage
// The MeetingNotification message is used to announce a shareholders meeting, ie, it provides information on the participation details and requirements for the meeting, the vote parameters and the resolutions. The MeetingNotification message may also be used to announce an update.
// To notify an update, the Amendment building block must be filled in. Any building block that is modified must be included in the amendment message. The information previously notified and not repeated in the amendment message remains valid.
// To update the resolutions of the agenda, the complete list of resolutions must be repeated in the amendment message. The resolutions that are deleted should be assigned the status Withdrawn.
type MeetingNotificationV02 struct {

	// Identifies the meeting notification message.
	MeetingNotificationIdentification *iso20022.MessageIdentification1 `xml:"MtgNtfctnId"`

	// Information specific to an amendment.
	Amendment *iso20022.AmendInformation1 `xml:"Amdmnt,omitempty"`

	// Defines the global status of the event contained in the notification.
	NotificationStatus *iso20022.NotificationStatus1 `xml:"NtfctnSts"`

	// Specifies information about the meeting. This component contains meeting identifications, various deadlines, contact persons, electronic and postal locations for accessing information and proxy assignment parameters.
	Meeting *iso20022.MeetingNotice2 `xml:"Mtg"`

	// Dates and details of the shareholders meeting.
	MeetingDetails []*iso20022.Meeting2 `xml:"MtgDtls"`

	// Party notifying the meeting.
	NotifyingParty *iso20022.PartyIdentification9Choice `xml:"NtifngPty"`

	// Specifies the institution that is the issuer of the security to which the meeting applies.
	Issuer *iso20022.PartyIdentification9Choice `xml:"Issr"`

	// Agents of the issuer.
	IssuerAgent []*iso20022.IssuerAgent1 `xml:"IssrAgt,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	Security []*iso20022.SecurityPosition5 `xml:"Scty"`

	// Detailed information of a resolution proposed to the vote.
	Resolution []*iso20022.Resolution2 `xml:"Rsltn,omitempty"`

	// Specifies the conditions to be allowed to vote, the different voting methods and options, the voting deadlines and the parameters of the incentive premium.
	Vote *iso20022.VoteParameters1 `xml:"Vote,omitempty"`

	// Specifies the entitlement ratio and the different deadlines for calculating the entitlement.
	EntitlementSpecification *iso20022.EntitlementAssessment1 `xml:"EntitlmntSpcfctn"`

	// Specifies requirements relative to the use of Power of Attorney.
	PowerOfAttorneyRequirements *iso20022.PowerOfAttorneyRequirements2 `xml:"PwrOfAttnyRqrmnts,omitempty"`
}

func (m *MeetingNotificationV02) AddMeetingNotificationIdentification() *iso20022.MessageIdentification1 {
	m.MeetingNotificationIdentification = new(iso20022.MessageIdentification1)
	return m.MeetingNotificationIdentification
}

func (m *MeetingNotificationV02) AddAmendment() *iso20022.AmendInformation1 {
	m.Amendment = new(iso20022.AmendInformation1)
	return m.Amendment
}

func (m *MeetingNotificationV02) AddNotificationStatus() *iso20022.NotificationStatus1 {
	m.NotificationStatus = new(iso20022.NotificationStatus1)
	return m.NotificationStatus
}

func (m *MeetingNotificationV02) AddMeeting() *iso20022.MeetingNotice2 {
	m.Meeting = new(iso20022.MeetingNotice2)
	return m.Meeting
}

func (m *MeetingNotificationV02) AddMeetingDetails() *iso20022.Meeting2 {
	newValue := new(iso20022.Meeting2)
	m.MeetingDetails = append(m.MeetingDetails, newValue)
	return newValue
}

func (m *MeetingNotificationV02) AddNotifyingParty() *iso20022.PartyIdentification9Choice {
	m.NotifyingParty = new(iso20022.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingNotificationV02) AddIssuer() *iso20022.PartyIdentification9Choice {
	m.Issuer = new(iso20022.PartyIdentification9Choice)
	return m.Issuer
}

func (m *MeetingNotificationV02) AddIssuerAgent() *iso20022.IssuerAgent1 {
	newValue := new(iso20022.IssuerAgent1)
	m.IssuerAgent = append(m.IssuerAgent, newValue)
	return newValue
}

func (m *MeetingNotificationV02) AddSecurity() *iso20022.SecurityPosition5 {
	newValue := new(iso20022.SecurityPosition5)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingNotificationV02) AddResolution() *iso20022.Resolution2 {
	newValue := new(iso20022.Resolution2)
	m.Resolution = append(m.Resolution, newValue)
	return newValue
}

func (m *MeetingNotificationV02) AddVote() *iso20022.VoteParameters1 {
	m.Vote = new(iso20022.VoteParameters1)
	return m.Vote
}

func (m *MeetingNotificationV02) AddEntitlementSpecification() *iso20022.EntitlementAssessment1 {
	m.EntitlementSpecification = new(iso20022.EntitlementAssessment1)
	return m.EntitlementSpecification
}

func (m *MeetingNotificationV02) AddPowerOfAttorneyRequirements() *iso20022.PowerOfAttorneyRequirements2 {
	m.PowerOfAttorneyRequirements = new(iso20022.PowerOfAttorneyRequirements2)
	return m.PowerOfAttorneyRequirements
}
