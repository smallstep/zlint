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

/************************************************************************
BRs: 7.1.2.3
cRLDistributionPoints
This extension MAY be present. If present, it MUST NOT be marked critical, and it MUST contain the
HTTP URL of the CA’s CRL service. See Section 13.2.1 for details.
*************************************************************************/

import (
	"strings"

	"github.com/smallstep/zcrypto/x509"
	"github.com/smallstep/zlint/util"
)

type subCertIssuerUrl struct{}

func (l *subCertIssuerUrl) Initialize() error {
	return nil
}

func (l *subCertIssuerUrl) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertIssuerUrl) Execute(c *x509.Certificate) *LintResult {
	for _, url := range c.IssuingCertificateURL {
		if strings.HasPrefix(url, "http://") {
			return &LintResult{Status: Pass}
		}
	}
	return &LintResult{Status: Warn}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_aia_does_not_contain_issuing_ca_url",
		Description:   "Subscriber certificates authorityInformationAccess extension should contain the HTTP URL of the issuing CA’s certificate",
		Citation:      "BRs: 7.1.2.3",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCertIssuerUrl{},
	})
}
