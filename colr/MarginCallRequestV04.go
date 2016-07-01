package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:colr.003.001.04 Document"`
	Message *MarginCallRequestV04 `xml:"MrgnCallReq"`
}

func (d *Document00300104) AddMessage() *MarginCallRequestV04 {
	d.Message = new(MarginCallRequestV04)
	return d.Message
}

// Scope
// The MarginCallRequest message is sent by the collateral taker or its collateral manager to the collateral giver or its collateral manager
// This message is used to request new collateral at the initiation of an exposure or request additional collateral for an existing exposure. It can also be used to recall collateral upon the collateral giver or its collateral manager's request.
// 
// The message definition is intended for use with the ISO20022 Business Application Header.
// 
// Usage
// When sent by the collateral taker the MarginCallRequest message is used to:
// - request new collateral at the initiation of an exposure
// - request additional collateral
// When sent by the collateral giver the MarginCallRequest message is used to:
// - request the return of collateral
type MarginCallRequestV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement4 `xml:"Agrmt,omitempty"`

	// Summation of the call amounts per margin type. It is provided for the purposes of carrying forward for future messages that are used to compare the margin call results to determine whether a call is agreed or full/partially disputed.
	MarginCallResult *iso20022.MarginCallResult3 `xml:"MrgnCallRslt"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailsDueToA *iso20022.MarginCall1 `xml:"MrgnDtlsDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailsDueToB *iso20022.MarginCall1 `xml:"MrgnDtlsDueToB,omitempty"`

	// Amount of expected margin that will be either delivered to party A by party B or recalled to party A from party B.
	RequirementDetailsDueToA *iso20022.MarginRequirement1Choice `xml:"RqrmntDtlsDueToA,omitempty"`

	// Amount of expected margin that will be either delivered to party B by party A or recalled to party B from party A.
	RequirementDetailsDueToB *iso20022.MarginRequirement1Choice `xml:"RqrmntDtlsDueToB,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party A by party B or recalled to party A from party B.
	ExpectedCollateralDueToA *iso20022.ExpectedCollateral2Choice `xml:"XpctdCollDueToA,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party B by party A or recalled to party B from party A.
	ExpectedCollateralDueToB *iso20022.ExpectedCollateral2Choice `xml:"XpctdCollDueToB,omitempty"`

	// Allows the reporting of the margin requirements for multiple accounts and report a single margin call amount made up of the aggregate of all the individual requirement amounts.
	MarginCallDetails []*iso20022.MarginCall2 `xml:"MrgnCallDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MarginCallRequestV04) SetTransactionIdentification(value string) {
	m.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (m *MarginCallRequestV04) AddObligation() *iso20022.Obligation4 {
	m.Obligation = new(iso20022.Obligation4)
	return m.Obligation
}

func (m *MarginCallRequestV04) AddAgreement() *iso20022.Agreement4 {
	m.Agreement = new(iso20022.Agreement4)
	return m.Agreement
}

func (m *MarginCallRequestV04) AddMarginCallResult() *iso20022.MarginCallResult3 {
	m.MarginCallResult = new(iso20022.MarginCallResult3)
	return m.MarginCallResult
}

func (m *MarginCallRequestV04) AddMarginDetailsDueToA() *iso20022.MarginCall1 {
	m.MarginDetailsDueToA = new(iso20022.MarginCall1)
	return m.MarginDetailsDueToA
}

func (m *MarginCallRequestV04) AddMarginDetailsDueToB() *iso20022.MarginCall1 {
	m.MarginDetailsDueToB = new(iso20022.MarginCall1)
	return m.MarginDetailsDueToB
}

func (m *MarginCallRequestV04) AddRequirementDetailsDueToA() *iso20022.MarginRequirement1Choice {
	m.RequirementDetailsDueToA = new(iso20022.MarginRequirement1Choice)
	return m.RequirementDetailsDueToA
}

func (m *MarginCallRequestV04) AddRequirementDetailsDueToB() *iso20022.MarginRequirement1Choice {
	m.RequirementDetailsDueToB = new(iso20022.MarginRequirement1Choice)
	return m.RequirementDetailsDueToB
}

func (m *MarginCallRequestV04) AddExpectedCollateralDueToA() *iso20022.ExpectedCollateral2Choice {
	m.ExpectedCollateralDueToA = new(iso20022.ExpectedCollateral2Choice)
	return m.ExpectedCollateralDueToA
}

func (m *MarginCallRequestV04) AddExpectedCollateralDueToB() *iso20022.ExpectedCollateral2Choice {
	m.ExpectedCollateralDueToB = new(iso20022.ExpectedCollateral2Choice)
	return m.ExpectedCollateralDueToB
}

func (m *MarginCallRequestV04) AddMarginCallDetails() *iso20022.MarginCall2 {
	newValue := new (iso20022.MarginCall2)
	m.MarginCallDetails = append(m.MarginCallDetails, newValue)
	return newValue
}

func (m *MarginCallRequestV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

