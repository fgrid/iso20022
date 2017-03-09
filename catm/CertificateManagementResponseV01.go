package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.008.001.01 Document"`
	Message *CertificateManagementResponseV01 `xml:"CertMgmtRspn"`
}

func (d *Document00800101) AddMessage() *CertificateManagementResponseV01 {
	d.Message = new(CertificateManagementResponseV01)
	return d.Message
}

// The CertificateManagementResponse is sent by a terminal manager in response to a CertificateManagementRequest to provide the outcome of the requested service.
type CertificateManagementResponseV01 struct {

	// Information related to the protocol management.
	Header *iso20022.Header29 `xml:"Hdr"`

	// Information related to the result of the certificate management request.
	CertificateManagementResponse *iso20022.CertificateManagementResponse1 `xml:"CertMgmtRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (c *CertificateManagementResponseV01) AddHeader() *iso20022.Header29 {
	c.Header = new(iso20022.Header29)
	return c.Header
}

func (c *CertificateManagementResponseV01) AddCertificateManagementResponse() *iso20022.CertificateManagementResponse1 {
	c.CertificateManagementResponse = new(iso20022.CertificateManagementResponse1)
	return c.CertificateManagementResponse
}

func (c *CertificateManagementResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType13 {
	c.SecurityTrailer = new(iso20022.ContentInformationType13)
	return c.SecurityTrailer
}
