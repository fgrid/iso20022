package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.004.001.01 Document"`
	Message *UndertakingAmendmentRequestV01 `xml:"UdrtkgAmdmntReq"`
}

func (d *Document00400101) AddMessage() *UndertakingAmendmentRequestV01 {
	d.Message = new(UndertakingAmendmentRequestV01)
	return d.Message
}

// The UndertakingAmendmentRequest message is sent by the party that requested issuance of the undertaking (applicant or obligor) to the party that issued the undertaking to request issuance of a proposed amendment to the undertaking. The undertaking could be a demand guarantee, standby letter of credit, counter-undertaking (counter-guarantee or counter-standby), or suretyship undertaking. The message provides details on proposed changes to the undertaking, for example, to the expiry date, amount, and/or terms and conditions. It may also be used to request termination or cancellation of the undertaking.
type UndertakingAmendmentRequestV01 struct {

	// Details related to the request for an amendment of an undertaking.
	UndertakingAmendmentRequestDetails *iso20022.Amendment3 `xml:"UdrtkgAmdmntReqDtls"`

	// Instructions specific to the bank receiving the message.
	InstructionsToBank []*iso20022.Max2000Text `xml:"InstrsToBk,omitempty"`

	// Digital signature of the undertaking amendment request.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentRequestV01) AddUndertakingAmendmentRequestDetails() *iso20022.Amendment3 {
	u.UndertakingAmendmentRequestDetails = new(iso20022.Amendment3)
	return u.UndertakingAmendmentRequestDetails
}

func (u *UndertakingAmendmentRequestV01) AddInstructionsToBank(value string) {
	u.InstructionsToBank = append(u.InstructionsToBank, (*iso20022.Max2000Text)(&value))
}

func (u *UndertakingAmendmentRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
