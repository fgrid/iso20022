package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200105 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.05 Document"`
	Message *AcceptorBatchTransferResponseV05 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200105) AddMessage() *AcceptorBatchTransferResponseV05 {
	d.Message = new(AcceptorBatchTransferResponseV05)
	return d.Message
}

// The AcceptorBatchTransferResponse is sent by the acquirer (or its agent) to inform the acceptor (or its agent) of the transfer in a previous AcceptorBatchTransfer of a collection of transactions.
type AcceptorBatchTransferResponseV05 struct {

	// Capture advice response message management information.
	Header *iso20022.Header25 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	BatchTransferResponse *iso20022.CardPaymentBatchTransferResponse4 `xml:"BtchTrfRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorBatchTransferResponseV05) AddHeader() *iso20022.Header25 {
	a.Header = new(iso20022.Header25)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV05) AddBatchTransferResponse() *iso20022.CardPaymentBatchTransferResponse4 {
	a.BatchTransferResponse = new(iso20022.CardPaymentBatchTransferResponse4)
	return a.BatchTransferResponse
}

func (a *AcceptorBatchTransferResponseV05) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	a.SecurityTrailer = new(iso20022.ContentInformationType12)
	return a.SecurityTrailer
}
