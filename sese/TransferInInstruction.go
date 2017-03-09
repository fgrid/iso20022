package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.01 Document"`
	Message *TransferInInstruction `xml:"sese.005.001.01"`
}

func (d *Document00500101) AddMessage() *TransferInInstruction {
	d.Message = new(TransferInInstruction)
	return d.Message
}

// Scope
// The TransferInInstruction message is sent by an instructing party, or an instructing party's designated agent, to the executing party.
// This message is used to instruct the receipt of a financial instrument, free of payment, at a given date, from a specified party. This message may also be used to instruct the receipt of a financial instrument from an own account or from a third party.
// Usage
// The TransferInInstruction message is used by an instructing party to instruct the executing party to receive a financial instrument from another account, either owned by the instructing party or by a third party.
type TransferInInstruction struct {

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message which was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *iso20022.Transfer3 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *iso20022.FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *iso20022.InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation1 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInInstruction) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInInstruction) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInInstruction) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInInstruction) AddTransferDetails() *iso20022.Transfer3 {
	t.TransferDetails = new(iso20022.Transfer3)
	return t.TransferDetails
}

func (t *TransferInInstruction) AddFinancialInstrumentDetails() *iso20022.FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(iso20022.FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferInInstruction) AddAccountDetails() *iso20022.InvestmentAccount10 {
	t.AccountDetails = new(iso20022.InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferInInstruction) AddSettlementDetails() *iso20022.DeliverInformation1 {
	t.SettlementDetails = new(iso20022.DeliverInformation1)
	return t.SettlementDetails
}

func (t *TransferInInstruction) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
