package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02900101 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Document"`
	Message *AgentCADeactivationCancellationRequestV01 `xml:"AgtCADeactvtnCxlReq"`
}

func (d *Document02900101) AddMessage() *AgentCADeactivationCancellationRequestV01 {
	d.Message = new(AgentCADeactivationCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to the CSD to request the cancellation of a previously sent corporate action deactivation instruction.
// Usage
// This message is used to request the cancellation of a pending deactivation instruction. The cancellation must apply to exactly the same level as the original instruction, ie to the entire CA event or to an option as per the original instruction.
// This message must be sent before the deactivation execution date.
// In case a corporate action or option is already deactivated, this message can not be used to reactivate the corporate action entire event or option; the notification advice message must be used to reactivate a corporate action or option.
type AgentCADeactivationCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Deactivation Instruction to be cancelled.
	AgentCADeactivationInstructionIdentification *iso20022.DocumentIdentification8 `xml:"AgtCADeactvtnInstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Details of the deactivation instruction to be cancelled.
	DeactivationInstructionDetails *iso20022.CorporateActionDeactivationInstruction1 `xml:"DeactvtnInstrDtls,omitempty"`
}

func (a *AgentCADeactivationCancellationRequestV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCADeactivationCancellationRequestV01) AddAgentCADeactivationInstructionIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCADeactivationInstructionIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCADeactivationInstructionIdentification
}

func (a *AgentCADeactivationCancellationRequestV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCADeactivationCancellationRequestV01) AddDeactivationInstructionDetails() *iso20022.CorporateActionDeactivationInstruction1 {
	a.DeactivationInstructionDetails = new(iso20022.CorporateActionDeactivationInstruction1)
	return a.DeactivationInstructionDetails
}
