package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.02 Document"`
	Message *TransferInInstructionV02 `xml:"TrfInInstrV02"`
}

func (d *Document00500102) AddMessage() *TransferInInstructionV02 {
	d.Message = new(TransferInInstructionV02)
	return d.Message
}

// Scope
// An instructing party, eg, an investment manager or its authorised representative, sends the TransferInInstruction message to the executing party, eg, a transfer agent, to instruct the receipt of a financial instrument, free of payment, on a given date from a specified party.
// This message may also be used to instruct the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInInstruction message is used to instruct the receipt of a financial instrument from another account, either owned by the instructing party or by a third party.
type TransferInInstructionV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message which was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *iso20022.Transfer5 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *iso20022.FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *iso20022.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation3 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInInstructionV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInInstructionV02) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInInstructionV02) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInInstructionV02) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInInstructionV02) AddTransferDetails() *iso20022.Transfer5 {
	t.TransferDetails = new(iso20022.Transfer5)
	return t.TransferDetails
}

func (t *TransferInInstructionV02) AddFinancialInstrumentDetails() *iso20022.FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(iso20022.FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferInInstructionV02) AddAccountDetails() *iso20022.InvestmentAccount22 {
	t.AccountDetails = new(iso20022.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferInInstructionV02) AddSettlementDetails() *iso20022.DeliverInformation3 {
	t.SettlementDetails = new(iso20022.DeliverInformation3)
	return t.SettlementDetails
}

func (t *TransferInInstructionV02) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInInstructionV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
