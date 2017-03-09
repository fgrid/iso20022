package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.01 Document"`
	Message *AcceptorBatchTransferResponseV01 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200101) AddMessage() *AcceptorBatchTransferResponseV01 {
	d.Message = new(AcceptorBatchTransferResponseV01)
	return d.Message
}

// Scope
// The AcceptorBatchTransferResponse message is sent by the acquirer to the card acceptor to acknowledge the proper reception of the AcceptorBatchTransfer.
// Usage
// The AcceptorBatchTransferResponse message is used by an acquirer to inform the card acceptor of the card payment transactions that could not be captured in the AcceptorBatchTransfer.
type AcceptorBatchTransferResponseV01 struct {

	// Capture advice response message management information.
	Header *iso20022.Header3 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	DataSet []*iso20022.CardPaymentDataSet2 `xml:"DataSet"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType1 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferResponseV01) AddHeader() *iso20022.Header3 {
	a.Header = new(iso20022.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV01) AddDataSet() *iso20022.CardPaymentDataSet2 {
	newValue := new(iso20022.CardPaymentDataSet2)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

func (a *AcceptorBatchTransferResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType1 {
	a.SecurityTrailer = new(iso20022.ContentInformationType1)
	return a.SecurityTrailer
}
