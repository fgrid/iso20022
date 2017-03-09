package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.001.01 Document"`
	Message *SecuritiesTransactionStatusQueryV01 `xml:"SctiesTxStsQry"`
}

func (d *Document02100101) AddMessage() *SecuritiesTransactionStatusQueryV01 {
	d.Message = new(SecuritiesTransactionStatusQueryV01)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesTransactionStatusQuery to an account servicer to request a status on a securities transaction.
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
type SecuritiesTransactionStatusQueryV01 struct {

	// Information that unambiguously identifies a SecuritiesTransactionStatusQuery message as know by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Description of the status advise requested.
	StatusAdviceRequested *iso20022.DocumentNumber2 `xml:"StsAdvcReqd"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesTransactionStatusQueryV01) AddIdentification() *iso20022.DocumentIdentification11 {
	s.Identification = new(iso20022.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesTransactionStatusQueryV01) AddStatusAdviceRequested() *iso20022.DocumentNumber2 {
	s.StatusAdviceRequested = new(iso20022.DocumentNumber2)
	return s.StatusAdviceRequested
}

func (s *SecuritiesTransactionStatusQueryV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	s.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionStatusQueryV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionStatusQueryV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	s.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesTransactionStatusQueryV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	s.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesTransactionStatusQueryV01) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
