package fxtr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800106 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.008.001.06 Document"`
	Message *ForeignExchangeTradeStatusNotificationV06 `xml:"FXTradStsNtfctn"`
}

func (d *Document00800106) AddMessage() *ForeignExchangeTradeStatusNotificationV06 {
	d.Message = new(ForeignExchangeTradeStatusNotificationV06)
	return d.Message
}

// Scope
// The ForeignExchangeTradeStatusNotification message is sent by a central system to the participant to notify the current status of a foreign exchange trade in the system.
// Usage
// This ForeignExchangeTradeStatusNotification message will be sent at specific times agreed upon by the central settlement system and a participant in a central settlement system.
type ForeignExchangeTradeStatusNotificationV06 struct {

	// Provides information on the status of a trade in a system.
	TradeData *iso20022.TradeData15 `xml:"TradData"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *iso20022.RegulatoryReporting6 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeStatusNotificationV06) AddTradeData() *iso20022.TradeData15 {
	f.TradeData = new(iso20022.TradeData15)
	return f.TradeData
}

func (f *ForeignExchangeTradeStatusNotificationV06) AddRegulatoryReporting() *iso20022.RegulatoryReporting6 {
	f.RegulatoryReporting = new(iso20022.RegulatoryReporting6)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeStatusNotificationV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
