package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01900105 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.019.001.05 Document"`
	Message *SecuritiesSettlementTransactionAllegementReportV05 `xml:"SctiesSttlmTxAllgmtRpt"`
}

func (d *Document01900105) AddMessage() *SecuritiesSettlementTransactionAllegementReportV05 {
	d.Message = new(SecuritiesSettlementTransactionAllegementReportV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementReport to an account owner to provide, at a specified time, the status and details of pending settlement allegements, for all or selected securities in a specified safekeeping account.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionAllegementReportV05 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement39 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Details of the allegement.
	AllegementDetails []*iso20022.SecuritiesTradeDetails49 `xml:"AllgmtDtls,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddStatementGeneralDetails() *iso20022.Statement39 {
	s.StatementGeneralDetails = new(iso20022.Statement39)
	return s.StatementGeneralDetails
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddAccountOwner() *iso20022.PartyIdentification98 {
	s.AccountOwner = new(iso20022.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddSafekeepingAccount() *iso20022.SecuritiesAccount24 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddAllegementDetails() *iso20022.SecuritiesTradeDetails49 {
	newValue := new(iso20022.SecuritiesTradeDetails49)
	s.AllegementDetails = append(s.AllegementDetails, newValue)
	return newValue
}
