package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.001.01 Document"`
	Message *CorporateActionMovementPreliminaryAdviceV01 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500101) AddMessage() *CorporateActionMovementPreliminaryAdviceV01 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdvice message to an account owner or its designated agent to pre-advise upcoming posting or reversal of securities and/or cash postings.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementPreliminaryAdviceV01 struct {

	// Information that unambiguously identifies a CorporateActionMovementPreliminaryAdvice message as know by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

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
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation3 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification7Choice `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*iso20022.CorporateActionOption10 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative6 `xml:"AddtlInf,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification10Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification10Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification10Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *iso20022.PartyIdentification10Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner. 
	ResellingAgent []*iso20022.PartyIdentification10Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *iso20022.PartyIdentification10Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person. 
	DropAgent *iso20022.PartyIdentification10Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change. 
	SolicitationAgent []*iso20022.PartyIdentification10Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation. 
	InformationAgent *iso20022.PartyIdentification10Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`

}


func (c *CorporateActionMovementPreliminaryAdviceV01) AddIdentification() *iso20022.DocumentIdentification11 {
	c.Identification = new(iso20022.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddMovementPreliminaryAdviceGeneralInformation() *iso20022.CorporateActionPreliminaryAdviceType1 {
	c.MovementPreliminaryAdviceGeneralInformation = new(iso20022.CorporateActionPreliminaryAdviceType1)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddPreviousMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification15 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification15)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddNotificationIdentification() *iso20022.DocumentIdentification15 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification15)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification15 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification15)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new (iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new (iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddReversalReason() *iso20022.CorporateActionReversalReason1 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason1)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation3 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation3)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddAccountDetails() *iso20022.AccountIdentification7Choice {
	c.AccountDetails = new(iso20022.AccountIdentification7Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddCorporateActionMovementDetails() *iso20022.CorporateActionOption10 {
	newValue := new (iso20022.CorporateActionOption10)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddAdditionalInformation() *iso20022.CorporateActionNarrative6 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative6)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	c.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	c.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddIssuerAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddPayingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddSubPayingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddRegistrar() *iso20022.PartyIdentification10Choice {
	c.Registrar = new(iso20022.PartyIdentification10Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddResellingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification10Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification10Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddDropAgent() *iso20022.PartyIdentification10Choice {
	c.DropAgent = new(iso20022.PartyIdentification10Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddSolicitationAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddInformationAgent() *iso20022.PartyIdentification10Choice {
	c.InformationAgent = new(iso20022.PartyIdentification10Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV01) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}

