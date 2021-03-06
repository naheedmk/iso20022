package model

// Choice between a code and or a data source scheme to determine the account type.
type PurposeCode2Choice struct {

	// Securities account purpose as an ISO 20022 code.
	Code *SecuritiesAccountPurposeType1Code `xml:"Cd"`

	// Securities account purpose as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PurposeCode2Choice) SetCode(value string) {
	p.Code = (*SecuritiesAccountPurposeType1Code)(&value)
}

func (p *PurposeCode2Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
