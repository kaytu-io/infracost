package aws

import (
	"github.com/kaytu-io/infracost/external/resources/aws"
	"github.com/kaytu-io/infracost/external/schema"
)

func getKinesisStreamRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_kinesis_stream",
		CoreRFunc: newKinesisStream,
	}
}

func newKinesisStream(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()
	StreamMode := d.Get("stream_mode_details.0.stream_mode").String()
	ShardCount := d.Get("shard_count").Int()

	return &aws.KinesisStream{
		Address:    d.Address,
		Region:     region,
		StreamMode: StreamMode,
		ShardCount: ShardCount,
	}
}
