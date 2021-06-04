package entity

type ID string

const EmptyID ID = ""

func StrToID(s string) ID {
	return ID(s)
}

func (i ID) String() string {
	return string(i)
}