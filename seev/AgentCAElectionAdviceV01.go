package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.012.001.01 Document"`
	Message *AgentCAElectionAdviceV01 `xml:"AgtCAElctnAdvc"`
}

func (d *Document01200101) AddMessage() *AgentCAElectionAdviceV01 {
	d.Message = new(AgentCAElectionAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer (or its agent) to provide information about the clients' election instruction, the registration details, the delivery details, etc. In case of an election with underlying resource movements, it also confirms that these have been completed. This message may also be sent in case of an amendment of an election, once the CSD has completed the required resource movements.
// Usage
// This message can be used for a new election advice or an amended election advice.
// If this message is used for a new election advice, the function of the message must be 'new election'.
// If this message is used for an amended election advice, the function of the message must be 'option change' and the identification of the previously sent election advice must be present.
// This message can include the cash movements and/or securities movements in the case of an election with underlying resource movements. Additionally, this message can include delivery, certification and beneficial owner details.
// Note: this information can be also sent separately in the Agent Corporate Action Information advice message.
type AgentCAElectionAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Provides information about the type of election advice and linked messages.
	ElectionAdviceTypeAndLinkage *iso20022.ElectionAdviceFunction1 `xml:"ElctnAdvcTpAndLkg"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the election(s).
	ElectionDetails *iso20022.CorporateActionElection3 `xml:"ElctnDtls"`

	// Provides additional information about the delivery details, beneficial owner details, etc.
	AdditionalInformation *iso20022.CorporateActionAdditionalInformation1 `xml:"AddtlInf,omitempty"`

	// Contact responsible for the transaction identified in the message.
	ContactDetails *iso20022.ContactPerson1 `xml:"CtctDtls,omitempty"`
}

func (a *AgentCAElectionAdviceV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAElectionAdviceV01) AddElectionAdviceTypeAndLinkage() *iso20022.ElectionAdviceFunction1 {
	a.ElectionAdviceTypeAndLinkage = new(iso20022.ElectionAdviceFunction1)
	return a.ElectionAdviceTypeAndLinkage
}

func (a *AgentCAElectionAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAElectionAdviceV01) AddElectionDetails() *iso20022.CorporateActionElection3 {
	a.ElectionDetails = new(iso20022.CorporateActionElection3)
	return a.ElectionDetails
}

func (a *AgentCAElectionAdviceV01) AddAdditionalInformation() *iso20022.CorporateActionAdditionalInformation1 {
	a.AdditionalInformation = new(iso20022.CorporateActionAdditionalInformation1)
	return a.AdditionalInformation
}

func (a *AgentCAElectionAdviceV01) AddContactDetails() *iso20022.ContactPerson1 {
	a.ContactDetails = new(iso20022.ContactPerson1)
	return a.ContactDetails
}
