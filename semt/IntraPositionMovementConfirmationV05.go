package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500105 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.001.05 Document"`
	Message *IntraPositionMovementConfirmationV05 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500105) AddMessage() *IntraPositionMovementConfirmationV05 {
	d.Message = new(IntraPositionMovementConfirmationV05)
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
type IntraPositionMovementConfirmationV05 struct {

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters21 `xml:"AddtlParams,omitempty"`

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
	IntraPositionDetails *iso20022.IntraPositionDetails34 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementConfirmationV05) AddAdditionalParameters() *iso20022.AdditionalParameters21 {
	i.AdditionalParameters = new(iso20022.AdditionalParameters21)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmationV05) AddAccountOwner() *iso20022.PartyIdentification92Choice {
	i.AccountOwner = new(iso20022.PartyIdentification92Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmationV05) AddSafekeepingAccount() *iso20022.SecuritiesAccount24 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount24)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmationV05) AddSafekeepingPlace() *iso20022.SafekeepingPlaceFormat10Choice {
	i.SafekeepingPlace = new(iso20022.SafekeepingPlaceFormat10Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmationV05) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	i.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmationV05) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes63 {
	i.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes63)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmationV05) AddIntraPositionDetails() *iso20022.IntraPositionDetails34 {
	i.IntraPositionDetails = new(iso20022.IntraPositionDetails34)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
