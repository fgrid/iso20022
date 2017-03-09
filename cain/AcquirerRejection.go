package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name           `xml:"urn:iso:std:iso:20022:tech:xsd:cain.013.001.01 Document"`
	Message *AcquirerRejection `xml:"AcqrrRjctn"`
}

func (d *Document01300101) AddMessage() *AcquirerRejection {
	d.Message = new(AcquirerRejection)
	return d.Message
}

// The AcquirerRejection message is sent by any party, to reject an Acquirer to Issuer message.
type AcquirerRejection struct {

	// Information related to the protocol management.
	Header *iso20022.Header19 `xml:"Hdr"`

	// Information related to the reject.
	Reject *iso20022.AcceptorRejection4 `xml:"Rjct"`
}

func (a *AcquirerRejection) AddHeader() *iso20022.Header19 {
	a.Header = new(iso20022.Header19)
	return a.Header
}

func (a *AcquirerRejection) AddReject() *iso20022.AcceptorRejection4 {
	a.Reject = new(iso20022.AcceptorRejection4)
	return a.Reject
}
