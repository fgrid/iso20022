package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02400101 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.01 Document"`
	Message *PaymentRegulatoryInformationNotificationV01 `xml:"PmtRgltryInfNtfctn"`
}

func (d *Document02400101) AddMessage() *PaymentRegulatoryInformationNotificationV01 {
	d.Message = new(PaymentRegulatoryInformationNotificationV01)
	return d.Message
}

// The PaymentRegulatoryInformationNotification message is sent by the reporting party to the registration agent to provide details on the transaction details, when a payment has to be recorded against the registered currency control contract.
//
// In some cases, the registration agent may also sent this message to the reporting party.
type PaymentRegulatoryInformationNotificationV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *iso20022.CurrencyControlHeader3 `xml:"GrpHdr"`

	// Notification of information related to a regulatory reporting on a payment.
	TransactionNotification []*iso20022.RegulatoryReportingNotification1 `xml:"TxNtfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentRegulatoryInformationNotificationV01) AddGroupHeader() *iso20022.CurrencyControlHeader3 {
	p.GroupHeader = new(iso20022.CurrencyControlHeader3)
	return p.GroupHeader
}

func (p *PaymentRegulatoryInformationNotificationV01) AddTransactionNotification() *iso20022.RegulatoryReportingNotification1 {
	newValue := new(iso20022.RegulatoryReportingNotification1)
	p.TransactionNotification = append(p.TransactionNotification, newValue)
	return newValue
}

func (p *PaymentRegulatoryInformationNotificationV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
