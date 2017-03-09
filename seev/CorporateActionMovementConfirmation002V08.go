package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03600208 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.036.002.08 Document"`
	Message *CorporateActionMovementConfirmation002V08 `xml:"CorpActnMvmntConf"`
}

func (d *Document03600208) AddMessage() *CorporateActionMovementConfirmation002V08 {
	d.Message = new(CorporateActionMovementConfirmation002V08)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementConfirmation message to an account owner or its designated agent to confirm posting of securities or cash as a result of a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionMovementConfirmation002V08 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification37 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification37 `xml:"MvmntPrlimryAdvcId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification17 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely  linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation118 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountAndBalance38 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction36 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *iso20022.CorporateActionOption136 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative35 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification104Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification104Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification104Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementConfirmation002V08) AddNotificationIdentification() *iso20022.DocumentIdentification37 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification37)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementConfirmation002V08) AddMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification37 {
	c.MovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification37)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementConfirmation002V08) AddInstructionIdentification() *iso20022.DocumentIdentification17 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification17)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementConfirmation002V08) AddOtherDocumentIdentification() *iso20022.DocumentIdentification38 {
	newValue := new(iso20022.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddEventsLinkage() *iso20022.CorporateActionEventReference4 {
	newValue := new(iso20022.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation118 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation118)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementConfirmation002V08) AddAccountDetails() *iso20022.AccountAndBalance38 {
	c.AccountDetails = new(iso20022.AccountAndBalance38)
	return c.AccountDetails
}

func (c *CorporateActionMovementConfirmation002V08) AddCorporateActionDetails() *iso20022.CorporateAction36 {
	c.CorporateActionDetails = new(iso20022.CorporateAction36)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementConfirmation002V08) AddCorporateActionConfirmationDetails() *iso20022.CorporateActionOption136 {
	c.CorporateActionConfirmationDetails = new(iso20022.CorporateActionOption136)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementConfirmation002V08) AddAdditionalInformation() *iso20022.CorporateActionNarrative35 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative35)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementConfirmation002V08) AddIssuerAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddSubPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
