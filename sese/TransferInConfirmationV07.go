package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700107 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.07 Document"`
	Message *TransferInConfirmationV07 `xml:"TrfInConf"`
}

func (d *Document00700107) AddMessage() *TransferInConfirmationV07 {
	d.Message = new(TransferInConfirmationV07)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferInConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to confirm the receipt of a financial instrument, free of payment, on a given date, from a specified party.
// This message may also be used to confirm the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// This message may also be used as an advice, that is, the message is used to provide account information.
// Usage
// The TransferInConfirmation message is used to confirm receipt of a financial instrument, either from another account owned by the instructing party or from a third party. The reference of the transfer confirmation is identified in TransferConfirmationReference.
// The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferInConfirmationV07 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Function of the transfer in, that is, whether the transfer in message is used as a confirmation or as or an advice. The absence of Function indicates the message is a confirmation.
	Function *iso20022.TransferInFunction2Code `xml:"Fctn,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument. 
	TransferDetails []*iso20022.Transfer33 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *iso20022.InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation17 `xml:"SttlmDtls,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (t *TransferInConfirmationV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInConfirmationV07) AddPoolReference() *iso20022.AdditionalReference6 {
	t.PoolReference = new(iso20022.AdditionalReference6)
	return t.PoolReference
}

func (t *TransferInConfirmationV07) AddPreviousReference() *iso20022.AdditionalReference6 {
	t.PreviousReference = new(iso20022.AdditionalReference6)
	return t.PreviousReference
}

func (t *TransferInConfirmationV07) AddRelatedReference() *iso20022.AdditionalReference6 {
	t.RelatedReference = new(iso20022.AdditionalReference6)
	return t.RelatedReference
}

func (t *TransferInConfirmationV07) SetFunction(value string) {
	t.Function = (*iso20022.TransferInFunction2Code)(&value)
}

func (t *TransferInConfirmationV07) SetMasterReference(value string) {
	t.MasterReference = (*iso20022.Max35Text)(&value)
}

func (t *TransferInConfirmationV07) AddTransferDetails() *iso20022.Transfer33 {
	newValue := new (iso20022.Transfer33)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInConfirmationV07) AddAccountDetails() *iso20022.InvestmentAccount56 {
	t.AccountDetails = new(iso20022.InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferInConfirmationV07) AddSettlementDetails() *iso20022.DeliverInformation17 {
	t.SettlementDetails = new(iso20022.DeliverInformation17)
	return t.SettlementDetails
}

func (t *TransferInConfirmationV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInConfirmationV07) AddCopyDetails() *iso20022.CopyInformation4 {
	t.CopyDetails = new(iso20022.CopyInformation4)
	return t.CopyDetails
}

func (t *TransferInConfirmationV07) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

