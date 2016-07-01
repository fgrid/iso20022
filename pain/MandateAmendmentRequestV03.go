package pain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.03 Document"`
	Message *MandateAmendmentRequestV03 `xml:"MndtAmdmntReq"`
}

func (d *Document01000103) AddMessage() *MandateAmendmentRequestV03 {
	d.Message = new(MandateAmendmentRequestV03)
	return d.Message
}

// Scope
// The MandateAmendmentRequest message is sent by the initiator of the request to his agent and/or counterparty. The initiator can both be the debtor or the creditor (or where appropriate the debtor agent).
// The MandateAmendmentRequest message is forwarded by the agent of the initiator to the agent of the counterparty.
// A MandateAmendmentRequest message is used to request the amendment of specific information in an existing mandate. The MandateAmendmentRequest message must reflect the new data of the element(s) to be amended and at a minimum a unique reference to the existing mandate. If accepted, this MandateAmendmentRequest message together with the MandateAcceptanceReport message confirming the acceptance will be considered as a valid amendment on an existing mandate, agreed upon by all parties. The amended mandate will from then on be considered the valid mandate.
// Usage
// The MandateAmendmentRequest message can contain one or more request(s) to amend a specific mandate.
// The messages can be exchanged between creditor and creditor agent or debtor and debtor agent and between creditor agent and debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor or debtor.
// The MandateAmendmentRequest message can be used in domestic and cross-border scenarios.
// If all elements in the existing Mandate need to be amended or the underlying contract is different, then the MandateAmendmentRequest message should not be used. The existing Mandate has to be cancelled and a new Mandate has to be initiated.
type MandateAmendmentRequestV03 struct {

	// Set of characteristics to identify the message and parties playing a role in the amendment of the mandate, but which are not part of the mandate.
	GroupHeader *iso20022.GroupHeader47 `xml:"GrpHdr"`

	// Set of elements used to provide details on the amendment request.
	UnderlyingAmendmentDetails []*iso20022.MandateAmendment3 `xml:"UndrlygAmdmntDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MandateAmendmentRequestV03) AddGroupHeader() *iso20022.GroupHeader47 {
	m.GroupHeader = new(iso20022.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateAmendmentRequestV03) AddUnderlyingAmendmentDetails() *iso20022.MandateAmendment3 {
	newValue := new (iso20022.MandateAmendment3)
	m.UnderlyingAmendmentDetails = append(m.UnderlyingAmendmentDetails, newValue)
	return newValue
}

func (m *MandateAmendmentRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

