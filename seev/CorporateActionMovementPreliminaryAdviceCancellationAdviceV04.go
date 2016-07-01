package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04400104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.04 Document"`
	Message *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04 `xml:"CorpActnMvmntPrlimryAdvcCxlAdvc"`
}

func (d *Document04400104) AddMessage() *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceCancellationAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdviceCancellationAdvice message to an account owner or its designated agent to cancel a previously announced CorporateActionMovementPreliminaryAdvice.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementPreliminaryAdviceCancellationAdviceV04 struct {

	// Specifies the status of the details of the event.
	CancellationAdviceGeneralInformation *iso20022.CorporateActionProcessingStatus1Choice `xml:"CxlAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification15 `xml:"MvmntPrlimryAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation53 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *iso20022.AccountIdentification13Choice `xml:"AcctDtls"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification46Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification46Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification46Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *iso20022.PartyIdentification46Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner. 
	ResellingAgent []*iso20022.PartyIdentification46Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *iso20022.PartyIdentification46Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person. 
	DropAgent *iso20022.PartyIdentification46Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*iso20022.PartyIdentification46Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation. 
	InformationAgent *iso20022.PartyIdentification46Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddCancellationAdviceGeneralInformation() *iso20022.CorporateActionProcessingStatus1Choice {
	c.CancellationAdviceGeneralInformation = new(iso20022.CorporateActionProcessingStatus1Choice)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification15 {
	c.MovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification15)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation53 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation53)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddAccountDetails() *iso20022.AccountIdentification13Choice {
	c.AccountDetails = new(iso20022.AccountIdentification13Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddIssuerAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddSubPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddRegistrar() *iso20022.PartyIdentification46Choice {
	c.Registrar = new(iso20022.PartyIdentification46Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddResellingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification46Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification46Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddDropAgent() *iso20022.PartyIdentification46Choice {
	c.DropAgent = new(iso20022.PartyIdentification46Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddSolicitationAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddInformationAgent() *iso20022.PartyIdentification46Choice {
	c.InformationAgent = new(iso20022.PartyIdentification46Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

