package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.005.001.02 Document"`
	Message *AccountingStatementOfHoldingsCancellationV02 `xml:"AcctgStmtOfHldgsCxlV02"`
}

func (d *Document00500102) AddMessage() *AccountingStatementOfHoldingsCancellationV02 {
	d.Message = new(AccountingStatementOfHoldingsCancellationV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent, sends the AccountingStatementofHoldingsCancellation message to the account owner, for example, a fund manager or an account owner's designated agent, to notify the cancellation of a previously sent AccountingStatementOfHoldings message.
// Usage
// The AccountingStatementOfHoldingsCancellation message is used to cancel a previously sent AccountingStatementOfHoldings message. This message must contain the reference of the message to be cancelled.
// This message may also contain all the details of the message to be cancelled, but this is not recommended.
type AccountingStatementOfHoldingsCancellationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// The Accounting Statement of Holdings message to cancel.
	StatementToBeCancelled *iso20022.AccountingStatementOfHoldings2 `xml:"StmtToBeCanc,omitempty"`

}


func (a *AccountingStatementOfHoldingsCancellationV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountingStatementOfHoldingsCancellationV02) AddPreviousReference() *iso20022.AdditionalReference2 {
	a.PreviousReference = new(iso20022.AdditionalReference2)
	return a.PreviousReference
}

func (a *AccountingStatementOfHoldingsCancellationV02) AddRelatedReference() *iso20022.AdditionalReference2 {
	a.RelatedReference = new(iso20022.AdditionalReference2)
	return a.RelatedReference
}

func (a *AccountingStatementOfHoldingsCancellationV02) AddMessagePagination() *iso20022.Pagination {
	a.MessagePagination = new(iso20022.Pagination)
	return a.MessagePagination
}

func (a *AccountingStatementOfHoldingsCancellationV02) AddStatementToBeCancelled() *iso20022.AccountingStatementOfHoldings2 {
	a.StatementToBeCancelled = new(iso20022.AccountingStatementOfHoldings2)
	return a.StatementToBeCancelled
}

