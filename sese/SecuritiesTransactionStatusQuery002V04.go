package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100204 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.021.002.04 Document"`
	Message *SecuritiesTransactionStatusQuery002V04 `xml:"SctiesTxStsQry"`
}

func (d *Document02100204) AddMessage() *SecuritiesTransactionStatusQuery002V04 {
	d.Message = new(SecuritiesTransactionStatusQuery002V04)
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
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
type SecuritiesTransactionStatusQuery002V04 struct {

	// Description of the status advise requested.
	StatusAdviceRequested *iso20022.DocumentNumber15 `xml:"StsAdvcReqd"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTransactionStatusQuery002V04) AddStatusAdviceRequested() *iso20022.DocumentNumber15 {
	s.StatusAdviceRequested = new(iso20022.DocumentNumber15)
	return s.StatusAdviceRequested
}

func (s *SecuritiesTransactionStatusQuery002V04) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesTransactionStatusQuery002V04) AddSafekeepingAccount() *iso20022.SecuritiesAccount27 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount27)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionStatusQuery002V04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
