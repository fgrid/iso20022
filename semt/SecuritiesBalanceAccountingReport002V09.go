package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300209 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.003.002.09 Document"`
	Message *SecuritiesBalanceAccountingReport002V09 `xml:"SctiesBalAcctgRpt"`
}

func (d *Document00300209) AddMessage() *SecuritiesBalanceAccountingReport002V09 {
	d.Message = new(SecuritiesBalanceAccountingReport002V09)
	return d.Message
}

// Scope 
// An account servicer sends a SecuritiesBalanceAccountingReport to an account owner to provide, at a moment in time, valuations of the portfolio together with details of each financial instrument holding.
// The account servicer/owner relationship may be:
// - an accounting agent acting on behalf of an account owner, or
// - a transfer agent acting on behalf of a fund manager or an account owner's designated agent.
// 
// Usage
// The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner.
// The message can be sent either audited or un-audited and may be provided on a trade date, contractual or settlement date basis. 
// This message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, the message can be used to either specify holdings at
// - the main account level, or, 
// - the sub-account level.
// This message can be used to report where the financial instruments are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated. 
// The SecuritiesBalanceAccountingReport message should not be used for trading purposes.
// There may be one or more intermediary parties, for example, an intermediary or a concentrator between the account owner and account servicer.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesBalanceAccountingReport002V09 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *iso20022.Statement51 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *iso20022.PartyIdentification111 `xml:"AcctSvcr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount33 `xml:"SfkpgAcct"`

	// Information about the party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*iso20022.Intermediary37 `xml:"IntrmyInf,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*iso20022.AggregateBalanceInformation32 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*iso20022.SubAccountIdentification45 `xml:"SubAcctDtls,omitempty"`

	// Total valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyTotalAmounts *iso20022.TotalValueInPageAndStatement4 `xml:"AcctBaseCcyTtlAmts,omitempty"`

	// Total valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyTotalAmounts *iso20022.TotalValueInPageAndStatement4 `xml:"AltrnRptgCcyTtlAmts,omitempty"`

}


func (s *SecuritiesBalanceAccountingReport002V09) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesBalanceAccountingReport002V09) AddStatementGeneralDetails() *iso20022.Statement51 {
	s.StatementGeneralDetails = new(iso20022.Statement51)
	return s.StatementGeneralDetails
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAccountServicer() *iso20022.PartyIdentification111 {
	s.AccountServicer = new(iso20022.PartyIdentification111)
	return s.AccountServicer
}

func (s *SecuritiesBalanceAccountingReport002V09) AddSafekeepingAccount() *iso20022.SecuritiesAccount33 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount33)
	return s.SafekeepingAccount
}

func (s *SecuritiesBalanceAccountingReport002V09) AddIntermediaryInformation() *iso20022.Intermediary37 {
	newValue := new (iso20022.Intermediary37)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReport002V09) AddBalanceForAccount() *iso20022.AggregateBalanceInformation32 {
	newValue := new (iso20022.AggregateBalanceInformation32)
	s.BalanceForAccount = append(s.BalanceForAccount, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReport002V09) AddSubAccountDetails() *iso20022.SubAccountIdentification45 {
	newValue := new (iso20022.SubAccountIdentification45)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAccountBaseCurrencyTotalAmounts() *iso20022.TotalValueInPageAndStatement4 {
	s.AccountBaseCurrencyTotalAmounts = new(iso20022.TotalValueInPageAndStatement4)
	return s.AccountBaseCurrencyTotalAmounts
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAlternateReportingCurrencyTotalAmounts() *iso20022.TotalValueInPageAndStatement4 {
	s.AlternateReportingCurrencyTotalAmounts = new(iso20022.TotalValueInPageAndStatement4)
	return s.AlternateReportingCurrencyTotalAmounts
}

