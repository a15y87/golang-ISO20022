package iso20022

import (
	"encoding/xml"
	"golang.org/x/text/currency"
	"time"
)

type CodeSet int

const (
	//BalanceType12Code
	Expected CodeSet = iota + 1
	OpeningAvailable
	InterimAvailable
	ClosingAvailable
	ForwardAvailable
	ClosingBooked
	InterimBooked
	OpeningBooked
	PreviouslyClosedBooked
	Information

	Booked
	Pending
	// Information

	Credit
	Debit

	All
	Changed
	Modified

	// AddressType2Code
	Postal
	POBox
	Residential
	Business
	MailTo
	DeliveryTo
)

var CodeSetString = [...]string{
	//BalanceType12Code
	"XPCD",
	"OPAV",
	"ITAV",
	"CLAV",
	"FWAV",
	"CLBD",
	"ITBD",
	"OPBD",
	"PRCD",
	"INFO",

	"BOOK",
	"PDNG",
	//"INFO",

	"CRDT",
	"DBIT",

	"ALLL",
	"CHNG",
	"MODF",

	//AddressType2Code
	"ADDR",
	"PBOX",
	"HOME",
	"BIZZ",
	"MLTO",
	"DLVY",
}

func (s CodeSet) String() string {
	return CodeSetString[s-1]
}
func (s CodeSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(s.String(), start)
}

type ISO8601DateTime struct {
	time.Time
}

func (t ISO8601DateTime) MarshalText() (result []byte, err error) {
	result, err = []byte(t.Format("2006-01-02T15:04:05-0700")), nil
	return
}

func (t *ISO8601DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(t.Format("2006-01-02T15:04:05-0700"), start)
	return nil
}

type ISO8601Time struct {
	time.Time
}

func (t ISO8601Time) MarshalText() (result []byte, err error) {
	result, err = []byte(t.Format("15:04:05")), nil
	return
}

func (t *ISO8601Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(t.Format("15:04:05"), start)
	return nil
}

type ISO8601Date struct {
	time.Time
}

func (t ISO8601Date) MarshalText() (result []byte, err error) {
	result, err = []byte(t.Format("2006-01-02")), nil
	return
}

func (t *ISO8601Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(t.Format("2006-01-02"), start)
	return nil
}

type CashAccountType2Choice struct {
	Code        CodeSet `xml:"Cd"`
	Proprietary string  `xml:"Prtry"`
}

type OrganisationIdentification8 struct {
	AnyBIC string                                `xml:"AnyBIC,omitempty"` //[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1} ISO 9362
	Other  []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

type Party11Choice struct {
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId,omitempty"`
	PrivateIdentification      *PersonIdentification5       `xml:"PrvtId,omitempty"`
}

type PartyIdentification43 struct {
	Name           string          `xml:"Nm,omitempty"`
	PostalAddress  *PostalAddress6 `xlm:"PstlAdr,omitempty"`
	Identification *Party11Choice  `xlm:"ID,omitempty"`
}

type CashAccount24 struct {
	Identification string                 `xml:"Id"`
	Type           CashAccountType2Choice `xml:"Tp,omitempty"`
	Currency       currency.Amount        `xml:"Ccy,omitempty"` //ActiveOrHistoricCurrencyCode
	Name           string                 `xml:"Nm,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDate       ISO8601Date `xml:"BirthDt"`
	ProvinceOfBirth string      `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth     string      `xml:"CityOfBirth"`
	CountryOfBirth  string      `xml:"CtryOfBirth"`
}

type PersonIdentification5 struct {
	DateAndPlaceOfBirth *DateAndPlaceOfBirth            `xml:"DtAndPlcOfBirth,omitempty"`
	Other               []*GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Identification string                                 `xml:"Id"`
	SchemeName     *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issuer         string                                 `xml:"Issr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Code        string `xml:"Cd,omitempty"` //ExternalPersonIdentification1Code 1:4
	Proprietary string `xml:"Prtry,omitempty"`
}

type Party12Choice struct {
	Party *PartyIdentification43                        `xml:"Pty,omitempty"`
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty"`
}

type ClearingSystemIdentification struct {
	Code        string `xml:"Cd,omitempty"` //ExternalClearingSystemIdentification1Code 1:5
	Proprietary string `sml:"Prtry,omitempty"`
}

type ClearingSystemMemberIdentification struct {
	ClearingSystemIdentification *ClearingSystemIdentification `xml:"ClrSysId,omitempty"`
	MemberIdentification         string                        `xml:"MmbId"`
}

type PostalAddress6 struct {
	// PostalAddress6
	AddressType        CodeSet  `xml:"AdrTp,omitempty"`
	Department         string   `xml:"Dept,omitempty"`
	SubDepartment      string   `xml:"SubDept,omitempty"`
	StreetName         string   `xml:"StrtNm,omitempty"`
	BuildingNumber     string   `xml:"BldgNb,omitempty"`
	PostCode           string   `xml:"PstCd,omitempty"`
	TownName           string   `xml:"TwnNm,omitempty"`
	CountrySubDivision string   `xml:"CtrySubDvsn,omitempty"`
	Country            string   `xml:"Ctry,omitempty"`
	AddressLine        []string `xml:"AdrLine,omitempty"`
}

type SchemeName struct {
	Code        string `xml:"Cd,omitempty"` //ExternalFinancialInstitutionIdentification1Code 1:4
	Proprietary string `sml:"Prtry,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Identification string      `xml:"Id"`
	SchemeName     *SchemeName `xml:"SchmeNm,omitempty"`
	Issuer         string      `xml:"Issr,omitempty"`
}

type FinancialInstitutionIdentification struct {
	BICFI                              string                              `xml:"BICFI,omitempty"`
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification `xml:"ClrSysMmbId,omitempty"`
	Name                               string                              `xml:"Nm,omitempty"`
	PostalAddress                      *PostalAddress6                     `xml:"PstlAdr,omitempty"`
	Other                              *GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

type BranchIdentification struct {
	Identification string          `xml:"Id,omitempty"`
	Name           string          `xml:"Nm,omitempty"`
	PostalAddress  *PostalAddress6 `xml:"PstlAdr,omitempty"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinancialInstitutionIdentification *FinancialInstitutionIdentification `sml:"FinInstnId"`
	BranchIdentification               *BranchIdentification               `sml:"BrnchId,omitempty"`
}

type FromToDate struct {
	FromDate ISO8601Date `xml:"FrDt"`
	ToDate   ISO8601Date `xml:"ToDt,omitempty"`
}

type FromToTime struct {
	FromTime ISO8601Time
	ToTime   ISO8601Time
}

type ReportingPeriod struct {
	FromToDate *FromToDate `xml:"FrToDt"`
	FromToTime *FromToTime `xml:"FrToTm"`
	Type       CodeSet     `xml:"Tp"`
}

type FloorLimit struct {
	Amount               currency.Amount `xml:"Amt"`
	CreditDebitIndicator CodeSet         `xml:"CdtDbtInd"`
}

type RequestedTransactionType struct {
	Status               CodeSet     `xml:"Sts"`
	CreditDebitIndicator CodeSet     `xml:"CdtDbtInd"`
	FloorLimit           *FloorLimit `xml:"FlrLmt,omitempty"`
}

type RequestedBalanceType struct {
	CodeOrProprietary *CodeOrProprietary `xml:"CdOrPrtry"`
	SubType           *SubType           `xml:"SubTp,omitempty"`
}

type SubType struct {
	Code        string `xml:"Cd,omitempty"` //ExternalBalanceSubType1Code
	Proprietary string `sml:"Prtry,omitempty"`
}

type CodeOrProprietary struct {
	Code        CodeSet `xml:"Cd,omitempty"`
	Proprietary string  `sml:"Prtry,omitempty"`
}

type ReportingRequest struct {
	Identification                     string                                        `xml:"Id,omitempty"`
	RequestedMessageNameIdentification string                                        `xml:"ReqdMsgNmId"`
	Account                            *CashAccount24                                `xml:"Acct,omitempty"`
	AccountOwner                       *Party12Choice                                `xml:"AcctOwnr"`
	AccountServicer                    *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`
	ReportingPeriod                    *ReportingPeriod                              `xml:"RptgPrd,omitempty"`
	RequestedTransactionType           *RequestedTransactionType                     `xml:"ReqdTxTp,omitempty"`
	RequestedBalanceType               *RequestedBalanceType                         `xml:"ReqdBalTp,omitempty"`
}

type GroupHeader struct {
	MessageIdentification string          `xml:"MsgId"`
	CreationDateTime      ISO8601DateTime `xml:"CreDtTm"`
	MessageSender         string          `xml:"MsgSndr,omitempty"`
}

type AccountReportingRequest struct {
	GroupHeader      *GroupHeader      `xml:"GrpHdr"`
	ReportingRequest *ReportingRequest `xml:"RptgReq"`
}

type AccountReportingRequestV03 struct {
	XMLName                 xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Document"`
	Xmlns_xsi               string                   `xml:"xmlns:xsi,attr"`
	AccountReportingRequest *AccountReportingRequest `xml:"AcctRptgReq"`
}

func (r AccountReportingRequestV03) init() {
	r.Xmlns_xsi="http://www.w3.org/2001/XMLSchema-instance"
	r.AccountReportingRequest = &AccountReportingRequest{}

}