package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500101 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.001.01 Document"`
	Message *IntraPositionMovementConfirmationV01 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500101) AddMessage() *IntraPositionMovementConfirmationV01 {
	d.Message = new(IntraPositionMovementConfirmationV01)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementConfirmation to an account owner to confirm the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementConfirmationV01 struct {

	// Information that unambiguously identifies an IntraPositionMovementConfirmation message as known by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters3 `xml:"AddtlParams,omitempty"`

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
	IntraPositionDetails *iso20022.IntraPositionDetails2 `xml:"IntraPosDtls"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (i *IntraPositionMovementConfirmationV01) AddIdentification() *iso20022.DocumentIdentification11 {
	i.Identification = new(iso20022.DocumentIdentification11)
	return i.Identification
}

func (i *IntraPositionMovementConfirmationV01) AddAdditionalParameters() *iso20022.AdditionalParameters3 {
	i.AdditionalParameters = new(iso20022.AdditionalParameters3)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmationV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	i.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmationV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmationV01) AddSafekeepingPlace() *iso20022.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(iso20022.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmationV01) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification11 {
	i.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification11)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmationV01) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes4 {
	i.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes4)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmationV01) AddIntraPositionDetails() *iso20022.IntraPositionDetails2 {
	i.IntraPositionDetails = new(iso20022.IntraPositionDetails2)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmationV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	i.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return i.MessageOriginator
}

func (i *IntraPositionMovementConfirmationV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	i.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return i.MessageRecipient
}

func (i *IntraPositionMovementConfirmationV01) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
