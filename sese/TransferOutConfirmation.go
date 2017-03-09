package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300101 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.003.001.01 Document"`
	Message *TransferOutConfirmation `xml:"sese.003.001.01"`
}

func (d *Document00300101) AddMessage() *TransferOutConfirmation {
	d.Message = new(TransferOutConfirmation)
	return d.Message
}

// Scope
// The TransferOutConfirmation message is sent by an executing party to the instructing party or the instructing party's designated agent.
// This message is used to confirm the delivery of a financial instrument, free of payment, at a given date, to a specified party. This message can be used to confirm the transfer of a financial instrument to an own account or to a third party.
// Usage
// The TransferOutConfirmation message is used by an executing party to confirm to the instructing party that the withdrawal of a financial instrument from the owner's account and its delivery to another own account, or to a third party, has taken place.
type TransferOutConfirmation struct {

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *iso20022.Transfer2 `xml:"TrfDtls"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *iso20022.FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *iso20022.InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *iso20022.ReceiveInformation2 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOutConfirmation) AddRelatedReference() *iso20022.AdditionalReference2 {
	t.RelatedReference = new(iso20022.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutConfirmation) AddPoolReference() *iso20022.AdditionalReference2 {
	t.PoolReference = new(iso20022.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutConfirmation) AddPreviousReference() *iso20022.AdditionalReference2 {
	t.PreviousReference = new(iso20022.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutConfirmation) AddTransferDetails() *iso20022.Transfer2 {
	t.TransferDetails = new(iso20022.Transfer2)
	return t.TransferDetails
}

func (t *TransferOutConfirmation) AddFinancialInstrumentDetails() *iso20022.FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(iso20022.FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferOutConfirmation) AddAccountDetails() *iso20022.InvestmentAccount10 {
	t.AccountDetails = new(iso20022.InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferOutConfirmation) AddSettlementDetails() *iso20022.ReceiveInformation2 {
	t.SettlementDetails = new(iso20022.ReceiveInformation2)
	return t.SettlementDetails
}

func (t *TransferOutConfirmation) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
