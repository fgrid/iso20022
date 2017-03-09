package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.01 Document"`
	Message *AcceptorBatchTransferV01 `xml:"AccptrBtchTrf"`
}

func (d *Document01100101) AddMessage() *AcceptorBatchTransferV01 {
	d.Message = new(AcceptorBatchTransferV01)
	return d.Message
}

// Scope
// The AcceptorBatchTransfer message is sent by the card acceptor to the acquirer to capture a collection of previously completed card payment transactions.
// Usage
// The AcceptorBatchTransfer message embeds the information required for transferring to the acquirer the data needed to perform the financial settlement of these transactions (capture).
type AcceptorBatchTransferV01 struct {

	// Batch capture message management information.
	Header *iso20022.Header3 `xml:"Hdr"`

	// Information related to the set of transaction.
	DataSet []*iso20022.CardPaymentDataSet1 `xml:"DataSet"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType1 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV01) AddHeader() *iso20022.Header3 {
	a.Header = new(iso20022.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferV01) AddDataSet() *iso20022.CardPaymentDataSet1 {
	newValue := new(iso20022.CardPaymentDataSet1)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

func (a *AcceptorBatchTransferV01) AddSecurityTrailer() *iso20022.ContentInformationType1 {
	a.SecurityTrailer = new(iso20022.ContentInformationType1)
	return a.SecurityTrailer
}
