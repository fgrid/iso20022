package admi

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01 Document"`
	Message *MessageRejectV01 `xml:"admi.002.001.01"`
}

func (d *Document00200101) AddMessage() *MessageRejectV01 {
	d.Message = new(MessageRejectV01)
	return d.Message
}

// Scope
// The MessageReject message is sent by a central system to notify the rejection of a previously received message.
// Usage
// The message provides specific information about the rejection reason.
type MessageRejectV01 struct {

	// Refers to the identification of the message previously received and for which the rejection is notified.
	RelatedReference *iso20022.MessageReference `xml:"RltdRef"`

	// General information about the reason of the message rejection.
	Reason *iso20022.RejectionReason2 `xml:"Rsn"`
}

func (m *MessageRejectV01) AddRelatedReference() *iso20022.MessageReference {
	m.RelatedReference = new(iso20022.MessageReference)
	return m.RelatedReference
}

func (m *MessageRejectV01) AddReason() *iso20022.RejectionReason2 {
	m.Reason = new(iso20022.RejectionReason2)
	return m.Reason
}
