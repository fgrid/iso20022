package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01900101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.019.001.01 Document"`
	Message *UndertakingStatusReportV01 `xml:"UdrtkgStsRpt"`
}

func (d *Document01900101) AddMessage() *UndertakingStatusReportV01 {
	d.Message = new(UndertakingStatusReportV01)
	return d.Message
}

// The UndertakingStatusReport message is exchanged between parties that have an interest in the referenced undertaking transaction. It notifies the recipient of the status of the transaction, such as acceptance or rejection, withdrawal, or non-conformation. The sender may add additional information, as appropriate.
type UndertakingStatusReportV01 struct {

	// Details of the undertaking status report
	UndertakingStatusReportDetails *iso20022.UndertakingStatusAdvice1 `xml:"UdrtkgStsRptDtls"`

	// Digital signature of the report.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingStatusReportV01) AddUndertakingStatusReportDetails() *iso20022.UndertakingStatusAdvice1 {
	u.UndertakingStatusReportDetails = new(iso20022.UndertakingStatusAdvice1)
	return u.UndertakingStatusReportDetails
}

func (u *UndertakingStatusReportV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
