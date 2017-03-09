package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02800101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.028.001.01 Document"`
	Message *AgentCADeactivationInstructionV01 `xml:"AgtCADeactvtnInstr"`
}

func (d *Document02800101) AddMessage() *AgentCADeactivationInstructionV01 {
	d.Message = new(AgentCADeactivationInstructionV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to the CSD to instruct the deactivation of a corporate action event or to deactivate one or more specific options of the corporate action. As of the deactivation date, the CSD is allowed to reject any related election instruction received from clients.
// Usage
// Deactivation refers only to the empowerment of the CSD to reject further elections. To withdraw an event, the Agent Corporate Action Notification Advice message must be used.
// This message can be used to deactivate all the options of a corporate action event, in which case, no option should be mentioned in the message.
// This message can also be used to deactivate one or more specific corporate action options, in which case, the option type and option number must be present.
// This message can only be used when the deactivation date is after the market deadline. Before the market deadline, an updated notification advice message must be sent with option availability status: inactive or cancelled.
// An un-effected deactivation (pending deactivation date/time) can be cancelled with an Agent Corporate Action Deactivation Cancellation Request.
// The amendment of a deactivation is effected by cancel/replace mechanism.
type AgentCADeactivationInstructionV01 struct {

	// Identification assigned by the Sender to unambiguously identify the instruction.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Information related to the deactivation of a CA event.
	DeactivationDetails *iso20022.CorporateActionDeactivationInstruction1 `xml:"DeactvtnDtls"`
}

func (a *AgentCADeactivationInstructionV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCADeactivationInstructionV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCADeactivationInstructionV01) AddDeactivationDetails() *iso20022.CorporateActionDeactivationInstruction1 {
	a.DeactivationDetails = new(iso20022.CorporateActionDeactivationInstruction1)
	return a.DeactivationDetails
}
