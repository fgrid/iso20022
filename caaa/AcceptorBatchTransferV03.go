package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.03 Document"`
	Message *AcceptorBatchTransferV03 `xml:"AccptrBtchTrf"`
}

func (d *Document01100103) AddMessage() *AcceptorBatchTransferV03 {
	d.Message = new(AcceptorBatchTransferV03)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV03 struct {

	// Batch capture message management information.
	Header *iso20022.Header3 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *iso20022.CardPaymentBatchTransfer2 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType9 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV03) AddHeader() *iso20022.Header3 {
	a.Header = new(iso20022.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferV03) AddBatchTransfer() *iso20022.CardPaymentBatchTransfer2 {
	a.BatchTransfer = new(iso20022.CardPaymentBatchTransfer2)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV03) AddSecurityTrailer() *iso20022.ContentInformationType9 {
	a.SecurityTrailer = new(iso20022.ContentInformationType9)
	return a.SecurityTrailer
}
