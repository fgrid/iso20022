package pain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400106 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.06 Document"`
	Message *CreditorPaymentActivationRequestStatusReportV06 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (d *Document01400106) AddMessage() *CreditorPaymentActivationRequestStatusReportV06 {
	d.Message = new(CreditorPaymentActivationRequestStatusReportV06)
	return d.Message
}

// The CreditorPaymentActivationRequestStatusReport message is sent by a party to the next party in the creditor payment activation request chain. It is used to inform the latter about the positive or negative status of a creditor payment activation request (either single or file).
type CreditorPaymentActivationRequestStatusReportV06 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader46 `xml:"GrpHdr"`

	// Original group information concerning the group of transactions, to which the status report message refers to.
	OriginalGroupInformationAndStatus *iso20022.OriginalGroupInformation28 `xml:"OrgnlGrpInfAndSts"`

	// Information concerning the original payment information, to which the status report message refers.
	OriginalPaymentInformationAndStatus []*iso20022.OriginalPaymentInstruction24 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestStatusReportV06) AddGroupHeader() *iso20022.GroupHeader46 {
	c.GroupHeader = new(iso20022.GroupHeader46)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestStatusReportV06) AddOriginalGroupInformationAndStatus() *iso20022.OriginalGroupInformation28 {
	c.OriginalGroupInformationAndStatus = new(iso20022.OriginalGroupInformation28)
	return c.OriginalGroupInformationAndStatus
}

func (c *CreditorPaymentActivationRequestStatusReportV06) AddOriginalPaymentInformationAndStatus() *iso20022.OriginalPaymentInstruction24 {
	newValue := new(iso20022.OriginalPaymentInstruction24)
	c.OriginalPaymentInformationAndStatus = append(c.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestStatusReportV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
