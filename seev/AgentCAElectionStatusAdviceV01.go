package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.015.001.01 Document"`
	Message *AgentCAElectionStatusAdviceV01 `xml:"AgtCAElctnStsAdvc"`
}

func (d *Document01500101) AddMessage() *AgentCAElectionStatusAdviceV01 {
	d.Message = new(AgentCAElectionStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to the CSD to report the status, or a change in status, of:
// - a corporate action election advice;
// - an election cancellation request; or
// - an election amendment request.
// Usage
// This message must be sent in response to an:
// - Agent Corporation Action Election Advice to provide the status of an election advice in the case of a rejection. However, it may also be used in all other situations, in which case, the building blocks Election Advice Identification and the Election Advice Status must be present.
// - Agent Corporation Action Election Cancellation Request to provide the status of the cancellation request, in which case, the building blocks Election Cancellation Request Identification and the Election Cancellation Request Status must be present.
// - Agent Corporation Action Election Amendment Request to provide the status of the amendment request, in which case, the building blocks Election Amendment Request Identification and the Election Amendment Request Status must be present.
type AgentCAElectionStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Election Advice for which a status is given.
	AgentCAElectionAdviceIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAElctnAdvcId"`

	// Identification of the linked Agent CA Election Cancellation Request for which a status is given.
	AgentCAElectionCancellationRequestIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAElctnCxlReqId"`

	// Identification of the linked Agent CA Election Amendment Request for which a status is given.
	AgentCAElectionAmendmentRequestIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAElctnAmdmntReqId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Status of the election advice sent by the CSD.
	ElectionAdviceStatus *iso20022.ElectionAdviceStatus1Choice `xml:"ElctnAdvcSts"`

	// Status of the election cancellation request sent by the CSD.
	ElectionCancellationRequestStatus *iso20022.ElectionCancellationStatus1Choice `xml:"ElctnCxlReqSts"`

	// Status of the amendment request sent by the CSD.
	ElectionAmendmentRequestStatus *iso20022.ElectionAmendmentStatus1Choice `xml:"ElctnAmdmntReqSts"`

}


func (a *AgentCAElectionStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAElectionStatusAdviceV01) AddAgentCAElectionAdviceIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAElectionAdviceIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAElectionAdviceIdentification
}

func (a *AgentCAElectionStatusAdviceV01) AddAgentCAElectionCancellationRequestIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAElectionCancellationRequestIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAElectionCancellationRequestIdentification
}

func (a *AgentCAElectionStatusAdviceV01) AddAgentCAElectionAmendmentRequestIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAElectionAmendmentRequestIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAElectionAmendmentRequestIdentification
}

func (a *AgentCAElectionStatusAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAElectionStatusAdviceV01) AddElectionAdviceStatus() *iso20022.ElectionAdviceStatus1Choice {
	a.ElectionAdviceStatus = new(iso20022.ElectionAdviceStatus1Choice)
	return a.ElectionAdviceStatus
}

func (a *AgentCAElectionStatusAdviceV01) AddElectionCancellationRequestStatus() *iso20022.ElectionCancellationStatus1Choice {
	a.ElectionCancellationRequestStatus = new(iso20022.ElectionCancellationStatus1Choice)
	return a.ElectionCancellationRequestStatus
}

func (a *AgentCAElectionStatusAdviceV01) AddElectionAmendmentRequestStatus() *iso20022.ElectionAmendmentStatus1Choice {
	a.ElectionAmendmentRequestStatus = new(iso20022.ElectionAmendmentStatus1Choice)
	return a.ElectionAmendmentRequestStatus
}

