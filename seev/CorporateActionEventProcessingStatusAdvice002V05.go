package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03200205 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.032.002.05 Document"`
	Message *CorporateActionEventProcessingStatusAdvice002V05 `xml:"CorpActnEvtPrcgStsAdvc"`
}

func (d *Document03200205) AddMessage() *CorporateActionEventProcessingStatusAdvice002V05 {
	d.Message = new(CorporateActionEventProcessingStatusAdvice002V05)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionEventProcessingStatusAdvice message to an account owner or its designated agent to report processing status of a corporate action event.
// The account servicer uses this message to provide a reason as to why a corporate action event has not been completed by the announced payment dates.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate), 
// using the relevant elements in the business application header (BAH).
type CorporateActionEventProcessingStatusAdvice002V05 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification17 `xml:"NtfctnId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification34 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation94 `xml:"CorpActnGnlInf"`

	// Information about the status of a corporate action.
	EventProcessingStatus []*iso20022.EventProcessingStatus4Choice `xml:"EvtPrcgSts"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative19 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *CorporateActionEventProcessingStatusAdvice002V05) AddNotificationIdentification() *iso20022.DocumentIdentification17 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification17)
	return c.NotificationIdentification
}

func (c *CorporateActionEventProcessingStatusAdvice002V05) AddOtherDocumentIdentification() *iso20022.DocumentIdentification34 {
	newValue := new (iso20022.DocumentIdentification34)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdvice002V05) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation94 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation94)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionEventProcessingStatusAdvice002V05) AddEventProcessingStatus() *iso20022.EventProcessingStatus4Choice {
	newValue := new (iso20022.EventProcessingStatus4Choice)
	c.EventProcessingStatus = append(c.EventProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdvice002V05) AddAdditionalInformation() *iso20022.CorporateActionNarrative19 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative19)
	return c.AdditionalInformation
}

func (c *CorporateActionEventProcessingStatusAdvice002V05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

