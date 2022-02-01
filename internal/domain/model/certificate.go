package model

import(
)

type certificate struct {
	token 		CertOntineToken
	domain 		CertDomain
	lifecycle	CertLifeCycle	// for cookie
}

type CertOntineToken	string
type CertDomain			string
type CertLifeCycle		int


type Certificate interface {
	GetToken()		string
	GetLifeCycle()	int	
	GetDomain()		string
}

func NewCertificate(t string, expr int, domain string) Certificate {

	return &certificate{
		token: CertOntineToken(t),
		domain: CertDomain(domain),
		lifecycle: CertLifeCycle(expr),
	}
}

func (c *certificate) GetToken() string {
	return string(c.token)
}

func (c *certificate) GetLifeCycle() int {
	return int(c.lifecycle)
}

func (c *certificate) GetDomain() string {
	return string(c.domain)
}
