package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.021.001.02 Document"`
	Message *SecuritiesStatementQueryV02 `xml:"SctiesStmtQry"`
}

func (d *Document02100102) AddMessage() *SecuritiesStatementQueryV02 {
	d.Message = new(SecuritiesStatementQueryV02)
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
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesStatementQueryV02 struct {

	// Description of the statement requested.
	StatementRequested *iso20022.DocumentNumber1 `xml:"StmtReqd"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement16 `xml:"StmtGnlDtls,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Additional specific query criteria.
	AdditionalQueryParameters []*iso20022.AdditionalQueryParameters3 `xml:"AddtlQryParams,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesStatementQueryV02) AddStatementRequested() *iso20022.DocumentNumber1 {
	s.StatementRequested = new(iso20022.DocumentNumber1)
	return s.StatementRequested
}

func (s *SecuritiesStatementQueryV02) AddStatementGeneralDetails() *iso20022.Statement16 {
	s.StatementGeneralDetails = new(iso20022.Statement16)
	return s.StatementGeneralDetails
}

func (s *SecuritiesStatementQueryV02) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesStatementQueryV02) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatementQueryV02) AddAdditionalQueryParameters() *iso20022.AdditionalQueryParameters3 {
	newValue := new (iso20022.AdditionalQueryParameters3)
	s.AdditionalQueryParameters = append(s.AdditionalQueryParameters, newValue)
	return newValue
}

func (s *SecuritiesStatementQueryV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

