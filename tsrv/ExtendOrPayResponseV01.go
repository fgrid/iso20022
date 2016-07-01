package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.015.001.01 Document"`
	Message *ExtendOrPayResponseV01 `xml:"XtndOrPayRspn"`
}

func (d *Document01500101) AddMessage() *ExtendOrPayResponseV01 {
	d.Message = new(ExtendOrPayResponseV01)
	return d.Message
}

// The ExtendOrPayResponse message is sent by the party that requested issuance of the undertaking (applicant or obligor) to the party that issued the undertaking, in response to the issuer's request for the applicant's response to the beneficiary’s request to extend or pay.
type ExtendOrPayResponseV01 struct {

	// Details of the extend or pay response.
	ExtendOrPayResponseDetails *iso20022.ExtendOrPayQuery2 `xml:"XtndOrPayRspnDtls"`

	// Digital signature of the response.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

}


func (e *ExtendOrPayResponseV01) AddExtendOrPayResponseDetails() *iso20022.ExtendOrPayQuery2 {
	e.ExtendOrPayResponseDetails = new(iso20022.ExtendOrPayQuery2)
	return e.ExtendOrPayResponseDetails
}

func (e *ExtendOrPayResponseV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	e.DigitalSignature = new(iso20022.PartyAndSignature2)
	return e.DigitalSignature
}

