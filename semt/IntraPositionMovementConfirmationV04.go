package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500104 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.001.04 Document"`
	Message *IntraPositionMovementConfirmationV04 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500104) AddMessage() *IntraPositionMovementConfirmationV04 {
	d.Message = new(IntraPositionMovementConfirmationV04)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementConfirmation to an account owner to confirm the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
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
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementConfirmationV04 struct {

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters10 `xml:"AddtlParams,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *iso20022.SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes36 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *iso20022.IntraPositionDetails27 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementConfirmationV04) AddAdditionalParameters() *iso20022.AdditionalParameters10 {
	i.AdditionalParameters = new(iso20022.AdditionalParameters10)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmationV04) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	i.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmationV04) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmationV04) AddSafekeepingPlace() *iso20022.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(iso20022.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmationV04) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmationV04) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes36 {
	i.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes36)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmationV04) AddIntraPositionDetails() *iso20022.IntraPositionDetails27 {
	i.IntraPositionDetails = new(iso20022.IntraPositionDetails27)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmationV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
