package structs

type Number struct{
	PhoneNumber string
	Name string
	Group string

	DateAdded string
}

func NewNumber (phoneNumber string,
	name string, group string,
	dateAdded string) Number{
		return Number{
			PhoneNumber: phoneNumber,
			Name: name,
			Group: group,
			DateAdded: dateAdded,
		}
}

func (nb *Number) ChangeName(NewName string) {
	nb.Name = NewName
}
func (nb *Number) ChangeGroup(NewGroup string){
	nb.Group = NewGroup
}