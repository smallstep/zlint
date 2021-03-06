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

/**************************************************************************
BRs: 7.1.2.3
keyUsage (optional)
If present, bit positions for keyCertSign and cRLSign MUST NOT be set.
***************************************************************************/

import (
	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type subCertKeyUsageBitSet struct{}

func (l *subCertKeyUsageBitSet) Initialize() error {
	return nil
}

func (l *subCertKeyUsageBitSet) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID) && !util.IsCACert(c)
}

func (l *subCertKeyUsageBitSet) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	if (c.KeyUsage & x509.KeyUsageCertSign) == x509.KeyUsageCertSign {
		return &LintResult{Status: Error}
	} else { //key usage doesn't allow cert signing or isn't present
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_key_usage_cert_sign_bit_set",
		Description:   "Subscriber Certificate: keyUsage if present, bit positions for keyCertSign and cRLSign MUST NOT be set.",
		Citation:      "BRs: 7.1.2.3",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCertKeyUsageBitSet{},
	})
}
