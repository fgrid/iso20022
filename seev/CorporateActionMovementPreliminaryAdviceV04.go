package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.001.04 Document"`
	Message *CorporateActionMovementPreliminaryAdviceV04 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500104) AddMessage() *CorporateActionMovementPreliminaryAdviceV04 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdvice message to an account owner or its designated agent to pre-advise upcoming posting or reversal of securities and/or cash postings.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementPreliminaryAdviceV04 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts preliminary advice is to continue or that the message is the last page of the multi-parts preliminary advice.
	Pagination *iso20022.Pagination `xml:"Pgntn,omitempty"`

	// General information about the movement preliminary advice document.
	MovementPreliminaryAdviceGeneralInformation *iso20022.CorporateActionPreliminaryAdviceType1 `xml:"MvmntPrlimryAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	PreviousMovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification15 `xml:"PrvsMvmntPrlimryAdvcId,omitempty"`

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification15 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *iso20022.DocumentIdentification15 `xml:"MvmntConfId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *iso20022.CorporateActionReversalReason1 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation54 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification21Choice `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*iso20022.CorporateActionOption52 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative6 `xml:"AddtlInf,omitempty"`

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


func (c *CorporateActionMovementPreliminaryAdviceV04) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddMovementPreliminaryAdviceGeneralInformation() *iso20022.CorporateActionPreliminaryAdviceType1 {
	c.MovementPreliminaryAdviceGeneralInformation = new(iso20022.CorporateActionPreliminaryAdviceType1)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddPreviousMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification15 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification15)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddNotificationIdentification() *iso20022.DocumentIdentification15 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification15)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification15 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification15)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new (iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new (iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddReversalReason() *iso20022.CorporateActionReversalReason1 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason1)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation54 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation54)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddAccountDetails() *iso20022.AccountIdentification21Choice {
	c.AccountDetails = new(iso20022.AccountIdentification21Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddCorporateActionMovementDetails() *iso20022.CorporateActionOption52 {
	newValue := new (iso20022.CorporateActionOption52)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddAdditionalInformation() *iso20022.CorporateActionNarrative6 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative6)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddIssuerAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddSubPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddRegistrar() *iso20022.PartyIdentification46Choice {
	c.Registrar = new(iso20022.PartyIdentification46Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddResellingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification46Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification46Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddDropAgent() *iso20022.PartyIdentification46Choice {
	c.DropAgent = new(iso20022.PartyIdentification46Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddSolicitationAgent() *iso20022.PartyIdentification46Choice {
	newValue := new (iso20022.PartyIdentification46Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddInformationAgent() *iso20022.PartyIdentification46Choice {
	c.InformationAgent = new(iso20022.PartyIdentification46Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

