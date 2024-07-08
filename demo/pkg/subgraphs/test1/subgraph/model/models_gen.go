// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BigObject struct {
	NestedObjects []*NestedObject `json:"nestedObjects"`
}

type DeeplyNestedObject struct {
	AFieldOnDeeplyNestedObject string  `json:"aFieldOnDeeplyNestedObject"`
	BFieldOnDeeplyNestedObject int     `json:"bFieldOnDeeplyNestedObject"`
	CFieldOnDeeplyNestedObject bool    `json:"cFieldOnDeeplyNestedObject"`
	DFieldOnDeeplyNestedObject float64 `json:"dFieldOnDeeplyNestedObject"`
	EFieldOnDeeplyNestedObject string  `json:"eFieldOnDeeplyNestedObject"`
	FFieldOnDeeplyNestedObject int     `json:"fFieldOnDeeplyNestedObject"`
	GFieldOnDeeplyNestedObject bool    `json:"gFieldOnDeeplyNestedObject"`
	HFieldOnDeeplyNestedObject float64 `json:"hFieldOnDeeplyNestedObject"`
	IFieldOnDeeplyNestedObject string  `json:"iFieldOnDeeplyNestedObject"`
	JFieldOnDeeplyNestedObject int     `json:"jFieldOnDeeplyNestedObject"`
	KFieldOnDeeplyNestedObject bool    `json:"kFieldOnDeeplyNestedObject"`
	LFieldOnDeeplyNestedObject float64 `json:"lFieldOnDeeplyNestedObject"`
	MFieldOnDeeplyNestedObject string  `json:"mFieldOnDeeplyNestedObject"`
	NFieldOnDeeplyNestedObject int     `json:"nFieldOnDeeplyNestedObject"`
	OFieldOnDeeplyNestedObject bool    `json:"oFieldOnDeeplyNestedObject"`
	PFieldOnDeeplyNestedObject float64 `json:"pFieldOnDeeplyNestedObject"`
	QFieldOnDeeplyNestedObject string  `json:"qFieldOnDeeplyNestedObject"`
	RFieldOnDeeplyNestedObject int     `json:"rFieldOnDeeplyNestedObject"`
	SFieldOnDeeplyNestedObject bool    `json:"sFieldOnDeeplyNestedObject"`
	TFieldOnDeeplyNestedObject float64 `json:"tFieldOnDeeplyNestedObject"`
	UFieldOnDeeplyNestedObject string  `json:"uFieldOnDeeplyNestedObject"`
	VFieldOnDeeplyNestedObject int     `json:"vFieldOnDeeplyNestedObject"`
	WFieldOnDeeplyNestedObject bool    `json:"wFieldOnDeeplyNestedObject"`
	XFieldOnDeeplyNestedObject float64 `json:"xFieldOnDeeplyNestedObject"`
	YFieldOnDeeplyNestedObject string  `json:"yFieldOnDeeplyNestedObject"`
	ZFieldOnDeeplyNestedObject int     `json:"zFieldOnDeeplyNestedObject"`
}

type Employee struct {
	ID               int     `json:"id"`
	FieldThrowsError *string `json:"fieldThrowsError,omitempty"`
}

func (Employee) IsEntity() {}

type NestedObject struct {
	DeeplyNestedObjects []*DeeplyNestedObject `json:"deeplyNestedObjects"`
}

type Query struct {
}

type Subscription struct {
}

type TimestampedString struct {
	// The value of the string.
	Value string `json:"value"`
	// The timestamp when the response was generated.
	UnixTime int `json:"unixTime"`
	// Sequence number
	Seq int `json:"seq"`
	// Total number of responses to be sent
	Total          int                    `json:"total"`
	InitialPayload map[string]interface{} `json:"initialPayload,omitempty"`
}
