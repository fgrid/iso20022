package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:colr.003.001.03 Document"`
	Message *MarginCallRequestV03 `xml:"MrgnCallReq"`
}

func (d *Document00300103) AddMessage() *MarginCallRequestV03 {
	d.Message = new(MarginCallRequestV03)
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
type MarginCallRequestV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement2 `xml:"Agrmt,omitempty"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailsDueToA *iso20022.MarginCall1 `xml:"MrgnDtlsDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailsDueToB *iso20022.MarginCall1 `xml:"MrgnDtlsDueToB,omitempty"`

	// Amount of expected margin that will be either delivered to party A by party B or recalled to party A from party B.
	RequirementDetailsDueToA *iso20022.MarginRequirement1Choice `xml:"RqrmntDtlsDueToA,omitempty"`

	// Amount of expected margin that will be either delivered to party B by party A or recalled to party B from party A.
	RequirementDetailsDueToB *iso20022.MarginRequirement1Choice `xml:"RqrmntDtlsDueToB,omitempty"`

	// Summation of the call amounts per margin type. It is provided for the purposes of carrying forward for future messages that are used to compare the margin call results to determine whether a call is agreed or full/partially disputed.
	MarginCallResult *iso20022.MarginCallResult3 `xml:"MrgnCallRslt"`

	// Provides details about the type of collateral that will be either delivered to party B by party A or recalled to party B from party A.
	ExpectedCollateralDueToB *iso20022.ExpectedCollateral1Choice `xml:"XpctdCollDueToB,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party A by party B or recalled to party A from party B.
	ExpectedCollateralDueToA *iso20022.ExpectedCollateral1Choice `xml:"XpctdCollDueToA,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MarginCallRequestV03) SetTransactionIdentification(value string) {
	m.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (m *MarginCallRequestV03) AddObligation() *iso20022.Obligation3 {
	m.Obligation = new(iso20022.Obligation3)
	return m.Obligation
}

func (m *MarginCallRequestV03) AddAgreement() *iso20022.Agreement2 {
	m.Agreement = new(iso20022.Agreement2)
	return m.Agreement
}

func (m *MarginCallRequestV03) AddMarginDetailsDueToA() *iso20022.MarginCall1 {
	m.MarginDetailsDueToA = new(iso20022.MarginCall1)
	return m.MarginDetailsDueToA
}

func (m *MarginCallRequestV03) AddMarginDetailsDueToB() *iso20022.MarginCall1 {
	m.MarginDetailsDueToB = new(iso20022.MarginCall1)
	return m.MarginDetailsDueToB
}

func (m *MarginCallRequestV03) AddRequirementDetailsDueToA() *iso20022.MarginRequirement1Choice {
	m.RequirementDetailsDueToA = new(iso20022.MarginRequirement1Choice)
	return m.RequirementDetailsDueToA
}

func (m *MarginCallRequestV03) AddRequirementDetailsDueToB() *iso20022.MarginRequirement1Choice {
	m.RequirementDetailsDueToB = new(iso20022.MarginRequirement1Choice)
	return m.RequirementDetailsDueToB
}

func (m *MarginCallRequestV03) AddMarginCallResult() *iso20022.MarginCallResult3 {
	m.MarginCallResult = new(iso20022.MarginCallResult3)
	return m.MarginCallResult
}

func (m *MarginCallRequestV03) AddExpectedCollateralDueToB() *iso20022.ExpectedCollateral1Choice {
	m.ExpectedCollateralDueToB = new(iso20022.ExpectedCollateral1Choice)
	return m.ExpectedCollateralDueToB
}

func (m *MarginCallRequestV03) AddExpectedCollateralDueToA() *iso20022.ExpectedCollateral1Choice {
	m.ExpectedCollateralDueToA = new(iso20022.ExpectedCollateral1Choice)
	return m.ExpectedCollateralDueToA
}

func (m *MarginCallRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

