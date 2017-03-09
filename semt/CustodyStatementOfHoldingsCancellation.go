package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.004.001.01 Document"`
	Message *CustodyStatementOfHoldingsCancellation `xml:"semt.004.001.01"`
}

func (d *Document00400101) AddMessage() *CustodyStatementOfHoldingsCancellation {
	d.Message = new(CustodyStatementOfHoldingsCancellation)
	return d.Message
}

// Scope
// The CustodyStatementOfHoldingsCancellation message is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a local agent (sub-custodian) acting on behalf of its global custodian customer, a custodian acting on behalf of an investment management institution or a broker/dealer, a fund administrator or fund intermediary, trustee or registrar, etc.
// This message is used to cancel a previously sent CustodyStatementOfHoldings message.
// Usage
// The CustodyStatementOfHoldingsCancellation message is sent by an account servicer to the account owner to cancel a previously sent CustodyStatementOfHoldings message.
// This message must contain the reference of the message to be cancelled. This message may also contain details of the message to be cancelled, but this is not recommended.
type CustodyStatementOfHoldingsCancellation struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Number used to sequence pages when it is not possible for data to be conveyed in a single message and the data has to be split across several pages (messages).
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// The Custody Statement of Holdings message to cancel.
	StatementToBeCancelled *iso20022.CustodyStatementOfHoldings1 `xml:"StmtToBeCanc,omitempty"`
}

func (c *CustodyStatementOfHoldingsCancellation) AddPreviousReference() *iso20022.AdditionalReference2 {
	c.PreviousReference = new(iso20022.AdditionalReference2)
	return c.PreviousReference
}

func (c *CustodyStatementOfHoldingsCancellation) AddRelatedReference() *iso20022.AdditionalReference2 {
	c.RelatedReference = new(iso20022.AdditionalReference2)
	return c.RelatedReference
}

func (c *CustodyStatementOfHoldingsCancellation) AddMessagePagination() *iso20022.Pagination {
	c.MessagePagination = new(iso20022.Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldingsCancellation) AddStatementToBeCancelled() *iso20022.CustodyStatementOfHoldings1 {
	c.StatementToBeCancelled = new(iso20022.CustodyStatementOfHoldings1)
	return c.StatementToBeCancelled
}
