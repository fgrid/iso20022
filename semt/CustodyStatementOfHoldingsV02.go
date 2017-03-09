package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.002.001.02 Document"`
	Message *CustodyStatementOfHoldingsV02 `xml:"CtdyStmtOfHldgsV02"`
}

func (d *Document00200102) AddMessage() *CustodyStatementOfHoldingsV02 {
	d.Message = new(CustodyStatementOfHoldingsV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent sends the CustodyStatementOfHoldings message to the account owner, for example, a fund manager or an account owner's designated agent to provide detailed holdings of the portfolio at a specified moment in time.
// The message provides, at a moment in time, the quantity and identification of the financial instruments that the account servicer holds for the account owner. The message can also include availability and the location of holdings to facilitate trading and minimise settlement issues.
// The message reports all information per financial instrument, ie, when a financial instrument is held at multiple places of safekeeping, the total holdings for all locations can be provided.
// Usage
// The CustodyStatementOfHoldings message is used to provide detailed quantity and availability information for financial instrument holdings of a portfolio. The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner.
// This message can be also be used to report where the financial instruments are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// This message can reflect all outstanding holding information or may only contain changes since the previously sent statement.
// The CustodyStatementOfHoldings message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, this message can be used to either specify holdings at
// - the main account level, or,
// - the sub-account level.
type CustodyStatementOfHoldingsV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// General information related to the custody statement of holdings.
	StatementGeneralDetails *iso20022.Statement7 `xml:"StmtGnlDtls"`

	// The safekeeping or investment account.
	AccountDetails *iso20022.SafekeepingAccount2 `xml:"AcctDtls"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*iso20022.AggregateBalanceInformation4 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*iso20022.SubAccountIdentification5 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *iso20022.TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (c *CustodyStatementOfHoldingsV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	c.MessageIdentification = new(iso20022.MessageIdentification1)
	return c.MessageIdentification
}

func (c *CustodyStatementOfHoldingsV02) AddPreviousReference() *iso20022.AdditionalReference2 {
	newValue := new(iso20022.AdditionalReference2)
	c.PreviousReference = append(c.PreviousReference, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddRelatedReference() *iso20022.AdditionalReference2 {
	newValue := new(iso20022.AdditionalReference2)
	c.RelatedReference = append(c.RelatedReference, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddMessagePagination() *iso20022.Pagination {
	c.MessagePagination = new(iso20022.Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldingsV02) AddStatementGeneralDetails() *iso20022.Statement7 {
	c.StatementGeneralDetails = new(iso20022.Statement7)
	return c.StatementGeneralDetails
}

func (c *CustodyStatementOfHoldingsV02) AddAccountDetails() *iso20022.SafekeepingAccount2 {
	c.AccountDetails = new(iso20022.SafekeepingAccount2)
	return c.AccountDetails
}

func (c *CustodyStatementOfHoldingsV02) AddBalanceForAccount() *iso20022.AggregateBalanceInformation4 {
	newValue := new(iso20022.AggregateBalanceInformation4)
	c.BalanceForAccount = append(c.BalanceForAccount, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddSubAccountDetails() *iso20022.SubAccountIdentification5 {
	newValue := new(iso20022.SubAccountIdentification5)
	c.SubAccountDetails = append(c.SubAccountDetails, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddTotalValues() *iso20022.TotalValueInPageAndStatement {
	c.TotalValues = new(iso20022.TotalValueInPageAndStatement)
	return c.TotalValues
}

func (c *CustodyStatementOfHoldingsV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
