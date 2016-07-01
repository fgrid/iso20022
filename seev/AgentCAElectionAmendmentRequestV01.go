package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.013.001.01 Document"`
	Message *AgentCAElectionAmendmentRequestV01 `xml:"AgtCAElctnAmdmntReq"`
}

func (d *Document01300101) AddMessage() *AgentCAElectionAmendmentRequestV01 {
	d.Message = new(AgentCAElectionAmendmentRequestV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer (or its agent) to request the authorisation of an amendment of a previously sent Agent Corporate Action Election Advice message.
// Usage
// This message is used to request the amendment of a previously sent Agent Corporate Action Election Advice message.
// Once the amendment request has been accepted by the issuer (or its agent), the CSD will process any resource movement and send an Agent Corporate Action Election Advice message with the function, option change, to confirm that the amendment has been booked at the CSD.
// This message is used when the terms and conditions of the corporate action event allow amendments.
type AgentCAElectionAmendmentRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the request.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Election Advice for which an amendment is requested.
	AgentCAElectionAdviceIdentification *iso20022.DocumentIdentification8 `xml:"AgtCAElctnAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the account.
	AccountDetails *iso20022.SecuritiesAccount7 `xml:"AcctDtls"`

	// Provides information about the original election to be amended.
	OriginalElectionDetails *iso20022.CorporateActionElection1 `xml:"OrgnlElctnDtls"`

	// Provides information about the amendments to the election.
	AmendedElectionDetails *iso20022.CorporateActionElection2 `xml:"AmddElctnDtls"`

	// Contact responsible for the transaction identified in the message.
	ContactDetails *iso20022.ContactPerson1 `xml:"CtctDtls,omitempty"`

}


func (a *AgentCAElectionAmendmentRequestV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAElectionAmendmentRequestV01) AddAgentCAElectionAdviceIdentification() *iso20022.DocumentIdentification8 {
	a.AgentCAElectionAdviceIdentification = new(iso20022.DocumentIdentification8)
	return a.AgentCAElectionAdviceIdentification
}

func (a *AgentCAElectionAmendmentRequestV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAElectionAmendmentRequestV01) AddAccountDetails() *iso20022.SecuritiesAccount7 {
	a.AccountDetails = new(iso20022.SecuritiesAccount7)
	return a.AccountDetails
}

func (a *AgentCAElectionAmendmentRequestV01) AddOriginalElectionDetails() *iso20022.CorporateActionElection1 {
	a.OriginalElectionDetails = new(iso20022.CorporateActionElection1)
	return a.OriginalElectionDetails
}

func (a *AgentCAElectionAmendmentRequestV01) AddAmendedElectionDetails() *iso20022.CorporateActionElection2 {
	a.AmendedElectionDetails = new(iso20022.CorporateActionElection2)
	return a.AmendedElectionDetails
}

func (a *AgentCAElectionAmendmentRequestV01) AddContactDetails() *iso20022.ContactPerson1 {
	a.ContactDetails = new(iso20022.ContactPerson1)
	return a.ContactDetails
}

