package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01800107 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.018.001.07 Document"`
	Message *SecuritiesTransactionPendingReportV07 `xml:"SctiesTxPdgRpt"`
}

func (d *Document01800107) AddMessage() *SecuritiesTransactionPendingReportV07 {
	d.Message = new(SecuritiesTransactionPendingReportV07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPendingReport to an account owner to provide, as at a specified time, the details of pending increases and decreases of holdings, for all or selected securities in a specified safekeeping account, for all or selected reasons why the transaction is pending.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The statement may also include future settlement or forward transactions which have become binding on the account owner.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesTransactionPendingReportV07 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *iso20022.Statement41 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Status information.
	Status []*iso20022.StatusAndReason27 `xml:"Sts,omitempty"`

	// Details of the transactions reported.
	Transactions []*iso20022.Transaction47 `xml:"Txs,omitempty"`
}

func (s *SecuritiesTransactionPendingReportV07) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPendingReportV07) AddStatementGeneralDetails() *iso20022.Statement41 {
	s.StatementGeneralDetails = new(iso20022.Statement41)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPendingReportV07) AddAccountOwner() *iso20022.PartyIdentification98 {
	s.AccountOwner = new(iso20022.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPendingReportV07) AddSafekeepingAccount() *iso20022.SecuritiesAccount24 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPendingReportV07) AddStatus() *iso20022.StatusAndReason27 {
	newValue := new(iso20022.StatusAndReason27)
	s.Status = append(s.Status, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReportV07) AddTransactions() *iso20022.Transaction47 {
	newValue := new(iso20022.Transaction47)
	s.Transactions = append(s.Transactions, newValue)
	return newValue
}
