package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.005.001.01 Document"`
	Message *AccountingStatementOfHoldingsCancellation `xml:"semt.005.001.01"`
}

func (d *Document00500101) AddMessage() *AccountingStatementOfHoldingsCancellation {
	d.Message = new(AccountingStatementOfHoldingsCancellation)
	return d.Message
}

// Scope
// The AccountingStatementOfHoldingsCancellation message is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a local agent (sub-custodian) acting on behalf of its global custodian customer, a custodian acting on behalf of an investment management institution or a broker/dealer, a fund administrator or fund intermediary, trustee or registrar.
// This message is used to cancel a previously sent AccountingStatementOfHoldings message.
// Usage
// The AccountingStatementOfHoldingsCancellation message is sent by an account servicer to the account owner to cancel a previously sent AccountingStatementOfHoldings message.
// This message must contain the reference of the message to be cancelled. This message may also contain all the details of the message to be cancelled, but this is not recommended.
type AccountingStatementOfHoldingsCancellation struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// The Accounting Statement of Holdings message to cancel.
	StatementToBeCancelled *iso20022.AccountingStatementOfHoldings1 `xml:"StmtToBeCanc,omitempty"`

}


func (a *AccountingStatementOfHoldingsCancellation) AddPreviousReference() *iso20022.AdditionalReference2 {
	a.PreviousReference = new(iso20022.AdditionalReference2)
	return a.PreviousReference
}

func (a *AccountingStatementOfHoldingsCancellation) AddRelatedReference() *iso20022.AdditionalReference2 {
	a.RelatedReference = new(iso20022.AdditionalReference2)
	return a.RelatedReference
}

func (a *AccountingStatementOfHoldingsCancellation) AddMessagePagination() *iso20022.Pagination {
	a.MessagePagination = new(iso20022.Pagination)
	return a.MessagePagination
}

func (a *AccountingStatementOfHoldingsCancellation) AddStatementToBeCancelled() *iso20022.AccountingStatementOfHoldings1 {
	a.StatementToBeCancelled = new(iso20022.AccountingStatementOfHoldings1)
	return a.StatementToBeCancelled
}

