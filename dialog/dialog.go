package dialog

type Variant struct {
	Id   string
	AdditionalId string // additional id or some info that will be appended to a button/link
	Text string // text visible to the user
	Url string // optional URL
	RowId int // from 1
}

type Dialog struct {
	Id       string
	Text     string // dialog header visible to the user
	Variants []Variant
}
