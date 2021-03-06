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
BRs: 7.1.2.1c certificatePolicies
This extension SHOULD NOT be present.
************************************************/

import (
	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type rootCAContainsCertPolicy struct{}

func (l *rootCAContainsCertPolicy) Initialize() error {
	return nil
}

func (l *rootCAContainsCertPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsRootCA(c)
}

func (l *rootCAContainsCertPolicy) Execute(c *x509.Certificate) *LintResult {
	if util.IsExtInCert(c, util.CertPolicyOID) {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_root_ca_contains_cert_policy",
		Description:   "Root CA Certificate: certificatePolicies SHOULD NOT be present.",
		Citation:      "BRs: 7.1.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &rootCAContainsCertPolicy{},
	})
}
