package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.021.001.01 Document"`
	Message *SecuritiesStatementQueryV01 `xml:"SctiesStmtQry"`
}

func (d *Document02100101) AddMessage() *SecuritiesStatementQueryV01 {
	d.Message = new(SecuritiesStatementQueryV01)
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
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesStatementQueryV01 struct {

	// Information that unambiguously identifies a SecuritiesStatementQuery message as know by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Description of the statement requested.
	StatementRequested *iso20022.DocumentNumber1 `xml:"StmtReqd"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement16 `xml:"StmtGnlDtls,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Additional specific query criteria.
	AdditionalQueryParameters []*iso20022.AdditionalQueryParameters1 `xml:"AddtlQryParams,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`

}


func (s *SecuritiesStatementQueryV01) AddIdentification() *iso20022.DocumentIdentification11 {
	s.Identification = new(iso20022.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesStatementQueryV01) AddStatementRequested() *iso20022.DocumentNumber1 {
	s.StatementRequested = new(iso20022.DocumentNumber1)
	return s.StatementRequested
}

func (s *SecuritiesStatementQueryV01) AddStatementGeneralDetails() *iso20022.Statement16 {
	s.StatementGeneralDetails = new(iso20022.Statement16)
	return s.StatementGeneralDetails
}

func (s *SecuritiesStatementQueryV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	s.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesStatementQueryV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatementQueryV01) AddAdditionalQueryParameters() *iso20022.AdditionalQueryParameters1 {
	newValue := new (iso20022.AdditionalQueryParameters1)
	s.AdditionalQueryParameters = append(s.AdditionalQueryParameters, newValue)
	return newValue
}

func (s *SecuritiesStatementQueryV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	s.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesStatementQueryV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	s.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesStatementQueryV01) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

