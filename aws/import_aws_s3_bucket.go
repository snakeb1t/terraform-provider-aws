package aws

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsS3BucketImportState(
	d *schema.ResourceData,
	meta interface{}) ([]*schema.ResourceData, error) {

	results := make([]*schema.ResourceData, 1)
	d.Set("force_destroy", false)
	d.Set("acl", "private")
	results[0] = d

	return results, nil
}
