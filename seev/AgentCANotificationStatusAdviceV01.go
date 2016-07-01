package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.011.001.01 Document"`
	Message *AgentCANotificationStatusAdviceV01 `xml:"AgtCANtfctnStsAdvc"`
}

func (d *Document01100101) AddMessage() *AgentCANotificationStatusAdviceV01 {
	d.Message = new(AgentCANotificationStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to report the status, or change in status, of a notification advice or notification cancellation request.
// Usage
// When this message is used to report the status of a notification advice then the building block Agent Corporate Action Notification Advice Identification must be present.
// When this message is used to provide the status of a notification cancellation request then the building block Notification Cancellation Request Identification must be present.
type AgentCANotificationStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Notification Advice for which a status is given.
	AgentCANotificationAdviceIdentification *iso20022.DocumentIdentification8 `xml:"AgtCANtfctnAdvcId"`

	// Identification of the linked Agent CA Notification Cancellation Request for which a status is given.
	AgentCANotificationCancellationRequestIdentification *iso20022.DocumentIdentification8 `xml:"AgtCANtfctnCxlReqId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation2 `xml:"CorpActnGnlInf"`

	// Status of the Notification Cancellation Request sent by the issuer (agent).
	NotificationCancellationRequestStatus *iso20022.NotificationCancellationRequestStatus1Choice `xml:"NtfctnCxlReqSts"`

	// Status of the notification advice sent by the issuer (agent).
	NotificationAdviceStatus *iso20022.NotificationAdviceStatus1Choice `xml:"NtfctnAdvcSts"`

}


func (a *AgentCANotificationStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCANotificationStatusAdviceV01) AddAgentCANotificationAdviceIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCANotificationAdviceIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCANotificationAdviceIdentification
}

func (a *AgentCANotificationStatusAdviceV01) AddAgentCANotificationCancellationRequestIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCANotificationCancellationRequestIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCANotificationCancellationRequestIdentification
}

func (a *AgentCANotificationStatusAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation2 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation2)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCANotificationStatusAdviceV01) AddNotificationCancellationRequestStatus() *iso20022.NotificationCancellationRequestStatus1Choice {
	a.NotificationCancellationRequestStatus = new(iso20022.NotificationCancellationRequestStatus1Choice)
	return a.NotificationCancellationRequestStatus
}

func (a *AgentCANotificationStatusAdviceV01) AddNotificationAdviceStatus() *iso20022.NotificationAdviceStatus1Choice {
	a.NotificationAdviceStatus = new(iso20022.NotificationAdviceStatus1Choice)
	return a.NotificationAdviceStatus
}

