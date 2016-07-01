package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.024.001.01 Document"`
	Message *AgentCAInformationStatusAdviceV01 `xml:"AgtCAInfStsAdvc"`
}

func (d *Document02400101) AddMessage() *AgentCAInformationStatusAdviceV01 {
	d.Message = new(AgentCAInformationStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to report the status, or change in status, of an information advice.
// Usage
// This message must be used in response to an Agent Corporate Action Information Advice in the case of a rejection. However, it may also be used to report other statuses.
// The information advice identification must be present to link this message to the information advice for which the status is provided.
type AgentCAInformationStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Information Advice for which a status is given.
	AgentCAInformationAdviceIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAInfAdvcId"`

	// Additional information about the corporate action such as the delivery details.
	CorporateActionAdditionalInformation *iso20022.CorporateActionAdditionalInformation1 `xml:"CorpActnAddtlInf,omitempty"`

	// Status of the information advice sent by the CSD.
	InformationStatusDetails *iso20022.CorporateActionInformationStatus1Choice `xml:"InfStsDtls"`

}


func (a *AgentCAInformationStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAInformationStatusAdviceV01) AddAgentCAInformationAdviceIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAInformationAdviceIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAInformationAdviceIdentification
}

func (a *AgentCAInformationStatusAdviceV01) AddCorporateActionAdditionalInformation() *iso20022.CorporateActionAdditionalInformation1 {
	a.CorporateActionAdditionalInformation = new(iso20022.CorporateActionAdditionalInformation1)
	return a.CorporateActionAdditionalInformation
}

func (a *AgentCAInformationStatusAdviceV01) AddInformationStatusDetails() *iso20022.CorporateActionInformationStatus1Choice {
	a.InformationStatusDetails = new(iso20022.CorporateActionInformationStatus1Choice)
	return a.InformationStatusDetails
}

