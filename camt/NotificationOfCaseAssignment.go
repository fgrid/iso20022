package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03000101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.01 Document"`
	Message *NotificationOfCaseAssignment `xml:"camt.030.001.01"`
}

func (d *Document03000101) AddMessage() *NotificationOfCaseAssignment {
	d.Message = new(NotificationOfCaseAssignment)
	return d.Message
}

// Scope
// The Notification Of Case Assignment message is sent by a case assignee to a case creator/case assigner.
// This message is used to inform the case assigner that:
// - the assignee is reassigning the case to the next agent in the transaction processing chain for further action
// - the assignee will work on the case himself, without re-assigning it to another party, and therefore indicating that the re-assignment has reached its end-point
// Usage
// The Notification Of Case Assignment message is used to notify the case creator or case assigner of further action undertaken by the case assignee in a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// The Notification Of Case Assignment message
// - covers one and only one case at a time. If the case assignee needs to inform a case creator or case assigner about several cases, then multiple Notification Of Case Assignment messages must be sent
// - except in the case where it is used to indicate that an agent is doing the correction himself, this message must be forwarded by all subsequent case assigner(s) until it reaches the case creator
// - must not be used in place of a Resolution Of Investigation or a Case Status Report message.
type NotificationOfCaseAssignment struct {

	// Specifies generic information about the notification.
	// The receiver of a notification is necessarily the party which assigned the case to the sender.
	Header *iso20022.ReportHeader `xml:"Hdr"`

	// Identifies the case.
	Case *iso20022.Case `xml:"Case"`

	// Identifies the current assignment of the case. 
	// 
	// The Assigner must be the emitter of the notification.
	// The Assignee must be another Party in the payment chain.
	Assignment *iso20022.CaseAssignment `xml:"Assgnmt"`

	// Information about the type of action taken.
	Notification *iso20022.CaseForwardingNotification `xml:"Ntfctn"`

}


func (n *NotificationOfCaseAssignment) AddHeader() *iso20022.ReportHeader {
	n.Header = new(iso20022.ReportHeader)
	return n.Header
}

func (n *NotificationOfCaseAssignment) AddCase() *iso20022.Case {
	n.Case = new(iso20022.Case)
	return n.Case
}

func (n *NotificationOfCaseAssignment) AddAssignment() *iso20022.CaseAssignment {
	n.Assignment = new(iso20022.CaseAssignment)
	return n.Assignment
}

func (n *NotificationOfCaseAssignment) AddNotification() *iso20022.CaseForwardingNotification {
	n.Notification = new(iso20022.CaseForwardingNotification)
	return n.Notification
}

