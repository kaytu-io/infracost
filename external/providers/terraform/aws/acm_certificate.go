package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getACMCertificate() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_acm_certificate",
		RFunc: NewACMCertificate,
	}
}
func NewACMCertificate(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.ACMCertificate{
		Address:                 d.Address,
		Region:                  d.Get("region").String(),
		CertificateAuthorityARN: d.Get("certificate_authority_arn").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
