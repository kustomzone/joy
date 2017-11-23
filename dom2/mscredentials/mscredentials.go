package mscredentials

import "github.com/matthewmueller/golly/js"

// js:"MSCredentials,omit"
type MSCredentials struct {
}

// GetAssertion
func (*MSCredentials) GetAssertion(challenge string, filter *mscredentialfilter.MSCredentialFilter, params *mssignatureparameters.MSSignatureParameters) (m msassertion.MSAssertion) {
	js.Rewrite("await $<.GetAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

// MakeCredential
func (*MSCredentials) MakeCredential(accountInfo *msaccountinfo.MSAccountInfo, params []*mscredentialparameters.MSCredentialParameters, challenge *string) (m msassertion.MSAssertion) {
	js.Rewrite("await $<.MakeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
