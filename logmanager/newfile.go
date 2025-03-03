package logmanager

type People struct {
	idCard  string
	name    string
	address string
}

func NewPeople(idcard string, name string, address string) *People {
	return &People{
		idCard:  idcard,
		name:    name,
		address: address,
	}
}

func (p *People) GetName() string {
	return p.name
}

func (p *People) GetID() string {
	return p.idCard
}
