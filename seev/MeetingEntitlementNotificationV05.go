package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300105 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.003.001.05 Document"`
	Message *MeetingEntitlementNotificationV05 `xml:"MtgEntitlmntNtfctn"`
}

func (d *Document00300105) AddMessage() *MeetingEntitlementNotificationV05 {
	d.Message = new(MeetingEntitlementNotificationV05)
	return d.Message
}

// Scope
// An account servicer sends the MeetingEntitlementNotification to an issuer, its agent, an intermediary or an account owner to advise the entitlement in relation to a shareholders meeting.
// Usage
// This message is sent to advise the quantity of securities held by an account owner. The balance is specified for the securities for which the meeting is taking place.
// This entitlement message is sent by the account servicer or the registrar to an intermediary, the issuer's agent or the issuer. It is also sent between the account servicer and the account owner or the party holding the right to vote.
// The message is also used to amend a previously sent MeetingEntitlementNotification. To notify an update, the RelatedReference must be included in the message.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingEntitlementNotificationV05 struct {

	// Identifies the meeting entitlement message to be modified.
	RelatedReference *iso20022.MessageIdentification `xml:"RltdRef,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference7 `xml:"MtgRef"`

	// Identifies the security for which the meeting is organised, the account and the positions of the security holder.
	Security []*iso20022.SecurityPosition9 `xml:"Scty"`

	// Defines the dates determining eligibility.
	Eligibility *iso20022.EligibilityDates1 `xml:"Elgblty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingEntitlementNotificationV05) AddRelatedReference() *iso20022.MessageIdentification {
	m.RelatedReference = new(iso20022.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingEntitlementNotificationV05) AddMeetingReference() *iso20022.MeetingReference7 {
	m.MeetingReference = new(iso20022.MeetingReference7)
	return m.MeetingReference
}

func (m *MeetingEntitlementNotificationV05) AddSecurity() *iso20022.SecurityPosition9 {
	newValue := new(iso20022.SecurityPosition9)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingEntitlementNotificationV05) AddEligibility() *iso20022.EligibilityDates1 {
	m.Eligibility = new(iso20022.EligibilityDates1)
	return m.Eligibility
}

func (m *MeetingEntitlementNotificationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
