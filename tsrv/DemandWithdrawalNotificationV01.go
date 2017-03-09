package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.017.001.01 Document"`
	Message *DemandWithdrawalNotificationV01 `xml:"DmndWdrwlNtfctn"`
}

func (d *Document01700101) AddMessage() *DemandWithdrawalNotificationV01 {
	d.Message = new(DemandWithdrawalNotificationV01)
	return d.Message
}

// The DemandWithdrawalNotification message is sent by the beneficiary to the party that issued the undertaking, either directly or via a presenting or nominated party, to inform the issuer or nominated party that it has elected to withdraw its demand under the demand guarantee or standby letter of credit.
type DemandWithdrawalNotificationV01 struct {

	// Details of the demand withdrawal notification.
	DemandWithdrawalNotificationDetails *iso20022.UndertakingDemandWithdrawal1 `xml:"DmndWdrwlNtfctnDtls"`

	// Digital signature of the notification.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (d *DemandWithdrawalNotificationV01) AddDemandWithdrawalNotificationDetails() *iso20022.UndertakingDemandWithdrawal1 {
	d.DemandWithdrawalNotificationDetails = new(iso20022.UndertakingDemandWithdrawal1)
	return d.DemandWithdrawalNotificationDetails
}

func (d *DemandWithdrawalNotificationV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	d.DigitalSignature = new(iso20022.PartyAndSignature2)
	return d.DigitalSignature
}
