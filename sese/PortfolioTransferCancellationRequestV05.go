package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.014.001.05 Document"`
	Message *PortfolioTransferCancellationRequestV05 `xml:"PrtflTrfCxlReq"`
}

func (d *Document01400105) AddMessage() *PortfolioTransferCancellationRequestV05 {
	d.Message = new(PortfolioTransferCancellationRequestV05)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee), sends the PortfolioTransferCancellationRequest message to the executing party, for example, a (old) plan manager (Transferor), to request the cancellation of a previously sent PortfolioTransferInstruction.
// Usage
// The PortfolioTransferCancellationRequest message is used to request the cancellation of an entire PortfolioTransferInstruction message, ie, all the product transfers that it contained. The cancellation request can be specified either by:
// - quoting the transfer references of all the product transfers listed in the PortfolioTransferInstruction message, or,
// - quoting the details of all the product transfers (this includes TransferReference) listed in PortfolioTransferInstruction message.
// The message identification of the PortfolioTransferInstruction may also be quoted in PreviousReference. It is also possible to request the cancellation of PortfolioTransferInstruction by just quoting its message identification in PreviousReference.
type PortfolioTransferCancellationRequestV05 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Choice between cancellation by transfer details or reference.
	Cancellation *iso20022.Cancellation6Choice `xml:"Cxl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

}


func (p *PortfolioTransferCancellationRequestV05) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferCancellationRequestV05) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PortfolioTransferCancellationRequestV05) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PortfolioTransferCancellationRequestV05) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PortfolioTransferCancellationRequestV05) AddCancellation() *iso20022.Cancellation6Choice {
	p.Cancellation = new(iso20022.Cancellation6Choice)
	return p.Cancellation
}

func (p *PortfolioTransferCancellationRequestV05) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}

