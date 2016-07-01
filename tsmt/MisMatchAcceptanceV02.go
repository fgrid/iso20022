package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02000102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.020.001.02 Document"`
	Message *MisMatchAcceptanceV02 `xml:"MisMtchAccptnc"`
}

func (d *Document02000102) AddMessage() *MisMatchAcceptanceV02 {
	d.Message = new(MisMatchAcceptanceV02)
	return d.Message
}

// Scope
// The MisMatchAcceptance message is sent by the party requested to accept or reject data set mis-matches to the matching application.
// This message is used to accept mis-matches between data sets and the related baseline.
// Usage
// The MisMatchAcceptance message can be sent by the party requested to accept or reject data set mis-matches to inform that it accepts the data sets.
// The message can be sent in response to a DataSetMatchReport message conveying mis-matches.
// The information about the acceptance of the mis-matched data sets will be forwarded by the matching application to the submitter of the data sets by a MisMatchAcceptanceNotification message.
// The rejection of mis-matched data sets can be achieved by sending a MisMatchRejection message.
type MisMatchAcceptanceV02 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *iso20022.MessageIdentification1 `xml:"AccptncId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the report that contained the difference. 
	DataSetMatchReportReference *iso20022.MessageIdentification1 `xml:"DataSetMtchRptRef"`

}


func (m *MisMatchAcceptanceV02) AddAcceptanceIdentification() *iso20022.MessageIdentification1 {
	m.AcceptanceIdentification = new(iso20022.MessageIdentification1)
	return m.AcceptanceIdentification
}

func (m *MisMatchAcceptanceV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	m.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return m.TransactionIdentification
}

func (m *MisMatchAcceptanceV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	m.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return m.SubmitterTransactionReference
}

func (m *MisMatchAcceptanceV02) AddDataSetMatchReportReference() *iso20022.MessageIdentification1 {
	m.DataSetMatchReportReference = new(iso20022.MessageIdentification1)
	return m.DataSetMatchReportReference
}

