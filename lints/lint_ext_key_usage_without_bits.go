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

/***********************************************************************
 This profile does not restrict the combinations of bits that may be
   set in an instantiation of the keyUsage extension.  However,
   appropriate values for keyUsage extensions for particular algorithms
   are specified in [RFC3279], [RFC4055], and [RFC4491].  When the
   keyUsage extension appears in a certificate, at least one of the bits
   MUST be set to 1.
***********************************************************************/

import (
	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type keyUsageBitsSet struct{}

func (l *keyUsageBitsSet) Initialize() error {
	return nil
}

func (l *keyUsageBitsSet) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *keyUsageBitsSet) Execute(c *x509.Certificate) *LintResult {
	if c.KeyUsage == 0 {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_key_usage_without_bits",
		Description:   "When the keyUsage extension is included, at least one bit MUST be set to 1",
		Citation:      "RFC 5280: 4.2.1.3",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &keyUsageBitsSet{},
	})
}
