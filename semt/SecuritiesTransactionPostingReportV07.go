package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700107 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.07 Document"`
	Message *SecuritiesTransactionPostingReportV07 `xml:"SctiesTxPstngRpt"`
}

func (d *Document01700107) AddMessage() *SecuritiesTransactionPostingReportV07 {
	d.Message = new(SecuritiesTransactionPostingReportV07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPostingReport to an account owner to provide the details of increases and decreases of holdings which occurred during a specified period, for all or selected securities in the specified safekeeping account or sub-safekeeping account which the account servicer holds for the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// This message may be used as a trade date based or a settlement date based statement.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesTransactionPostingReportV07 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *iso20022.Statement44 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*iso20022.FinancialInstrumentDetails25 `xml:"FinInstrmDtls,omitempty"`

	// Details at sub-account level.
	SubAccountDetails []*iso20022.SubAccountIdentification49 `xml:"SubAcctDtls,omitempty"`
}

func (s *SecuritiesTransactionPostingReportV07) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPostingReportV07) AddStatementGeneralDetails() *iso20022.Statement44 {
	s.StatementGeneralDetails = new(iso20022.Statement44)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPostingReportV07) AddAccountOwner() *iso20022.PartyIdentification98 {
	s.AccountOwner = new(iso20022.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPostingReportV07) AddSafekeepingAccount() *iso20022.SecuritiesAccount24 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPostingReportV07) AddFinancialInstrumentDetails() *iso20022.FinancialInstrumentDetails25 {
	newValue := new(iso20022.FinancialInstrumentDetails25)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

func (s *SecuritiesTransactionPostingReportV07) AddSubAccountDetails() *iso20022.SubAccountIdentification49 {
	newValue := new(iso20022.SubAccountIdentification49)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}
