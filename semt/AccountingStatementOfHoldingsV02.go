package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.003.001.02 Document"`
	Message *AccountingStatementOfHoldingsV02 `xml:"AcctgStmtOfHldgsV02"`
}

func (d *Document00300102) AddMessage() *AccountingStatementOfHoldingsV02 {
	d.Message = new(AccountingStatementOfHoldingsV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent sends the AccountStatementOfHoldings message to the account owner, for example, a fund manager or an account owner's designated agent to provide detailed holdings of the portfolio at a specified moment in time.
// The message provides, at a moment in time, valuations of the portfolio together with details of each financial instrument holding.
// The message can be sent either audited or un-audited and may be provided on a trade date or settlement date basis.
// Usage
// The AccountingStatementOfHoldings message is used to provide valuation detail for each financial instrument held in a portfolio. The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner.
// This message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, the message can be used to either specify holdings at
// - the main account level, or,
// - the sub-account level.
// This message can be used to report where the financial instruments are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// The AccountingStatementOfHoldings message should not be used for trading purposes.
type AccountingStatementOfHoldingsV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// General information related to the accounting statement of holdings.
	StatementGeneralDetails *iso20022.Statement6 `xml:"StmtGnlDtls"`

	// The safekeeping or investment account.
	AccountDetails *iso20022.SafekeepingAccount2 `xml:"AcctDtls"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*iso20022.AggregateBalanceInformation3 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*iso20022.SubAccountIdentification3 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *iso20022.TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountingStatementOfHoldingsV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	a.MessageIdentification = new(iso20022.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountingStatementOfHoldingsV02) AddPreviousReference() *iso20022.AdditionalReference2 {
	newValue := new(iso20022.AdditionalReference2)
	a.PreviousReference = append(a.PreviousReference, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddRelatedReference() *iso20022.AdditionalReference2 {
	newValue := new(iso20022.AdditionalReference2)
	a.RelatedReference = append(a.RelatedReference, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddMessagePagination() *iso20022.Pagination {
	a.MessagePagination = new(iso20022.Pagination)
	return a.MessagePagination
}

func (a *AccountingStatementOfHoldingsV02) AddStatementGeneralDetails() *iso20022.Statement6 {
	a.StatementGeneralDetails = new(iso20022.Statement6)
	return a.StatementGeneralDetails
}

func (a *AccountingStatementOfHoldingsV02) AddAccountDetails() *iso20022.SafekeepingAccount2 {
	a.AccountDetails = new(iso20022.SafekeepingAccount2)
	return a.AccountDetails
}

func (a *AccountingStatementOfHoldingsV02) AddBalanceForAccount() *iso20022.AggregateBalanceInformation3 {
	newValue := new(iso20022.AggregateBalanceInformation3)
	a.BalanceForAccount = append(a.BalanceForAccount, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddSubAccountDetails() *iso20022.SubAccountIdentification3 {
	newValue := new(iso20022.SubAccountIdentification3)
	a.SubAccountDetails = append(a.SubAccountDetails, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddTotalValues() *iso20022.TotalValueInPageAndStatement {
	a.TotalValues = new(iso20022.TotalValueInPageAndStatement)
	return a.TotalValues
}

func (a *AccountingStatementOfHoldingsV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
