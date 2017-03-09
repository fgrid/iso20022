package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01800101 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.018.001.01 Document"`
	Message *AgentCAGlobalDistributionStatusAdviceV01 `xml:"AgtCAGblDstrbtnStsAdvc"`
}

func (d *Document01800101) AddMessage() *AgentCAGlobalDistributionStatusAdviceV01 {
	d.Message = new(AgentCAGlobalDistributionStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to the CSD to authorise/prohibit the CSD to process the entitlement movements.
// Usage
// This message is used to authorise/prohibit the CSD to process the movements requested in the Global Distribution Authorisation Request message.
// Once the amendment request has been accepted by the issuer (or its agent), the CSD will process any resource movement and send an Agent Corporate Action Election Advice message with the function, option change, to confirm that the amendment has been booked at the CSD.
// The issuer (or its agent) can provide the status in 2 different ways:
// - Provide a global status, in which case the building block Global Movement Status must be present; or
// - Provide a status by individual movements, in which case, the building block Individual Movement Status must be present. An individual movement cannot be rejected.
type AgentCAGlobalDistributionStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Global Distribution Authorisation Request for which a status is given.
	AgentCAGlobalDistributionAuthorisationRequestIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAGblDstrbtnAuthstnReqId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the status of the global movement.
	GlobalMovementStatus *iso20022.GlobalDistributionStatus1 `xml:"GblMvmntSts"`

	// Provides information about the status of an individual movement.
	IndividualMovementStatus []*iso20022.IndividualMovementStatus1 `xml:"IndvMvmntSts"`
}

func (a *AgentCAGlobalDistributionStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAGlobalDistributionStatusAdviceV01) AddAgentCAGlobalDistributionAuthorisationRequestIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAGlobalDistributionAuthorisationRequestIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAGlobalDistributionAuthorisationRequestIdentification
}

func (a *AgentCAGlobalDistributionStatusAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAGlobalDistributionStatusAdviceV01) AddGlobalMovementStatus() *iso20022.GlobalDistributionStatus1 {
	a.GlobalMovementStatus = new(iso20022.GlobalDistributionStatus1)
	return a.GlobalMovementStatus
}

func (a *AgentCAGlobalDistributionStatusAdviceV01) AddIndividualMovementStatus() *iso20022.IndividualMovementStatus1 {
	newValue := new(iso20022.IndividualMovementStatus1)
	a.IndividualMovementStatus = append(a.IndividualMovementStatus, newValue)
	return newValue
}
