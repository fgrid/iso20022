package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300104 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.013.001.04 Document"`
	Message *InterestPaymentRequestV04 `xml:"IntrstPmtReq"`
}

func (d *Document01300104) AddMessage() *InterestPaymentRequestV04 {
	d.Message = new(InterestPaymentRequestV04)
	return d.Message
}

// Scope
// The InterestPaymentRequest message is sent by either;
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager
// It is used to request payment or advise the amount due for interest calculated for a specified period. The interest is based on the amount of collateral that has been held during the calculation period. It is possible to send these messages on a bi-lateral basis for matching.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The InterestPaymentRequest message is used to advise the interest amount calculated for a specific period or request payment for an interest amount for a specific period.
type InterestPaymentRequestV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement4 `xml:"Agrmt"`

	// Provides details on the interest amount due to party A.
	InterestDueToA *iso20022.InterestAmount1 `xml:"IntrstDueToA,omitempty"`

	// Provides details on the interest amount due to party B.
	InterestDueToB *iso20022.InterestAmount1 `xml:"IntrstDueToB,omitempty"`

	// Provides the net interest amount due to A or due to B in case of collateral held and posted during a period.
	NetAmountDetails *iso20022.InterestResult1 `xml:"NetAmtDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InterestPaymentRequestV04) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InterestPaymentRequestV04) AddObligation() *iso20022.Obligation4 {
	i.Obligation = new(iso20022.Obligation4)
	return i.Obligation
}

func (i *InterestPaymentRequestV04) AddAgreement() *iso20022.Agreement4 {
	i.Agreement = new(iso20022.Agreement4)
	return i.Agreement
}

func (i *InterestPaymentRequestV04) AddInterestDueToA() *iso20022.InterestAmount1 {
	i.InterestDueToA = new(iso20022.InterestAmount1)
	return i.InterestDueToA
}

func (i *InterestPaymentRequestV04) AddInterestDueToB() *iso20022.InterestAmount1 {
	i.InterestDueToB = new(iso20022.InterestAmount1)
	return i.InterestDueToB
}

func (i *InterestPaymentRequestV04) AddNetAmountDetails() *iso20022.InterestResult1 {
	i.NetAmountDetails = new(iso20022.InterestResult1)
	return i.NetAmountDetails
}

func (i *InterestPaymentRequestV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
