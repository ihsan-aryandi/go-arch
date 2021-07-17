package entity

type User struct {
	Id       int64 `field_name:"id" min:"5" max:"10" required:"true" exact_length:"10" is_number:"true" `
	Username string
	Password string
}
