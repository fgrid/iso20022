package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.01 Document"`
	Message *ATMRejectV01 `xml:"ATMRjct"`
}

func (d *Document00500101) AddMessage() *ATMRejectV01 {
	d.Message = new(ATMRejectV01)
	return d.Message
}

// The ATMReject message is sent by any entity to reject a received message.
type ATMRejectV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header22 `xml:"Hdr"`

	// Information related to the reject of a message from an ATM or an ATM manager.
	ATMReject *iso20022.ATMReject1 `xml:"ATMRjct"`

}


func (a *ATMRejectV01) AddHeader() *iso20022.Header22 {
	a.Header = new(iso20022.Header22)
	return a.Header
}

func (a *ATMRejectV01) AddATMReject() *iso20022.ATMReject1 {
	a.ATMReject = new(iso20022.ATMReject1)
	return a.ATMReject
}

