package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.014.001.01 Document"`
	Message *ExtendOrPayRequestV01 `xml:"XtndOrPayReq"`
}

func (d *Document01400101) AddMessage() *ExtendOrPayRequestV01 {
	d.Message = new(ExtendOrPayRequestV01)
	return d.Message
}

// The ExtendOrPayRequest message is sent by the party that issued the undertaking to the party that requested issuance of the undertaking (applicant or obligor), to request the applicant's response to a beneficiary request to extend or pay.
type ExtendOrPayRequestV01 struct {

	// Details of the extend or pay request.
	ExtendOrPayRequestDetails *iso20022.ExtendOrPayQuery1 `xml:"XtndOrPayReqDtls"`

	// Digital signature of the request.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

}


func (e *ExtendOrPayRequestV01) AddExtendOrPayRequestDetails() *iso20022.ExtendOrPayQuery1 {
	e.ExtendOrPayRequestDetails = new(iso20022.ExtendOrPayQuery1)
	return e.ExtendOrPayRequestDetails
}

func (e *ExtendOrPayRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	e.DigitalSignature = new(iso20022.PartyAndSignature2)
	return e.DigitalSignature
}

