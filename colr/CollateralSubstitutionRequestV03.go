package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:colr.010.001.03 Document"`
	Message *CollateralSubstitutionRequestV03 `xml:"CollSbstitnReq"`
}

func (d *Document01000103) AddMessage() *CollateralSubstitutionRequestV03 {
	d.Message = new(CollateralSubstitutionRequestV03)
	return d.Message
}

// Scope
// This CollateralSubstitutionRequest message is sent by either the collateral giver or its collateral manager to the collateral taker or its collateral manager. It is used to request a substitution of collateral by specifying the collateral to be returned and proposing the new type(s) of collateral to be delivered. Note: There are cases where the collateral taker can initiate the CollateralSubstitutionRequest message, for example in case of breach in the concentration limit.
// 
// The message definition is intended for use with the ISO20022 Business Application Header.
// 
// Usage
// The CollateralSubstitutionRequest message can be sent by either the collateral giver or collateral taker in order to substitute collateral that is already held by the other party.
type CollateralSubstitutionRequestV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement2 `xml:"Agrmt,omitempty"`

	// Provides details about the collateral that will be returned.
	CollateralSubstitutionReturn *iso20022.CollateralSubstitution2 `xml:"CollSbstitnRtr"`

	// Provides details about the collateral that will be delivered.
	CollateralSubstitutionDeliver *iso20022.CollateralSubstitution3 `xml:"CollSbstitnDlvr,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *CollateralSubstitutionRequestV03) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (c *CollateralSubstitutionRequestV03) AddObligation() *iso20022.Obligation3 {
	c.Obligation = new(iso20022.Obligation3)
	return c.Obligation
}

func (c *CollateralSubstitutionRequestV03) AddAgreement() *iso20022.Agreement2 {
	c.Agreement = new(iso20022.Agreement2)
	return c.Agreement
}

func (c *CollateralSubstitutionRequestV03) AddCollateralSubstitutionReturn() *iso20022.CollateralSubstitution2 {
	c.CollateralSubstitutionReturn = new(iso20022.CollateralSubstitution2)
	return c.CollateralSubstitutionReturn
}

func (c *CollateralSubstitutionRequestV03) AddCollateralSubstitutionDeliver() *iso20022.CollateralSubstitution3 {
	c.CollateralSubstitutionDeliver = new(iso20022.CollateralSubstitution3)
	return c.CollateralSubstitutionDeliver
}

func (c *CollateralSubstitutionRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

