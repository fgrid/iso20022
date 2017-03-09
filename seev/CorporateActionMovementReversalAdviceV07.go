package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03700107 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.07 Document"`
	Message *CorporateActionMovementReversalAdviceV07 `xml:"CorpActnMvmntRvslAdvc"`
}

func (d *Document03700107) AddMessage() *CorporateActionMovementReversalAdviceV07 {
	d.Message = new(CorporateActionMovementReversalAdviceV07)
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
type CorporateActionMovementReversalAdviceV07 struct {

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *iso20022.DocumentIdentification31 `xml:"MvmntConfId"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *iso20022.CorporateActionReversalReason3 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation89 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountAndBalance36 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction33 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *iso20022.CorporateActionOption119 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative31 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification71Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification71Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification71Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementReversalAdviceV07) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification31 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification31)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementReversalAdviceV07) AddOtherDocumentIdentification() *iso20022.DocumentIdentification32 {
	newValue := new(iso20022.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV07) AddEventsLinkage() *iso20022.CorporateActionEventReference3 {
	newValue := new(iso20022.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV07) AddReversalReason() *iso20022.CorporateActionReversalReason3 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason3)
	return c.ReversalReason
}

func (c *CorporateActionMovementReversalAdviceV07) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation89 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation89)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementReversalAdviceV07) AddAccountDetails() *iso20022.AccountAndBalance36 {
	c.AccountDetails = new(iso20022.AccountAndBalance36)
	return c.AccountDetails
}

func (c *CorporateActionMovementReversalAdviceV07) AddCorporateActionDetails() *iso20022.CorporateAction33 {
	c.CorporateActionDetails = new(iso20022.CorporateAction33)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementReversalAdviceV07) AddCorporateActionConfirmationDetails() *iso20022.CorporateActionOption119 {
	c.CorporateActionConfirmationDetails = new(iso20022.CorporateActionOption119)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementReversalAdviceV07) AddAdditionalInformation() *iso20022.CorporateActionNarrative31 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative31)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementReversalAdviceV07) AddIssuerAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV07) AddPayingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV07) AddSubPayingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
