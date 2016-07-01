package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02600101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.026.001.01 Document"`
	Message *AgentCAStandingInstructionCancellationRequestV01 `xml:"AgtCAStgInstrCxlReq"`
}

func (d *Document02600101) AddMessage() *AgentCAStandingInstructionCancellationRequestV01 {
	d.Message = new(AgentCAStandingInstructionCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer agent to request the cancellation of a previously sent Agent Corporate Action Standing Instruction.
// Usage
// This message is used to request the cancellation of a standing instruction.
// This message must contain the identification of the standing instruction to be cancelled. It may also contain details of the standing instruction to be cancelled, but this is not recommended.
type AgentCAStandingInstructionCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Standing Instruction Request to be cancelled.
	AgentCAStandingInstructionRequestIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAStgInstrReqId"`

	// General information about the standing instruction.
	StandingInstructionGeneralInformation *iso20022.CorporateActionStandingInstructionGeneralInformation1 `xml:"StgInstrGnlInf"`

	// Information related to the standing instruction.
	StandingInstructionDetails *iso20022.CorporateActionStandingInstruction1 `xml:"StgInstrDtls,omitempty"`

}


func (a *AgentCAStandingInstructionCancellationRequestV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAStandingInstructionCancellationRequestV01) AddAgentCAStandingInstructionRequestIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAStandingInstructionRequestIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAStandingInstructionRequestIdentification
}

func (a *AgentCAStandingInstructionCancellationRequestV01) AddStandingInstructionGeneralInformation() *iso20022.CorporateActionStandingInstructionGeneralInformation1 {
	a.StandingInstructionGeneralInformation = new(iso20022.CorporateActionStandingInstructionGeneralInformation1)
	return a.StandingInstructionGeneralInformation
}

func (a *AgentCAStandingInstructionCancellationRequestV01) AddStandingInstructionDetails() *iso20022.CorporateActionStandingInstruction1 {
	a.StandingInstructionDetails = new(iso20022.CorporateActionStandingInstruction1)
	return a.StandingInstructionDetails
}

