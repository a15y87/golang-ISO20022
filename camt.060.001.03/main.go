package AccountReportingRequestV03

import (
	"encoding/xml"
	"github.com/a15y87/golang-ISO20022"
)

type AccountReportingRequestV03 struct {
	XMLName                 xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Document"`
	Xmlns_xsi               string                   `xml:"xmlns:xsi,attr"`
	AccountReportingRequest *iso20022.AccountReportingRequest `xml:"AcctRptgReq"`
}

func New() (a *AccountReportingRequestV03) {
	a = &AccountReportingRequestV03{Xmlns_xsi: "http://www.w3.org/2001/XMLSchema-instance"}
	a.AccountReportingRequest = &iso20022.AccountReportingRequest{}
	a.AccountReportingRequest.GroupHeader = &iso20022.GroupHeader59{}
	a.AccountReportingRequest.ReportingRequest = &iso20022.ReportingRequest3{}
	a.AccountReportingRequest.ReportingRequest.Account = &iso20022.CashAccount24{}
	a.AccountReportingRequest.ReportingRequest.AccountOwner = &iso20022.Party12Choice{Party: &iso20022.PartyIdentification43{}}
	a.AccountReportingRequest.ReportingRequest.ReportingPeriod = &iso20022.ReportingPeriod1{}
	a.AccountReportingRequest.ReportingRequest.ReportingPeriod.FromToDate = &iso20022.FromToDate{}
	a.AccountReportingRequest.ReportingRequest.ReportingPeriod.FromToTime = &iso20022.FromToTime{}
	return
}
