package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03700208 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.002.08 Document"`
	Message *CorporateActionMovementReversalAdvice002V08 `xml:"CorpActnMvmntRvslAdvc"`
}

func (d *Document03700208) AddMessage() *CorporateActionMovementReversalAdvice002V08 {
	d.Message = new(CorporateActionMovementReversalAdvice002V08)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementReversalAdvice message to an account owner or its designated agent to reverse previously confirmed posting of securities or cash.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionMovementReversalAdvice002V08 struct {

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *iso20022.DocumentIdentification37 `xml:"MvmntConfId"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *iso20022.CorporateActionReversalReason4 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation118 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountAndBalance40 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction36 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *iso20022.CorporateActionOption126 `xml:"CorpActnConfDtls"`

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

func (c *CorporateActionMovementReversalAdvice002V08) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification37 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification37)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementReversalAdvice002V08) AddOtherDocumentIdentification() *iso20022.DocumentIdentification38 {
	newValue := new(iso20022.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V08) AddEventsLinkage() *iso20022.CorporateActionEventReference4 {
	newValue := new(iso20022.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V08) AddReversalReason() *iso20022.CorporateActionReversalReason4 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason4)
	return c.ReversalReason
}

func (c *CorporateActionMovementReversalAdvice002V08) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation118 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation118)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementReversalAdvice002V08) AddAccountDetails() *iso20022.AccountAndBalance40 {
	c.AccountDetails = new(iso20022.AccountAndBalance40)
	return c.AccountDetails
}

func (c *CorporateActionMovementReversalAdvice002V08) AddCorporateActionDetails() *iso20022.CorporateAction36 {
	c.CorporateActionDetails = new(iso20022.CorporateAction36)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementReversalAdvice002V08) AddCorporateActionConfirmationDetails() *iso20022.CorporateActionOption126 {
	c.CorporateActionConfirmationDetails = new(iso20022.CorporateActionOption126)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementReversalAdvice002V08) AddAdditionalInformation() *iso20022.CorporateActionNarrative35 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative35)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementReversalAdvice002V08) AddIssuerAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V08) AddPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V08) AddSubPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
