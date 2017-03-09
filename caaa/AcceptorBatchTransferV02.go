package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.02 Document"`
	Message *AcceptorBatchTransferV02 `xml:"AccptrBtchTrf"`
}

func (d *Document01100102) AddMessage() *AcceptorBatchTransferV02 {
	d.Message = new(AcceptorBatchTransferV02)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV02 struct {

	// Batch capture message management information.
	Header *iso20022.Header3 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *iso20022.CardPaymentBatchTransfer1 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType4 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV02) AddHeader() *iso20022.Header3 {
	a.Header = new(iso20022.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferV02) AddBatchTransfer() *iso20022.CardPaymentBatchTransfer1 {
	a.BatchTransfer = new(iso20022.CardPaymentBatchTransfer1)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV02) AddSecurityTrailer() *iso20022.ContentInformationType4 {
	a.SecurityTrailer = new(iso20022.ContentInformationType4)
	return a.SecurityTrailer
}
