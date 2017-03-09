package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.002.001.02 Document"`
	Message *MeetingCancellationV02 `xml:"MtgCxl"`
}

func (d *Document00200102) AddMessage() *MeetingCancellationV02 {
	d.Message = new(MeetingCancellationV02)
	return d.Message
}

// Scope
// The MeetingCancellation message is sent by the party that sent the MeetingNotification message to the original receiver. It is sent to cancel the previous MeetingNotification message or to advise the cancellation of a meeting.
// Usage
// The MeetingCancellation message is used in two different situations.
// First, it is used to cancel a previously sent MeetingNotification message. In this case, the MessageCancellation, the MeetingReference and the Reason building blocks need to be present.
// Second, it is used to advise that the meeting is cancelled. In this case, only the MeetingReference and Reason building blocks need to be present.
type MeetingCancellationV02 struct {

	// Identifies the cancellation message.
	CancellationIdentification *iso20022.MessageIdentification1 `xml:"CxlId"`

	// Information indicating that the cancellation of a message previously sent is requested (and not the cancellation of the meeting).
	MessageCancellation *iso20022.AmendInformation1 `xml:"MsgCxl,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference2 `xml:"MtgRef"`

	// Party notifying the cancellation of the meeting.
	NotifyingParty *iso20022.PartyIdentification9Choice `xml:"NtifngPty,omitempty"`

	// Identifies the security for which the meeting was organised.
	Security []*iso20022.SecurityPosition5 `xml:"Scty,omitempty"`

	// Defines the justification for the cancellation.
	Reason *iso20022.MeetingCancellationReason1 `xml:"Rsn"`
}

func (m *MeetingCancellationV02) AddCancellationIdentification() *iso20022.MessageIdentification1 {
	m.CancellationIdentification = new(iso20022.MessageIdentification1)
	return m.CancellationIdentification
}

func (m *MeetingCancellationV02) AddMessageCancellation() *iso20022.AmendInformation1 {
	m.MessageCancellation = new(iso20022.AmendInformation1)
	return m.MessageCancellation
}

func (m *MeetingCancellationV02) AddMeetingReference() *iso20022.MeetingReference2 {
	m.MeetingReference = new(iso20022.MeetingReference2)
	return m.MeetingReference
}

func (m *MeetingCancellationV02) AddNotifyingParty() *iso20022.PartyIdentification9Choice {
	m.NotifyingParty = new(iso20022.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingCancellationV02) AddSecurity() *iso20022.SecurityPosition5 {
	newValue := new(iso20022.SecurityPosition5)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingCancellationV02) AddReason() *iso20022.MeetingCancellationReason1 {
	m.Reason = new(iso20022.MeetingCancellationReason1)
	return m.Reason
}
