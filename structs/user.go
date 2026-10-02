package structs

type User struct{
    Id int 
    Name string
    Password string
    Address string
}

func NewUser (name string, 
    password string,
    address string) User{
    return User{
        Name: name,
        Password: password,
        Address: address,
    }
}
