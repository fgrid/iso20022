package colr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.009.001.03 Document"`
	Message *MarginCallDisputeNotificationV03 `xml:"MrgnCallDsptNtfctn"`
}

func (d *Document00900103) AddMessage() *MarginCallDisputeNotificationV03 {
	d.Message = new(MarginCallDisputeNotificationV03)
	return d.Message
}

// Scope
// The MarginCallDisputeNotification message is sent by the collateral taker or its collateral manager to the collateral giver or its collateral manager to acknowledge the notification of the dispute (either full or partial dispute) of the MarginCallRequest. The message will detail the amount of the dispute and the reason.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// When there is a dispute by the collateral giver to the collateral taker a MarginCallDisputeNotification message is sent with the disputed amount (full or partial) stating the reason why the margin call is being disputed.
type MarginCallDisputeNotificationV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation3 `xml:"Oblgtn"`

	// Details of the dispute notification.
	DisputeNotification *iso20022.DisputeNotification1Choice `xml:"DsptNtfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MarginCallDisputeNotificationV03) SetTransactionIdentification(value string) {
	m.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (m *MarginCallDisputeNotificationV03) AddObligation() *iso20022.Obligation3 {
	m.Obligation = new(iso20022.Obligation3)
	return m.Obligation
}

func (m *MarginCallDisputeNotificationV03) AddDisputeNotification() *iso20022.DisputeNotification1Choice {
	m.DisputeNotification = new(iso20022.DisputeNotification1Choice)
	return m.DisputeNotification
}

func (m *MarginCallDisputeNotificationV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
