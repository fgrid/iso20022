package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.02 Document"`
	Message *TransferInConfirmationV02 `xml:"TrfInConfV02"`
}

func (d *Document00700102) AddMessage() *TransferInConfirmationV02 {
	d.Message = new(TransferInConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the TransferInConfirmation message to the instructing party, eg, an investment manager or its authorised representative, to confirm the receipt of a financial instrument, free of payment, on a given date, from a specified party.
// This message may also be used to confirm the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInConfirmation message is used to confirm receipt of a financial instrument, either from another account owned by the instructing party or from a third party. The reference of the transfer confirmation is identified in TransferConfirmationReference.
// The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferInConfirmationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *iso20022.Transfer7 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *iso20022.FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *iso20022.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation4 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInConfirmationV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInConfirmationV02) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInConfirmationV02) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInConfirmationV02) AddTransferDetails() *iso20022.Transfer7 {
	t.TransferDetails = new(iso20022.Transfer7)
	return t.TransferDetails
}

func (t *TransferInConfirmationV02) AddFinancialInstrumentDetails() *iso20022.FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(iso20022.FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferInConfirmationV02) AddAccountDetails() *iso20022.InvestmentAccount22 {
	t.AccountDetails = new(iso20022.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferInConfirmationV02) AddSettlementDetails() *iso20022.DeliverInformation4 {
	t.SettlementDetails = new(iso20022.DeliverInformation4)
	return t.SettlementDetails
}

func (t *TransferInConfirmationV02) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInConfirmationV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
