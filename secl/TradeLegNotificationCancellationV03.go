package secl

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:secl.002.001.03 Document"`
	Message *TradeLegNotificationCancellationV03 `xml:"TradLegNtfctnCxl"`
}

func (d *Document00200103) AddMessage() *TradeLegNotificationCancellationV03 {
	d.Message = new(TradeLegNotificationCancellationV03)
	return d.Message
}

// Scope
// The TradeLegNotificationCancellation message is sent by the central counterparty (CCP) to a clearing member to notify the cancellation of a TradeLegNotification message previously sent.
// 
// The message definition is intended for use with the ISO20022 Business Application Header.
// 
// Usage
// The previously sent message must be the Trade Leg Notification message.
type TradeLegNotificationCancellationV03 struct {

	// Provides the identification of the account owner, that is the clearing member (individual clearing member or general clearing member).
	ClearingMember *iso20022.PartyIdentification35Choice `xml:"ClrMmb"`

	// Identifies the clearing member account at the Central counterparty through which the trade must be cleared.
	ClearingAccount *iso20022.SecuritiesAccount18 `xml:"ClrAcct"`

	// An account opened by the central counterparty in the name of the clearing member or its settlement agent within the account structure, for settlement purposes (gives information about the clearing member/its settlement agent account at the central securities depository).
	DeliveryAccount *iso20022.SecuritiesAccount19 `xml:"DlvryAcct,omitempty"`

	// Provides details about the non clearing member identification and account.
	NonClearingMember *iso20022.PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides clearing details such as the non clearing member identification or the settlement netting (or not) eligibility code.
	ClearingDetails *iso20022.Clearing4 `xml:"ClrDtls,omitempty"`

	// Provides details about the trade leg such as the trade date, the settlement date or the trading currency.
	TradeLegDetails *iso20022.TradeLeg8 `xml:"TradLegDtls"`

	// Provides details about the settlement details of the trade leg such as the settlement currency or the place of settlement.
	SettlementDetails *iso20022.Settlement1 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (t *TradeLegNotificationCancellationV03) AddClearingMember() *iso20022.PartyIdentification35Choice {
	t.ClearingMember = new(iso20022.PartyIdentification35Choice)
	return t.ClearingMember
}

func (t *TradeLegNotificationCancellationV03) AddClearingAccount() *iso20022.SecuritiesAccount18 {
	t.ClearingAccount = new(iso20022.SecuritiesAccount18)
	return t.ClearingAccount
}

func (t *TradeLegNotificationCancellationV03) AddDeliveryAccount() *iso20022.SecuritiesAccount19 {
	t.DeliveryAccount = new(iso20022.SecuritiesAccount19)
	return t.DeliveryAccount
}

func (t *TradeLegNotificationCancellationV03) AddNonClearingMember() *iso20022.PartyIdentificationAndAccount31 {
	t.NonClearingMember = new(iso20022.PartyIdentificationAndAccount31)
	return t.NonClearingMember
}

func (t *TradeLegNotificationCancellationV03) AddClearingDetails() *iso20022.Clearing4 {
	t.ClearingDetails = new(iso20022.Clearing4)
	return t.ClearingDetails
}

func (t *TradeLegNotificationCancellationV03) AddTradeLegDetails() *iso20022.TradeLeg8 {
	t.TradeLegDetails = new(iso20022.TradeLeg8)
	return t.TradeLegDetails
}

func (t *TradeLegNotificationCancellationV03) AddSettlementDetails() *iso20022.Settlement1 {
	t.SettlementDetails = new(iso20022.Settlement1)
	return t.SettlementDetails
}

func (t *TradeLegNotificationCancellationV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

