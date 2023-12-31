package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getSQSQueueRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_sqs_queue",
		RFunc: NewSQSQueue,
	}
}

func NewSQSQueue(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.SQSQueue{
		Address:   d.Address,
		Region:    d.Get("region").String(),
		FifoQueue: d.Get("fifo_queue").Bool(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
