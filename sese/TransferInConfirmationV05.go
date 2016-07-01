package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.05 Document"`
	Message *TransferInConfirmationV05 `xml:"TrfInConf"`
}

func (d *Document00700105) AddMessage() *TransferInConfirmationV05 {
	d.Message = new(TransferInConfirmationV05)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferInConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to confirm the receipt of a financial instrument, free of payment, on a given date, from a specified party.
// This message may also be used to confirm the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInConfirmation message is used to confirm receipt of a financial instrument, either from another account owned by the instructing party or from a third party. The reference of the transfer confirmation is identified in TransferConfirmationReference.
// The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferInConfirmationV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument. 
	TransferDetails []*iso20022.Transfer29 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *iso20022.InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation12 `xml:"SttlmDtls,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (t *TransferInConfirmationV05) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInConfirmationV05) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInConfirmationV05) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInConfirmationV05) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInConfirmationV05) SetMasterReference(value string) {
	t.MasterReference = (*iso20022.Max35Text)(&value)
}

func (t *TransferInConfirmationV05) AddTransferDetails() *iso20022.Transfer29 {
	newValue := new (iso20022.Transfer29)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInConfirmationV05) AddAccountDetails() *iso20022.InvestmentAccount40 {
	t.AccountDetails = new(iso20022.InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferInConfirmationV05) AddSettlementDetails() *iso20022.DeliverInformation12 {
	t.SettlementDetails = new(iso20022.DeliverInformation12)
	return t.SettlementDetails
}

func (t *TransferInConfirmationV05) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInConfirmationV05) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInConfirmationV05) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

