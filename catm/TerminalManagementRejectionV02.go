package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.004.001.02 Document"`
	Message *TerminalManagementRejectionV02 `xml:"TermnlMgmtRjctn"`
}

func (d *Document00400102) AddMessage() *TerminalManagementRejectionV02 {
	d.Message = new(TerminalManagementRejectionV02)
	return d.Message
}

// The TerminalManagementRejection message is sent by the terminal manager to reject a message request sent by an acceptor, to indicate that the received message could not be processed.
type TerminalManagementRejectionV02 struct {

	// Rejection message management information.
	Header *iso20022.Header6 `xml:"Hdr"`

	// Information related to the reject.
	Reject *iso20022.AcceptorRejection2 `xml:"Rjct"`
}

func (t *TerminalManagementRejectionV02) AddHeader() *iso20022.Header6 {
	t.Header = new(iso20022.Header6)
	return t.Header
}

func (t *TerminalManagementRejectionV02) AddReject() *iso20022.AcceptorRejection2 {
	t.Reject = new(iso20022.AcceptorRejection2)
	return t.Reject
}
