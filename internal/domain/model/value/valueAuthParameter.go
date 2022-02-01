package value

type valueAuthParameter struct {
	Domain 		AuthDomain
	Expiration  AuthExpiration
	Schema 		AuthSchema
}

type AuthDomain		string
type AuthExpiration	int
type AuthSchema		int

type AuthParameter interface {
	GetDomain() string
	GetExpiration() int
	GetSchema() int
}

func NewAuthParameter(domain string, expir int, schema int) AuthParameter {
	return &valueAuthParameter{
		Domain:		AuthDomain(domain),
		Expiration:	AuthExpiration(expir),
		Schema:		AuthSchema(schema),
	}
}

func (v *valueAuthParameter) GetExpiration() int {
	return int(v.Expiration)
}

func (v *valueAuthParameter) GetDomain() string {
	return string(v.Domain)
}

func (v *valueAuthParameter) GetSchema() int {
	return int(v.Schema)
}