package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02000101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.020.001.01 Document"`
	Message *AgentCAMovementCancellationRequestV01 `xml:"AgtCAMvmntCxlReq"`
}

func (d *Document02000101) AddMessage() *AgentCAMovementCancellationRequestV01 {
	d.Message = new(AgentCAMovementCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to request the cancellation of (a) movement(s) previously sent via an Agent Corporate Action Movement Instruction.
// Usage
// This message may be used to cancel an entire Agent Corporate Action Movement Instruction message that was previously sent by the issuer (or its agent) or specific movements.
// This message must contain the identification of the Agent Corporate Action Movement Instruction containing the movement(s) to be cancelled, the agent identification and the corporate action references. This message must also contain details of the movement(s) to be cancelled.
type AgentCAMovementCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Movement Instruction to be cancelled.
	AgentCAMovementInstructionIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAMvmntInstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Details of the movement instructions to be cancelled.
	MovementDetails *iso20022.MovementInstruction1 `xml:"MvmntDtls,omitempty"`

}


func (a *AgentCAMovementCancellationRequestV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAMovementCancellationRequestV01) AddAgentCAMovementInstructionIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAMovementInstructionIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAMovementInstructionIdentification
}

func (a *AgentCAMovementCancellationRequestV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAMovementCancellationRequestV01) AddMovementDetails() *iso20022.MovementInstruction1 {
	a.MovementDetails = new(iso20022.MovementInstruction1)
	return a.MovementDetails
}

