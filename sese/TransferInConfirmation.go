package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.01 Document"`
	Message *TransferInConfirmation `xml:"sese.007.001.01"`
}

func (d *Document00700101) AddMessage() *TransferInConfirmation {
	d.Message = new(TransferInConfirmation)
	return d.Message
}

// Scope
// TheTransferInConfirmation message is sent by an executing party to the instructing party or the instructing party's designated agent.
// This message is used to confirm the receipt of a financial instrument, free of payment, at a given date, from a specified party. This message can also be used to confirm the transfer a financial instrument from an own account or from a third party.
// Usage
// TheTransferInConfirmation message is used by an executing party to confirm to the instructing party receipt of a financial instrument, either from another account owned by the instructing party or from a third party.
type TransferInConfirmation struct {

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *iso20022.Transfer4 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *iso20022.FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *iso20022.InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation2 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInConfirmation) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInConfirmation) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInConfirmation) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInConfirmation) AddTransferDetails() *iso20022.Transfer4 {
	t.TransferDetails = new(iso20022.Transfer4)
	return t.TransferDetails
}

func (t *TransferInConfirmation) AddFinancialInstrumentDetails() *iso20022.FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(iso20022.FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferInConfirmation) AddAccountDetails() *iso20022.InvestmentAccount10 {
	t.AccountDetails = new(iso20022.InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferInConfirmation) AddSettlementDetails() *iso20022.DeliverInformation2 {
	t.SettlementDetails = new(iso20022.DeliverInformation2)
	return t.SettlementDetails
}

func (t *TransferInConfirmation) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
