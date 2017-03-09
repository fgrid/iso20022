package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02200101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.022.001.01 Document"`
	Message *AgentCAMovementStatusAdviceV01 `xml:"AgtCAMvmntStsAdvc"`
}

func (d *Document02200101) AddMessage() *AgentCAMovementStatusAdviceV01 {
	d.Message = new(AgentCAMovementStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to report the status, or a change in status, of
// - a global distribution status advice released by an issuer (or its agent);
// - a movement instruction released by an issuer (or its agent);
// - a movement cancellation request sent by the issuer (or its agent); and
// - the non-settlement of the movements at the CSD.
// Usage
// This message is used to report the status of:
// - the movements resulting from a movement instruction message, in which case, the Agent Corporate Action Movement Instruction Identification must be present;
// - the movements resulting from a global distribution status advice message (with the status, authorised), in which case, the Agent Corporate Action Global Distribution Status Advice Identification must be present;
// - the movement cancellation request, in which case, the Agent Corporate Action Movement Cancellation Request Identification must be present; and
// - the movements resulting from an election status advice (if the status of the election advice is rejected or if the status of the election cancellation request or amendment request is accepted) in case there is a settlement problem. The Election Status Advice Identification must be present.
// In the case of a failed settlement, the message contains details of the movement, such as account details, securities or cash information and the reason of the failure.
// This message should not be used to provide the confirmation of the settlement; the Agent Corporate Action Movement Confirmation message should be used instead.
type AgentCAMovementStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent Corporate Action Election Status Advice.
	AgentCAElectionStatusAdviceIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAElctnStsAdvcId"`

	// Identification of the Agent Corporate Action Global Distribution Status Advice.
	AgentCAGlobalDistributionStatusAdviceIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAGblDstrbtnStsAdvcId"`

	// Identification of the linked Agent CA Movement Instruction for which a status is given.
	AgentCAMovementInstructionIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAMvmntInstrId"`

	// Identification of the linked Agent CA Movement Cancellation Request for which a status is given.
	AgentCAMovementCancellationRequestIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAMvmntCxlReqId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Status of the movement instruction.
	MovementStatusDetails *iso20022.CorporateActionMovementStatus1Choice `xml:"MvmntStsDtls"`

	// Status of the movement cancellation request.
	MovementCancellationStatusDetails *iso20022.CorporateMovementStatus2 `xml:"MvmntCxlStsDtls"`
}

func (a *AgentCAMovementStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAMovementStatusAdviceV01) AddAgentCAElectionStatusAdviceIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAElectionStatusAdviceIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAElectionStatusAdviceIdentification
}

func (a *AgentCAMovementStatusAdviceV01) AddAgentCAGlobalDistributionStatusAdviceIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAGlobalDistributionStatusAdviceIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAGlobalDistributionStatusAdviceIdentification
}

func (a *AgentCAMovementStatusAdviceV01) AddAgentCAMovementInstructionIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAMovementInstructionIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAMovementInstructionIdentification
}

func (a *AgentCAMovementStatusAdviceV01) AddAgentCAMovementCancellationRequestIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAMovementCancellationRequestIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAMovementCancellationRequestIdentification
}

func (a *AgentCAMovementStatusAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAMovementStatusAdviceV01) AddMovementStatusDetails() *iso20022.CorporateActionMovementStatus1Choice {
	a.MovementStatusDetails = new(iso20022.CorporateActionMovementStatus1Choice)
	return a.MovementStatusDetails
}

func (a *AgentCAMovementStatusAdviceV01) AddMovementCancellationStatusDetails() *iso20022.CorporateMovementStatus2 {
	a.MovementCancellationStatusDetails = new(iso20022.CorporateMovementStatus2)
	return a.MovementCancellationStatusDetails
}
