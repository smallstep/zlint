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

import (
	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type rootCAKeyUsageMustBeCritical struct{}

func (l *rootCAKeyUsageMustBeCritical) Initialize() error {
	return nil
}

func (l *rootCAKeyUsageMustBeCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsRootCA(c) && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *rootCAKeyUsageMustBeCritical) Execute(c *x509.Certificate) *LintResult {
	keyUsageExtension := util.GetExtFromCert(c, util.KeyUsageOID)
	if keyUsageExtension.Critical {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_root_ca_key_usage_must_be_critical",
		Description:   "Root CA certificates MUST have Key Usage Extension marked critical",
		Citation:      "BRs: 7.1.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.RFC2459Date,
		Lint:          &rootCAKeyUsageMustBeCritical{},
	})
}
