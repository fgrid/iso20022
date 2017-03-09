package seev

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03300104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.04 Document"`
	Message *CorporateActionInstructionV04 `xml:"CorpActnInstr"`
}

func (d *Document03300104) AddMessage() *CorporateActionInstructionV04 {
	d.Message = new(CorporateActionInstructionV04)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstruction message to an account servicer to instruct election on a corporate action event.
// This message is used to provide the custodian with instructions on how the account owner wishes to proceed with a corporate action event. Instructions include investment decisions regarding the exercise of rights issues, the election of stock or cash when the option is available, and decisions on the conversion or tendering of securities.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionV04 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *iso20022.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent cancelled instruction document.
	CancelledInstructionIdentification *iso20022.DocumentIdentification15 `xml:"CancInstrId,omitempty"`

	// Identification of a previously sent instruction cancellation request document.
	InstructionCancellationRequestIdentification *iso20022.DocumentIdentification15 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation55 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountAndBalance17 `xml:"AcctDtls"`

	// Provides information about the beneficial owner of the securities.
	BeneficialOwnerDetails []*iso20022.PartyIdentification56 `xml:"BnfclOwnrDtls,omitempty"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *iso20022.CorporateActionOption57 `xml:"CorpActnInstr"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative21 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionV04) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*iso20022.YesNoIndicator)(&value)
}

func (c *CorporateActionInstructionV04) AddCancelledInstructionIdentification() *iso20022.DocumentIdentification15 {
	c.CancelledInstructionIdentification = new(iso20022.DocumentIdentification15)
	return c.CancelledInstructionIdentification
}

func (c *CorporateActionInstructionV04) AddInstructionCancellationRequestIdentification() *iso20022.DocumentIdentification15 {
	c.InstructionCancellationRequestIdentification = new(iso20022.DocumentIdentification15)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionV04) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new(iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionV04) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new(iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionInstructionV04) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation55 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation55)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionV04) AddAccountDetails() *iso20022.AccountAndBalance17 {
	c.AccountDetails = new(iso20022.AccountAndBalance17)
	return c.AccountDetails
}

func (c *CorporateActionInstructionV04) AddBeneficialOwnerDetails() *iso20022.PartyIdentification56 {
	newValue := new(iso20022.PartyIdentification56)
	c.BeneficialOwnerDetails = append(c.BeneficialOwnerDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionV04) AddCorporateActionInstruction() *iso20022.CorporateActionOption57 {
	c.CorporateActionInstruction = new(iso20022.CorporateActionOption57)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionV04) AddAdditionalInformation() *iso20022.CorporateActionNarrative21 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative21)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
