package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500104 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.04 Document"`
	Message *InterestPaymentStatementV04 `xml:"IntrstPmtStmt"`
}

func (d *Document01500104) AddMessage() *InterestPaymentStatementV04 {
	d.Message = new(InterestPaymentStatementV04)
	return d.Message
}

// Scope
// The InterestPaymentStatement message is sent by either the collateral giver or its collateral manager to the collateral taker or its collateral manager. It is used to report the interest amounts calculated based on the effective posted collateral amount, over a specific period of time agreed by both parties.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The InterestPaymentStatement message is used for reporting the interest per period on collateral held.
type InterestPaymentStatementV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Provides general information on the report such as the statement identification.
	StatementParameters *iso20022.Statement32 `xml:"StmtParams"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn,omitempty"`

	// Provides details on the interest statement.
	InterestStatement *iso20022.InterestStatement4 `xml:"IntrstStmt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InterestPaymentStatementV04) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InterestPaymentStatementV04) AddAgreement() *iso20022.Agreement4 {
	i.Agreement = new(iso20022.Agreement4)
	return i.Agreement
}

func (i *InterestPaymentStatementV04) AddObligation() *iso20022.Obligation4 {
	i.Obligation = new(iso20022.Obligation4)
	return i.Obligation
}

func (i *InterestPaymentStatementV04) AddStatementParameters() *iso20022.Statement32 {
	i.StatementParameters = new(iso20022.Statement32)
	return i.StatementParameters
}

func (i *InterestPaymentStatementV04) AddPagination() *iso20022.Pagination {
	i.Pagination = new(iso20022.Pagination)
	return i.Pagination
}

func (i *InterestPaymentStatementV04) AddInterestStatement() *iso20022.InterestStatement4 {
	i.InterestStatement = new(iso20022.InterestStatement4)
	return i.InterestStatement
}

func (i *InterestPaymentStatementV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
