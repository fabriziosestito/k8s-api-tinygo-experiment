package publicsuffix

import _ "embed"

type uint40String string

type list struct{}

type uint32String string

func PublicSuffix(domain string) (publicSuffix string, icann bool) {
	panic("stub")
}

func EffectiveTLDPlusOne(domain string) (string, error) {
	panic("stub")
}

type Embedme interface{}
