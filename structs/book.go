package structs

type Book struct{
    numbers map[string]Number
}

func NewBook() Book{
    return Book{}
}

func (b *Book) AddNumber(number Number) error{
    if _, ok :=b.numbers[number.PhoneNumber]; ok{
        return ErrNumberAlreadyExist
    }
    b.numbers[number.PhoneNumber] = number
    return nil
}

func (b *Book) GetNumber (phonenumber string) (Number, error){
    if _, ok := b.numbers[phonenumber]; !ok{
        return Number{}, ErrNumberNotFound
    }
    return b.numbers[phonenumber], nil
}

func (b *Book) GetNumbers () map[string]Number{
    return b.numbers
}

func (b *Book) DeleteNumber(number Number) error{
    if _, ok :=b.numbers[number.PhoneNumber]; !ok{
        return ErrNumberNotFound
    }
    delete(b.numbers,number.PhoneNumber)
    return nil
}