package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03800103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.03 Document"`
	Message *CorporateActionNarrativeV03 `xml:"CorpActnNrrtv"`
}

func (d *Document03800103) AddMessage() *CorporateActionNarrativeV03 {
	d.Message = new(CorporateActionNarrativeV03)
	return d.Message
}

// Scope
// The CorporateActionNarrative message is sent between an account servicer and an account owner or its designated agent to cater for tax reclaims, restrictions, documentation requirements. This message is bi-directional.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionNarrativeV03 struct {

	// General information about the safekeeping account and the account owner.
	AccountDetails *iso20022.AccountIdentification14Choice `xml:"AcctDtls,omitempty"`

	// Provides information about the securitised right for entitlement.
	UnderlyingSecurity *iso20022.SecurityIdentification14 `xml:"UndrlygScty,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation40 `xml:"CorpActnGnlInf"`

	// Provides additional information.
	AdditionalInformation *iso20022.UpdatedAdditionalInformation2 `xml:"AddtlInf"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *CorporateActionNarrativeV03) AddAccountDetails() *iso20022.AccountIdentification14Choice {
	c.AccountDetails = new(iso20022.AccountIdentification14Choice)
	return c.AccountDetails
}

func (c *CorporateActionNarrativeV03) AddUnderlyingSecurity() *iso20022.SecurityIdentification14 {
	c.UnderlyingSecurity = new(iso20022.SecurityIdentification14)
	return c.UnderlyingSecurity
}

func (c *CorporateActionNarrativeV03) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation40 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation40)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNarrativeV03) AddAdditionalInformation() *iso20022.UpdatedAdditionalInformation2 {
	c.AdditionalInformation = new(iso20022.UpdatedAdditionalInformation2)
	return c.AdditionalInformation
}

func (c *CorporateActionNarrativeV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

