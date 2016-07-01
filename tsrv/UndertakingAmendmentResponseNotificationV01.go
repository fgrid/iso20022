package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.009.001.01 Document"`
	Message *UndertakingAmendmentResponseNotificationV01 `xml:"UdrtkgAmdmntRspnNtfctn"`
}

func (d *Document00900101) AddMessage() *UndertakingAmendmentResponseNotificationV01 {
	d.Message = new(UndertakingAmendmentResponseNotificationV01)
	return d.Message
}

// The UndertakingAmendmentResponseNotification message is sent by the advising party to the party that issued the undertaking, either directly or via one or more other parties, to notify the recipient of the acceptance or rejection by the beneficiary of the amendment. On receipt of this message or the UndertakingAmendmentResponse message, the issuer may also send the UndertakingAmendmentResponseNotification to the applicant.
type UndertakingAmendmentResponseNotificationV01 struct {

	// Details related to the proposed amendment response notification.
	UndertakingAmendmentResponseNotificationDetails *iso20022.Amendment9 `xml:"UdrtkgAmdmntRspnNtfctnDtls"`

	// Additional information reported by the beneficiary.
	AdditionalInformation *iso20022.Max2000Text `xml:"AddtlInf,omitempty"`

	// Digital signature of the response notification.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

}


func (u *UndertakingAmendmentResponseNotificationV01) AddUndertakingAmendmentResponseNotificationDetails() *iso20022.Amendment9 {
	u.UndertakingAmendmentResponseNotificationDetails = new(iso20022.Amendment9)
	return u.UndertakingAmendmentResponseNotificationDetails
}

func (u *UndertakingAmendmentResponseNotificationV01) SetAdditionalInformation(value string) {
	u.AdditionalInformation = (*iso20022.Max2000Text)(&value)
}

func (u *UndertakingAmendmentResponseNotificationV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}

