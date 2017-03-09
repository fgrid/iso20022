package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500208 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.002.08 Document"`
	Message *CorporateActionMovementPreliminaryAdvice002V08 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500208) AddMessage() *CorporateActionMovementPreliminaryAdvice002V08 {
	d.Message = new(CorporateActionMovementPreliminaryAdvice002V08)
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
type CorporateActionMovementPreliminaryAdvice002V08 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts preliminary advice is to continue or that the message is the last page of the multi-parts preliminary advice.
	Pagination *iso20022.Pagination `xml:"Pgntn,omitempty"`

	// General information about the movement preliminary advice document.
	MovementPreliminaryAdviceGeneralInformation *iso20022.CorporateActionPreliminaryAdviceType2 `xml:"MvmntPrlimryAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	PreviousMovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification37 `xml:"PrvsMvmntPrlimryAdvcId,omitempty"`

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification37 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *iso20022.DocumentIdentification37 `xml:"MvmntConfId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification17 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *iso20022.CorporateActionReversalReason4 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation117 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification36Choice `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction38 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*iso20022.CorporateActionOption135 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative37 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification104Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification104Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification104Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *iso20022.PartyIdentification104Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*iso20022.PartyIdentification104Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *iso20022.PartyIdentification104Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *iso20022.PartyIdentification104Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*iso20022.PartyIdentification104Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *iso20022.PartyIdentification104Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddMovementPreliminaryAdviceGeneralInformation() *iso20022.CorporateActionPreliminaryAdviceType2 {
	c.MovementPreliminaryAdviceGeneralInformation = new(iso20022.CorporateActionPreliminaryAdviceType2)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPreviousMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification37 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification37)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddNotificationIdentification() *iso20022.DocumentIdentification37 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification37)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification37 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification37)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddInstructionIdentification() *iso20022.DocumentIdentification17 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification17)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddOtherDocumentIdentification() *iso20022.DocumentIdentification38 {
	newValue := new(iso20022.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddEventsLinkage() *iso20022.CorporateActionEventReference4 {
	newValue := new(iso20022.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddReversalReason() *iso20022.CorporateActionReversalReason4 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason4)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation117 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation117)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddAccountDetails() *iso20022.AccountIdentification36Choice {
	c.AccountDetails = new(iso20022.AccountIdentification36Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddCorporateActionDetails() *iso20022.CorporateAction38 {
	c.CorporateActionDetails = new(iso20022.CorporateAction38)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddCorporateActionMovementDetails() *iso20022.CorporateActionOption135 {
	newValue := new(iso20022.CorporateActionOption135)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddAdditionalInformation() *iso20022.CorporateActionNarrative37 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative37)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddIssuerAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddSubPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddRegistrar() *iso20022.PartyIdentification104Choice {
	c.Registrar = new(iso20022.PartyIdentification104Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddResellingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification104Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification104Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddDropAgent() *iso20022.PartyIdentification104Choice {
	c.DropAgent = new(iso20022.PartyIdentification104Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddSolicitationAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddInformationAgent() *iso20022.PartyIdentification104Choice {
	c.InformationAgent = new(iso20022.PartyIdentification104Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
