package role

type Role int

const (
	Buyer Role = iota
	Manager
	Admin
)

//func (r Role) String() string {
//	return string(r)
//}
