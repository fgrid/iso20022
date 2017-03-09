package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.016.001.01 Document"`
	Message *DemandRefusalNotificationV01 `xml:"DmndRfslNtfctn"`
}

func (d *Document01600101) AddMessage() *DemandRefusalNotificationV01 {
	d.Message = new(DemandRefusalNotificationV01)
	return d.Message
}

// The DemandRefusalNotification message is sent to the beneficiary or presenter by the party obligated on the undertaking and to whom a demand for payment has been made, either directly or via one or more advising parties. It notifies the beneficiary or presenter that the demand has been refused.
type DemandRefusalNotificationV01 struct {

	// Details of the demand refusal notification.
	DemandRefusalNotificationDetails []*iso20022.DemandRefusal1 `xml:"DmndRfslNtfctnDtls,omitempty"`

	// Digital signature of the notification.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (d *DemandRefusalNotificationV01) AddDemandRefusalNotificationDetails() *iso20022.DemandRefusal1 {
	newValue := new(iso20022.DemandRefusal1)
	d.DemandRefusalNotificationDetails = append(d.DemandRefusalNotificationDetails, newValue)
	return newValue
}

func (d *DemandRefusalNotificationV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	d.DigitalSignature = new(iso20022.PartyAndSignature2)
	return d.DigitalSignature
}
