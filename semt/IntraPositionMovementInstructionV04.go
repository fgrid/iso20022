package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.04 Document"`
	Message *IntraPositionMovementInstructionV04 `xml:"IntraPosMvmntInstr"`
}

func (d *Document01300104) AddMessage() *IntraPositionMovementInstructionV04 {
	d.Message = new(IntraPositionMovementInstructionV04)
	return d.Message
}

// Scope
// An account owner sends a IntraPositionMovementInstruction to an account servicer to instruct the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with its local agent (sub-custodian), or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
type IntraPositionMovementInstructionV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *iso20022.Identification14 `xml:"CorpActnEvtId,omitempty"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages36 `xml:"Lnkgs,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *iso20022.SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes63 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *iso20022.IntraPositionDetails33 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (i *IntraPositionMovementInstructionV04) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (i *IntraPositionMovementInstructionV04) AddCorporateActionEventIdentification() *iso20022.Identification14 {
	i.CorporateActionEventIdentification = new(iso20022.Identification14)
	return i.CorporateActionEventIdentification
}

func (i *IntraPositionMovementInstructionV04) AddNumberCounts() *iso20022.NumberCount1Choice {
	i.NumberCounts = new(iso20022.NumberCount1Choice)
	return i.NumberCounts
}

func (i *IntraPositionMovementInstructionV04) AddLinkages() *iso20022.Linkages36 {
	newValue := new (iso20022.Linkages36)
	i.Linkages = append(i.Linkages, newValue)
	return newValue
}

func (i *IntraPositionMovementInstructionV04) AddAccountOwner() *iso20022.PartyIdentification92Choice {
	i.AccountOwner = new(iso20022.PartyIdentification92Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementInstructionV04) AddSafekeepingAccount() *iso20022.SecuritiesAccount24 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount24)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementInstructionV04) AddSafekeepingPlace() *iso20022.SafekeepingPlaceFormat10Choice {
	i.SafekeepingPlace = new(iso20022.SafekeepingPlaceFormat10Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementInstructionV04) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	i.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementInstructionV04) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes63 {
	i.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes63)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementInstructionV04) AddIntraPositionDetails() *iso20022.IntraPositionDetails33 {
	i.IntraPositionDetails = new(iso20022.IntraPositionDetails33)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementInstructionV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}

