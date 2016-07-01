package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03100206 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.031.002.06 Document"`
	Message *CorporateActionNotification002V06 `xml:"CorpActnNtfctn"`
}

func (d *Document03100206) AddMessage() *CorporateActionNotification002V06 {
	d.Message = new(CorporateActionNotification002V06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionNotification message to an account owner or its designated agent to notify details of a corporate action event and optionally account information, eligible balance and entitlements.
// It may also include possible elections or choices available to the account owner. The account servicer can initially send the CorporateActionNotification message as a preliminary advice, subsequently replaced by another CorporateActionNotification message with complete or confirmed information.
// It may also be sent to an account owner or its designated agent, to remind of event details and/or of missing or incomplete instructions for a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionNotification002V06 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts notification is to continue or that the message is the last page of the multi-parts notification.
	Pagination *iso20022.Pagination `xml:"Pgntn,omitempty"`

	// General information about the event notification type, status and contents.
	NotificationGeneralInformation *iso20022.CorporateActionNotification6 `xml:"NtfctnGnlInf"`

	// Identification of a previously sent notification document.
	PreviousNotificationIdentification *iso20022.DocumentIdentification37 `xml:"PrvsNtfctnId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification17 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation103 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification36Choice `xml:"AcctDtls"`

	// Provides details on rights credited to the account as for instance trading period, expiry date, renounceability.
	IntermediateSecurity *iso20022.FinancialInstrumentAttributes73 `xml:"IntrmdtScty,omitempty"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction40 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionOptionDetails []*iso20022.CorporateActionOption127 `xml:"CorpActnOptnDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative41 `xml:"AddtlInf,omitempty"`

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


func (c *CorporateActionNotification002V06) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CorporateActionNotification002V06) AddNotificationGeneralInformation() *iso20022.CorporateActionNotification6 {
	c.NotificationGeneralInformation = new(iso20022.CorporateActionNotification6)
	return c.NotificationGeneralInformation
}

func (c *CorporateActionNotification002V06) AddPreviousNotificationIdentification() *iso20022.DocumentIdentification37 {
	c.PreviousNotificationIdentification = new(iso20022.DocumentIdentification37)
	return c.PreviousNotificationIdentification
}

func (c *CorporateActionNotification002V06) AddInstructionIdentification() *iso20022.DocumentIdentification17 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification17)
	return c.InstructionIdentification
}

func (c *CorporateActionNotification002V06) AddOtherDocumentIdentification() *iso20022.DocumentIdentification38 {
	newValue := new (iso20022.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddEventsLinkage() *iso20022.CorporateActionEventReference4 {
	newValue := new (iso20022.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation103 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation103)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNotification002V06) AddAccountDetails() *iso20022.AccountIdentification36Choice {
	c.AccountDetails = new(iso20022.AccountIdentification36Choice)
	return c.AccountDetails
}

func (c *CorporateActionNotification002V06) AddIntermediateSecurity() *iso20022.FinancialInstrumentAttributes73 {
	c.IntermediateSecurity = new(iso20022.FinancialInstrumentAttributes73)
	return c.IntermediateSecurity
}

func (c *CorporateActionNotification002V06) AddCorporateActionDetails() *iso20022.CorporateAction40 {
	c.CorporateActionDetails = new(iso20022.CorporateAction40)
	return c.CorporateActionDetails
}

func (c *CorporateActionNotification002V06) AddCorporateActionOptionDetails() *iso20022.CorporateActionOption127 {
	newValue := new (iso20022.CorporateActionOption127)
	c.CorporateActionOptionDetails = append(c.CorporateActionOptionDetails, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddAdditionalInformation() *iso20022.CorporateActionNarrative41 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative41)
	return c.AdditionalInformation
}

func (c *CorporateActionNotification002V06) AddIssuerAgent() *iso20022.PartyIdentification104Choice {
	newValue := new (iso20022.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new (iso20022.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddSubPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new (iso20022.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddRegistrar() *iso20022.PartyIdentification104Choice {
	c.Registrar = new(iso20022.PartyIdentification104Choice)
	return c.Registrar
}

func (c *CorporateActionNotification002V06) AddResellingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new (iso20022.PartyIdentification104Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification104Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification104Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionNotification002V06) AddDropAgent() *iso20022.PartyIdentification104Choice {
	c.DropAgent = new(iso20022.PartyIdentification104Choice)
	return c.DropAgent
}

func (c *CorporateActionNotification002V06) AddSolicitationAgent() *iso20022.PartyIdentification104Choice {
	newValue := new (iso20022.PartyIdentification104Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionNotification002V06) AddInformationAgent() *iso20022.PartyIdentification104Choice {
	c.InformationAgent = new(iso20022.PartyIdentification104Choice)
	return c.InformationAgent
}

func (c *CorporateActionNotification002V06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

