/*behavior pattern*/
package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtector(user string, passwordName string, hashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{user: user, passwordName: passwordName, hashAlgorithm: hashAlgorithm}
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

type SHA struct{}

func (S SHA) Hash(p *PasswordProtector) {
	fmt.Printf("hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (M MD5) Hash(p *PasswordProtector) {
	fmt.Printf("hashing using MD5 for %s\n", p.passwordName)
}
