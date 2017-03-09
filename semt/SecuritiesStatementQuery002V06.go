package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100206 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.021.002.06 Document"`
	Message *SecuritiesStatementQuery002V06 `xml:"SctiesStmtQry"`
}

func (d *Document02100206) AddMessage() *SecuritiesStatementQuery002V06 {
	d.Message = new(SecuritiesStatementQuery002V06)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesStatementQuery to an account servicer to request any existing securities statement.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct to a central securities depository or another settlement market infrastructure.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesStatementQuery002V06 struct {

	// Description of the statement requested.
	StatementRequested *iso20022.DocumentNumber14 `xml:"StmtReqd"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement54 `xml:"StmtGnlDtls,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Additional specific query criteria.
	AdditionalQueryParameters []*iso20022.AdditionalQueryParameters12 `xml:"AddtlQryParams,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatementQuery002V06) AddStatementRequested() *iso20022.DocumentNumber14 {
	s.StatementRequested = new(iso20022.DocumentNumber14)
	return s.StatementRequested
}

func (s *SecuritiesStatementQuery002V06) AddStatementGeneralDetails() *iso20022.Statement54 {
	s.StatementGeneralDetails = new(iso20022.Statement54)
	return s.StatementGeneralDetails
}

func (s *SecuritiesStatementQuery002V06) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesStatementQuery002V06) AddSafekeepingAccount() *iso20022.SecuritiesAccount27 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount27)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatementQuery002V06) AddAdditionalQueryParameters() *iso20022.AdditionalQueryParameters12 {
	newValue := new(iso20022.AdditionalQueryParameters12)
	s.AdditionalQueryParameters = append(s.AdditionalQueryParameters, newValue)
	return newValue
}

func (s *SecuritiesStatementQuery002V06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
