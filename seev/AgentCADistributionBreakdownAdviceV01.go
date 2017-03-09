package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.016.001.01 Document"`
	Message *AgentCADistributionBreakdownAdviceV01 `xml:"AgtCADstrbtnBrkdwnAdvc"`
}

func (d *Document01600101) AddMessage() *AgentCADistributionBreakdownAdviceV01 {
	d.Message = new(AgentCADistributionBreakdownAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to provide distribution breakdown information for the proceeds that are to be delivered outside the CSD (e.g. when the proceeds are not eligible in the CSD).
// Usage
// This message is used to provide distribution breakdown information (securities and/or cash) per account for a specific corporate action option.
// Note: the delivery details are sent through the Agent Corporate Action Information Advice.
type AgentCADistributionBreakdownAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the CA option and the entitlements.
	CorporateActionDistributionDetails *iso20022.EntitlementAdvice1 `xml:"CorpActnDstrbtnDtls"`
}

func (a *AgentCADistributionBreakdownAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCADistributionBreakdownAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCADistributionBreakdownAdviceV01) AddCorporateActionDistributionDetails() *iso20022.EntitlementAdvice1 {
	a.CorporateActionDistributionDetails = new(iso20022.EntitlementAdvice1)
	return a.CorporateActionDistributionDetails
}
