package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************
RFC 5280: 4.1.2.4
The issuer field identifies the entity that has signed and issued the
   certificate.  The issuer field MUST contain a non-empty distinguished
   name (DN).  The issuer field is defined as the X.501 type Name
   [X.501].
************************************************/

import (
	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type issuerFieldEmpty struct{}

func (l *issuerFieldEmpty) Initialize() error {
	return nil
}

func (l *issuerFieldEmpty) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *issuerFieldEmpty) Execute(c *x509.Certificate) *LintResult {
	if &c.Issuer != nil && util.NotAllNameFieldsAreEmpty(&c.Issuer) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_issuer_field_empty",
		Description:   "Certificate issuer field MUST NOT be empty and must have a non-empty distingushed name",
		Citation:      "RFC 5280: 4.1.2.4",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &issuerFieldEmpty{},
	})
}
