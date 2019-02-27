package aws

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func resourceAwsS3BucketImportState(
	d *schema.ResourceData,
	meta interface{}) ([]*schema.ResourceData, error) {

	results := make([]*schema.ResourceData, 1)
	idData := strings.Split(d.Id(), "_")

	d.Set("force_destroy", false)
	if len(idData) == 2 {
		d.Set("acl", idData[1])
	} else {
		d.Set("acl", "private")
	}
	d.SetId(idData[0])
	results[0] = d

	return results, nil
}
