package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.01 Document"`
	Message *IntraPositionMovementInstructionV01 `xml:"IntraPosMvmntInstr"`
}

func (d *Document01300101) AddMessage() *IntraPositionMovementInstructionV01 {
	d.Message = new(IntraPositionMovementInstructionV01)
	return d.Message
}

// Scope
// An account owner sends a IntraPositionMovementInstruction to an account servicer to instruct the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with its local agent (sub-custodian), or
// - an investment management institution which manage a fund account opened at a custodian, or
// - broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementInstructionV01 struct {

	// Information that unambiguously identifies an IntraPositionMovementTransaction and an IntraPositionMovementInstruction message as known by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.TransactionAndDocumentIdentification1 `xml:"Id"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *iso20022.Identification1 `xml:"CorpActnEvtId,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages *iso20022.Linkages2 `xml:"Lnkgs,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *iso20022.SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes4 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *iso20022.IntraPositionDetails1 `xml:"IntraPosDtls"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (i *IntraPositionMovementInstructionV01) AddIdentification() *iso20022.TransactionAndDocumentIdentification1 {
	i.Identification = new(iso20022.TransactionAndDocumentIdentification1)
	return i.Identification
}

func (i *IntraPositionMovementInstructionV01) AddCorporateActionEventIdentification() *iso20022.Identification1 {
	i.CorporateActionEventIdentification = new(iso20022.Identification1)
	return i.CorporateActionEventIdentification
}

func (i *IntraPositionMovementInstructionV01) AddLinkages() *iso20022.Linkages2 {
	i.Linkages = new(iso20022.Linkages2)
	return i.Linkages
}

func (i *IntraPositionMovementInstructionV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	i.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementInstructionV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementInstructionV01) AddSafekeepingPlace() *iso20022.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(iso20022.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementInstructionV01) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification11 {
	i.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification11)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementInstructionV01) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes4 {
	i.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes4)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementInstructionV01) AddIntraPositionDetails() *iso20022.IntraPositionDetails1 {
	i.IntraPositionDetails = new(iso20022.IntraPositionDetails1)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementInstructionV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	i.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return i.MessageOriginator
}

func (i *IntraPositionMovementInstructionV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	i.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return i.MessageRecipient
}

func (i *IntraPositionMovementInstructionV01) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
