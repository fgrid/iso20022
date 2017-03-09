package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700104 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.007.001.04 Document"`
	Message *CollateralProposalV04 `xml:"CollPrpsl"`
}

func (d *Document00700104) AddMessage() *CollateralProposalV04 {
	d.Message = new(CollateralProposalV04)
	return d.Message
}

// Scope
// The CollateralProposal message is sent by the collateral giver or its collateral manager to the collateral taker or its collateral manager, to propose the collateral to be delivered. This message is sent once the Margin Call is agreed or partially agreed and can be used for new collateral at the initiation of an exposure or for additional collateral for variation of an existing exposure. This message is used for both initial collateral proposal and subsequent counter proposals.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// This message is sent once the Margin Call is agreed or partially agreed and can be used for new collateral at the initiation of an exposure or for additional collateral for variation of an existing exposure. The collateral proposal can include securities, cash and other types of collateral.
type CollateralProposalV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement4 `xml:"Agrmt,omitempty"`

	// Indicates whether this is an initial or counter proposal.
	TypeAndDetails *iso20022.Proposal4 `xml:"TpAndDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralProposalV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (c *CollateralProposalV04) AddObligation() *iso20022.Obligation4 {
	c.Obligation = new(iso20022.Obligation4)
	return c.Obligation
}

func (c *CollateralProposalV04) AddAgreement() *iso20022.Agreement4 {
	c.Agreement = new(iso20022.Agreement4)
	return c.Agreement
}

func (c *CollateralProposalV04) AddTypeAndDetails() *iso20022.Proposal4 {
	c.TypeAndDetails = new(iso20022.Proposal4)
	return c.TypeAndDetails
}

func (c *CollateralProposalV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
