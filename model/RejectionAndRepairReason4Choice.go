package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason4Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason22Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason4Choice) SetCode(value string) {
	r.Code = (*RejectionReason22Code)(&value)
}

func (r *RejectionAndRepairReason4Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
