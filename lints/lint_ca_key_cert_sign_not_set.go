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
BRs: 7.1.2.1b
This extension MUST be present and MUST be marked critical. Bit positions for keyCertSign and cRLSign MUST be set.
If the Root CA Private Key is used for signing OCSP responses, then the digitalSignature bit MUST be set.
************************************************/

import (
	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type caKeyCertSignNotSet struct{}

func (l *caKeyCertSignNotSet) Initialize() error {
	return nil
}

func (l *caKeyCertSignNotSet) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *caKeyCertSignNotSet) Execute(c *x509.Certificate) *LintResult {
	if c.KeyUsage&x509.KeyUsageCertSign != 0 {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_key_cert_sign_not_set",
		Description:   "Root CA Certificate: Bit positions for keyCertSign and cRLSign MUST be set.",
		Citation:      "BRs: 7.1.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caKeyCertSignNotSet{},
	})
}
