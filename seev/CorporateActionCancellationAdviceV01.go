package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03900101 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.01 Document"`
	Message *CorporateActionCancellationAdviceV01 `xml:"CorpActnCxlAdvc"`
}

func (d *Document03900101) AddMessage() *CorporateActionCancellationAdviceV01 {
	d.Message = new(CorporateActionCancellationAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionCancellationAdvice message to an account owner or its designated agent to cancel a previously announced corporate action event in case of error from the account servicer or in case of withdrawal by the issuer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionCancellationAdviceV01 struct {

	// Information that unambiguously identifies a CorporateActionCancellationAdvice message as know by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// General information about the event cancellation status and cancellation reason.
	CancellationAdviceGeneralInformation *iso20022.CorporateActionCancellation1 `xml:"CxlAdvcGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation8 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountsDetails *iso20022.AccountIdentification6Choice `xml:"AcctsDtls"`

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

func (c *CorporateActionCancellationAdviceV01) AddIdentification() *iso20022.DocumentIdentification11 {
	c.Identification = new(iso20022.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionCancellationAdviceV01) AddCancellationAdviceGeneralInformation() *iso20022.CorporateActionCancellation1 {
	c.CancellationAdviceGeneralInformation = new(iso20022.CorporateActionCancellation1)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionCancellationAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation8 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation8)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionCancellationAdviceV01) AddAccountsDetails() *iso20022.AccountIdentification6Choice {
	c.AccountsDetails = new(iso20022.AccountIdentification6Choice)
	return c.AccountsDetails
}

func (c *CorporateActionCancellationAdviceV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	c.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionCancellationAdviceV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	c.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionCancellationAdviceV01) AddIssuerAgent() *iso20022.PartyIdentification10Choice {
	newValue := new(iso20022.PartyIdentification10Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV01) AddPayingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new(iso20022.PartyIdentification10Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV01) AddSubPayingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new(iso20022.PartyIdentification10Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV01) AddRegistrar() *iso20022.PartyIdentification10Choice {
	c.Registrar = new(iso20022.PartyIdentification10Choice)
	return c.Registrar
}

func (c *CorporateActionCancellationAdviceV01) AddResellingAgent() *iso20022.PartyIdentification10Choice {
	newValue := new(iso20022.PartyIdentification10Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV01) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification10Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification10Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionCancellationAdviceV01) AddDropAgent() *iso20022.PartyIdentification10Choice {
	c.DropAgent = new(iso20022.PartyIdentification10Choice)
	return c.DropAgent
}

func (c *CorporateActionCancellationAdviceV01) AddSolicitationAgent() *iso20022.PartyIdentification10Choice {
	newValue := new(iso20022.PartyIdentification10Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV01) AddInformationAgent() *iso20022.PartyIdentification10Choice {
	c.InformationAgent = new(iso20022.PartyIdentification10Choice)
	return c.InformationAgent
}

func (c *CorporateActionCancellationAdviceV01) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
