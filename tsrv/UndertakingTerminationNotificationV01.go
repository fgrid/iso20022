package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.012.001.01 Document"`
	Message *UndertakingTerminationNotificationV01 `xml:"UdrtkgTermntnNtfctn"`
}

func (d *Document01200101) AddMessage() *UndertakingTerminationNotificationV01 {
	d.Message = new(UndertakingTerminationNotificationV01)
	return d.Message
}

// The UndertakingTerminationNotification message is sent to the applicant by the party that issued the undertaking to give notification of the termination or cancelation of the referenced undertaking.
type UndertakingTerminationNotificationV01 struct {

	// Details of the termination notification.
	UndertakingTerminationNotificationDetails *iso20022.UndertakingTerminationNotice1 `xml:"UdrtkgTermntnNtfctnDtls"`

	// Digital signature of the notification.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingTerminationNotificationV01) AddUndertakingTerminationNotificationDetails() *iso20022.UndertakingTerminationNotice1 {
	u.UndertakingTerminationNotificationDetails = new(iso20022.UndertakingTerminationNotice1)
	return u.UndertakingTerminationNotificationDetails
}

func (u *UndertakingTerminationNotificationV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
