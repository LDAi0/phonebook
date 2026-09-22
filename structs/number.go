package structs

import "time"

type Number struct{
	PhoneNumber string
	Name string
	Group string

	DateAdded time.Time
}

func NewNumber (phoneNumber string,
	name string, group string) Number{
		return Number{
			PhoneNumber: phoneNumber,
			Name: name,
			Group: group,
			DateAdded: time.Now(),
		}
}

func (nb *Number) ChangeName(NewName string) {
	nb.Name = NewName
}
func (nb *Number) ChangeGroup(NewGroup string){
	nb.Group = NewGroup
}