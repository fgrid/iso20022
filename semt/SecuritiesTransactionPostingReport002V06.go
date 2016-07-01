package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700206 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.002.06 Document"`
	Message *SecuritiesTransactionPostingReport002V06 `xml:"SctiesTxPstngRpt"`
}

func (d *Document01700206) AddMessage() *SecuritiesTransactionPostingReport002V06 {
	d.Message = new(SecuritiesTransactionPostingReport002V06)
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
type SecuritiesTransactionPostingReport002V06 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *iso20022.Statement56 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*iso20022.FinancialInstrumentDetails23 `xml:"FinInstrmDtls,omitempty"`

	// Details at sub-account level.
	SubAccountDetails []*iso20022.SubAccountIdentification47 `xml:"SubAcctDtls,omitempty"`

}


func (s *SecuritiesTransactionPostingReport002V06) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPostingReport002V06) AddStatementGeneralDetails() *iso20022.Statement56 {
	s.StatementGeneralDetails = new(iso20022.Statement56)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPostingReport002V06) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPostingReport002V06) AddSafekeepingAccount() *iso20022.SecuritiesAccount27 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount27)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPostingReport002V06) AddFinancialInstrumentDetails() *iso20022.FinancialInstrumentDetails23 {
	newValue := new (iso20022.FinancialInstrumentDetails23)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

func (s *SecuritiesTransactionPostingReport002V06) AddSubAccountDetails() *iso20022.SubAccountIdentification47 {
	newValue := new (iso20022.SubAccountIdentification47)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

