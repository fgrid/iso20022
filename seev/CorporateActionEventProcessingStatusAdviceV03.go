package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.032.001.03 Document"`
	Message *CorporateActionEventProcessingStatusAdviceV03 `xml:"CorpActnEvtPrcgStsAdvc"`
}

func (d *Document03200103) AddMessage() *CorporateActionEventProcessingStatusAdviceV03 {
	d.Message = new(CorporateActionEventProcessingStatusAdviceV03)
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
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionEventProcessingStatusAdviceV03 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification9 `xml:"NtfctnId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification14 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation34 `xml:"CorpActnGnlInf"`

	// Information about the status of a corporate action.
	EventProcessingStatus []*iso20022.EventProcessingStatus1Choice `xml:"EvtPrcgSts"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *CorporateActionEventProcessingStatusAdviceV03) AddNotificationIdentification() *iso20022.DocumentIdentification9 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification9)
	return c.NotificationIdentification
}

func (c *CorporateActionEventProcessingStatusAdviceV03) AddOtherDocumentIdentification() *iso20022.DocumentIdentification14 {
	newValue := new (iso20022.DocumentIdentification14)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdviceV03) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation34 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation34)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionEventProcessingStatusAdviceV03) AddEventProcessingStatus() *iso20022.EventProcessingStatus1Choice {
	newValue := new (iso20022.EventProcessingStatus1Choice)
	c.EventProcessingStatus = append(c.EventProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdviceV03) AddAdditionalInformation() *iso20022.CorporateActionNarrative10 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionEventProcessingStatusAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

