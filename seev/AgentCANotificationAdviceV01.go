package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.009.001.01 Document"`
	Message *AgentCANotificationAdviceV01 `xml:"AgtCANtfctnAdvc"`
}

func (d *Document00900101) AddMessage() *AgentCANotificationAdviceV01 {
	d.Message = new(AgentCANotificationAdviceV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to:
// - Provide a CSD with the details of a corporate action along with the possible options available to the clients of that CSD; and
// - to update a corporate action notification. A notification advice can be initially sent as a preliminary advice and subsequently replaced by another notification advice with updated information.
// Usage
// This message is used:
// - to provide a CSD with the details of a corporate action along with the possible options available to the clients of that CSD. The information can be complete or incomplete.
// - to update a corporate action notification advice. A notification advice can be initially sent as a preliminary advice and subsequently replaced by another notification advice with updated information. As per SMPG recommendation, all the information should be provided in the update, not only updated information.
// An Agent Corporate Action Notification Status Advice is sent in reply to the Agent Corporate Action Notification Advice.
// Note: The amendment of a corporate action notification is done through a replacement mechanism in line with both the ISO 15022 messages used in the flow between the CSD and its clients, and the ISO 20022 proxy voting messages.
type AgentCANotificationAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Provides information about the type of  notification advice and linked message.
	NotificationTypeAndLinkage *iso20022.LinkedCorporateAction1 `xml:"NtfctnTpAndLkg"`

	// Provides general information about the notification advice.
	NotificationGeneralInformation *iso20022.CorporateActionNotification1 `xml:"NtfctnGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation2 `xml:"CorpActnGnlInf"`

	// Provides details information about the CA event.
	CorporateActionDetails *iso20022.CorporateAction2 `xml:"CorpActnDtls"`

	// Provides detailed information about the option of the CA event.
	CorporateActionOptionDetails []*iso20022.CorporateActionOption1 `xml:"CorpActnOptnDtls,omitempty"`

	// Provides information about the contact responsible for the transaction identified in the message.
	ContactDetails []*iso20022.ContactPerson1 `xml:"CtctDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative2 `xml:"AddtlInf,omitempty"`

}


func (a *AgentCANotificationAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCANotificationAdviceV01) AddNotificationTypeAndLinkage() *iso20022.LinkedCorporateAction1 {
	a.NotificationTypeAndLinkage = new(iso20022.LinkedCorporateAction1)
	return a.NotificationTypeAndLinkage
}

func (a *AgentCANotificationAdviceV01) AddNotificationGeneralInformation() *iso20022.CorporateActionNotification1 {
	a.NotificationGeneralInformation = new(iso20022.CorporateActionNotification1)
	return a.NotificationGeneralInformation
}

func (a *AgentCANotificationAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation2 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation2)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCANotificationAdviceV01) AddCorporateActionDetails() *iso20022.CorporateAction2 {
	a.CorporateActionDetails = new(iso20022.CorporateAction2)
	return a.CorporateActionDetails
}

func (a *AgentCANotificationAdviceV01) AddCorporateActionOptionDetails() *iso20022.CorporateActionOption1 {
	newValue := new (iso20022.CorporateActionOption1)
	a.CorporateActionOptionDetails = append(a.CorporateActionOptionDetails, newValue)
	return newValue
}

func (a *AgentCANotificationAdviceV01) AddContactDetails() *iso20022.ContactPerson1 {
	newValue := new (iso20022.ContactPerson1)
	a.ContactDetails = append(a.ContactDetails, newValue)
	return newValue
}

func (a *AgentCANotificationAdviceV01) AddAdditionalInformation() *iso20022.CorporateActionNarrative2 {
	a.AdditionalInformation = new(iso20022.CorporateActionNarrative2)
	return a.AdditionalInformation
}

