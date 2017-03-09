package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.008.001.01 Document"`
	Message *UndertakingAmendmentResponseV01 `xml:"UdrtkgAmdmntRspn"`
}

func (d *Document00800101) AddMessage() *UndertakingAmendmentResponseV01 {
	d.Message = new(UndertakingAmendmentResponseV01)
	return d.Message
}

// The UndertakingAmendmentResponse message is sent by the beneficiary to the party that issued the undertaking, either directly or via one or more advising parties, to indicate acceptance or rejection by the beneficiary of the amendment.
type UndertakingAmendmentResponseV01 struct {

	// Details related to the proposed amendment response.
	UndertakingAmendmentResponseDetails *iso20022.Amendment7 `xml:"UdrtkgAmdmntRspnDtls"`

	// Digital signature of the response.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentResponseV01) AddUndertakingAmendmentResponseDetails() *iso20022.Amendment7 {
	u.UndertakingAmendmentResponseDetails = new(iso20022.Amendment7)
	return u.UndertakingAmendmentResponseDetails
}

func (u *UndertakingAmendmentResponseV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
