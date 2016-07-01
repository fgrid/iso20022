package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.04 Document"`
	Message *MeetingNotificationV04 `xml:"MtgNtfctn"`
}

func (d *Document00100104) AddMessage() *MeetingNotificationV04 {
	d.Message = new(MeetingNotificationV04)
	return d.Message
}

// Scope
// A notifying party, eg, an issuer, its agent or an intermediary, sends the MeetingNotification message to a party holding the right to vote, to announce a shareholders meeting.
// Usage
// The MeetingNotification message is used to announce a shareholders meeting, for example, it provides information on the participation details and requirements for the meeting, the vote parameters and the resolutions. The MeetingNotification message may also be used to announce an update.
// To notify an update, the Amendment building block must be filled in. Any building block that is modified must be included in the amendment message. The information previously notified and not repeated in the amendment message remains valid.
// To update the resolutions of the agenda, the complete list of resolutions must be repeated in the amendment message. The resolutions that are deleted should be assigned the status Withdrawn.
type MeetingNotificationV04 struct {

	// Identifies the meeting notification message.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Information specific to an amendment.
	Amendment *iso20022.AmendInformation1 `xml:"Amdmnt,omitempty"`

	// Defines the global status of the event contained in the notification.
	NotificationStatus *iso20022.NotificationStatus1 `xml:"NtfctnSts"`

	// Specifies information about the meeting. This component contains meeting identifications, various deadlines, contact persons, electronic and postal locations for accessing information and proxy assignment parameters.
	Meeting *iso20022.MeetingNotice3 `xml:"Mtg"`

	// Dates and details of the shareholders meeting.
	MeetingDetails []*iso20022.Meeting3 `xml:"MtgDtls"`

	// Party notifying the meeting.
	NotifyingParty *iso20022.PartyIdentification9Choice `xml:"NtifngPty"`

	// Specifies the institution that is the issuer of the security to which the meeting applies.
	Issuer *iso20022.IssuerInformation1 `xml:"Issr"`

	// Agents of the issuer.
	IssuerAgent []*iso20022.IssuerAgent1 `xml:"IssrAgt,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	Security []*iso20022.SecurityPosition6 `xml:"Scty"`

	// Detailed information of a resolution proposed to the vote.
	Resolution []*iso20022.Resolution2 `xml:"Rsltn,omitempty"`

	// Specifies the conditions to be allowed to vote, the different voting methods and options, the voting deadlines and the parameters of the incentive premium.
	Vote *iso20022.VoteParameters3 `xml:"Vote,omitempty"`

	// Specifies the entitlement ratio and the different deadlines for calculating the entitlement.
	EntitlementSpecification *iso20022.EntitlementAssessment2 `xml:"EntitlmntSpcfctn"`

	// Specifies requirements relative to the use of Power of Attorney.
	PowerOfAttorneyRequirements *iso20022.PowerOfAttorneyRequirements2 `xml:"PwrOfAttnyRqrmnts,omitempty"`

	// Provides additional narrative information about the corporate event.
	AdditionalInformation *iso20022.CorporateEventNarrative2 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`

}


func (m *MeetingNotificationV04) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingNotificationV04) AddAmendment() *iso20022.AmendInformation1 {
	m.Amendment = new(iso20022.AmendInformation1)
	return m.Amendment
}

func (m *MeetingNotificationV04) AddNotificationStatus() *iso20022.NotificationStatus1 {
	m.NotificationStatus = new(iso20022.NotificationStatus1)
	return m.NotificationStatus
}

func (m *MeetingNotificationV04) AddMeeting() *iso20022.MeetingNotice3 {
	m.Meeting = new(iso20022.MeetingNotice3)
	return m.Meeting
}

func (m *MeetingNotificationV04) AddMeetingDetails() *iso20022.Meeting3 {
	newValue := new (iso20022.Meeting3)
	m.MeetingDetails = append(m.MeetingDetails, newValue)
	return newValue
}

func (m *MeetingNotificationV04) AddNotifyingParty() *iso20022.PartyIdentification9Choice {
	m.NotifyingParty = new(iso20022.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingNotificationV04) AddIssuer() *iso20022.IssuerInformation1 {
	m.Issuer = new(iso20022.IssuerInformation1)
	return m.Issuer
}

func (m *MeetingNotificationV04) AddIssuerAgent() *iso20022.IssuerAgent1 {
	newValue := new (iso20022.IssuerAgent1)
	m.IssuerAgent = append(m.IssuerAgent, newValue)
	return newValue
}

func (m *MeetingNotificationV04) AddSecurity() *iso20022.SecurityPosition6 {
	newValue := new (iso20022.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingNotificationV04) AddResolution() *iso20022.Resolution2 {
	newValue := new (iso20022.Resolution2)
	m.Resolution = append(m.Resolution, newValue)
	return newValue
}

func (m *MeetingNotificationV04) AddVote() *iso20022.VoteParameters3 {
	m.Vote = new(iso20022.VoteParameters3)
	return m.Vote
}

func (m *MeetingNotificationV04) AddEntitlementSpecification() *iso20022.EntitlementAssessment2 {
	m.EntitlementSpecification = new(iso20022.EntitlementAssessment2)
	return m.EntitlementSpecification
}

func (m *MeetingNotificationV04) AddPowerOfAttorneyRequirements() *iso20022.PowerOfAttorneyRequirements2 {
	m.PowerOfAttorneyRequirements = new(iso20022.PowerOfAttorneyRequirements2)
	return m.PowerOfAttorneyRequirements
}

func (m *MeetingNotificationV04) AddAdditionalInformation() *iso20022.CorporateEventNarrative2 {
	m.AdditionalInformation = new(iso20022.CorporateEventNarrative2)
	return m.AdditionalInformation
}

func (m *MeetingNotificationV04) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}

