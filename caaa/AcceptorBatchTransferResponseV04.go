package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.04 Document"`
	Message *AcceptorBatchTransferResponseV04 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200104) AddMessage() *AcceptorBatchTransferResponseV04 {
	d.Message = new(AcceptorBatchTransferResponseV04)
	return d.Message
}

// The AcceptorBatchTransferResponse is sent by the acquirer (or its agent) to inform the acceptor (or its agent) of the transfer in a previous AcceptorBatchTransfer of a collection of transactions.
type AcceptorBatchTransferResponseV04 struct {

	// Capture advice response message management information.
	Header *iso20022.Header12 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	BatchTransferResponse *iso20022.CardPaymentBatchTransferResponse3 `xml:"BtchTrfRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`

}


func (a *AcceptorBatchTransferResponseV04) AddHeader() *iso20022.Header12 {
	a.Header = new(iso20022.Header12)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV04) AddBatchTransferResponse() *iso20022.CardPaymentBatchTransferResponse3 {
	a.BatchTransferResponse = new(iso20022.CardPaymentBatchTransferResponse3)
	return a.BatchTransferResponse
}

func (a *AcceptorBatchTransferResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	a.SecurityTrailer = new(iso20022.ContentInformationType12)
	return a.SecurityTrailer
}

