package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.031.001.01 Document"`
	Message *CorporateActionNotificationV01 `xml:"CorpActnNtfctn"`
}

func (d *Document03100101) AddMessage() *CorporateActionNotificationV01 {
	d.Message = new(CorporateActionNotificationV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionNotification message to an account owner or its designated agent to notify details of a corporate action event and optionally account information, eligible balance and entitlements. It may also include possible elections or choices available to the account owner.
// The account servicer can initially send the CorporateActionNotification message as a preliminary advice, subsequently replaced by another CorporateActionNotification message with complete or confirmed information.
// It may also be sent to an account owner or its designated agent, to remind of event details and/or of missing or incomplete instructions for a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionNotificationV01 struct {

	// Information that unambiguously identifies a CorporateActionNotification message as know by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// General information about the event notification type, status and contents.
	NotificationGeneralInformation *iso20022.CorporateActionNotification2 `xml:"NtfctnGnlInf"`

	// Identification of a previously sent notification document.
	PreviousNotificationIdentification *iso20022.DocumentIdentification15 `xml:"PrvsNtfctnId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation11 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification7Choice `xml:"AcctDtls"`

	// Provides details on rights credited to the account as for instance trading period, expiry date, renounceability.
	IntermediateSecurity *iso20022.FinancialInstrumentAttributes3 `xml:"IntrmdtScty,omitempty"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction3 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionOptionDetails []*iso20022.CorporateActionOption3 `xml:"CorpActnOptnDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative11 `xml:"AddtlInf,omitempty"`

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


func (c *CorporateActionNotificationV01) AddIdentification() *iso20022.DocumentIdentification11 {
	c.Identification = new(iso20022.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionNotificationV01) AddNotificationGeneralInformation() *iso20022.CorporateActionNotification2 {
	c.NotificationGeneralInformation = new(iso20022.CorporateActionNotification2)
	return c.NotificationGeneralInformation
}

func (c *CorporateActionNotificationV01) AddPreviousNotificationIdentification() *iso20022.DocumentIdentification15 {
	c.PreviousNotificationIdentification = new(iso20022.DocumentIdentification15)
	return c.PreviousNotificationIdentification
}

func (c *CorporateActionNotificationV01) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionNotificationV01) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new (iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new (iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation11 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation11)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNotificationV01) AddAccountDetails() *iso20022.AccountIdentification7Choice {
	c.AccountDetails = new(iso20022.AccountIdentification7Choice)
	return c.AccountDetails
}

func (c *CorporateActionNotificationV01) AddIntermediateSecurity() *iso20022.FinancialInstrumentAttributes3 {
	c.IntermediateSecurity = new(iso20022.FinancialInstrumentAttributes3)
	return c.IntermediateSecurity
}

func (c *CorporateActionNotificationV01) AddCorporateActionDetails() *iso20022.CorporateAction3 {
	c.CorporateActionDetails = new(iso20022.CorporateAction3)
	return c.CorporateActionDetails
}

func (c *CorporateActionNotificationV01) AddCorporateActionOptionDetails() *iso20022.CorporateActionOption3 {
	newValue := new (iso20022.CorporateActionOption3)
	c.CorporateActionOptionDetails = append(c.CorporateActionOptionDetails, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddAdditionalInformation() *iso20022.CorporateActionNarrative11 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative11)
	return c.AdditionalInformation
}

func (c *CorporateActionNotificationV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	c.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionNotificationV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	c.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionNotificationV01) AddIssuerAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddPayingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddSubPayingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddRegistrar() *iso20022.PartyIdentification10Choice {
	c.Registrar = new(iso20022.PartyIdentification10Choice)
	return c.Registrar
}

func (c *CorporateActionNotificationV01) AddResellingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification10Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification10Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionNotificationV01) AddDropAgent() *iso20022.PartyIdentification10Choice {
	c.DropAgent = new(iso20022.PartyIdentification10Choice)
	return c.DropAgent
}

func (c *CorporateActionNotificationV01) AddSolicitationAgent() *iso20022.PartyIdentification10Choice {
	newValue := new (iso20022.PartyIdentification10Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV01) AddInformationAgent() *iso20022.PartyIdentification10Choice {
	c.InformationAgent = new(iso20022.PartyIdentification10Choice)
	return c.InformationAgent
}

func (c *CorporateActionNotificationV01) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}

