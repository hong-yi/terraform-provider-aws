// Code generated by generators/servicefilters/main.go; DO NOT EDIT.

package namevaluesfilters

import ( // nosemgrep: aws-sdk-go-multiple-service-imports
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// []*SERVICE.Filter handling

// AutoScalingFilters returns autoscaling service filters.
func (filters NameValuesFilters) AutoScalingFilters() []*autoscaling.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*autoscaling.Filter, 0, len(m))

	for k, v := range m {
		filter := &autoscaling.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// DatabasemigrationserviceFilters returns databasemigrationservice service filters.
func (filters NameValuesFilters) DatabasemigrationserviceFilters() []*databasemigrationservice.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*databasemigrationservice.Filter, 0, len(m))

	for k, v := range m {
		filter := &databasemigrationservice.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// DocDBFilters returns docdb service filters.
func (filters NameValuesFilters) DocDBFilters() []*docdb.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*docdb.Filter, 0, len(m))

	for k, v := range m {
		filter := &docdb.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// EC2Filters returns ec2 service filters.
func (filters NameValuesFilters) EC2Filters() []*ec2.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*ec2.Filter, 0, len(m))

	for k, v := range m {
		filter := &ec2.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// ElasticinferenceFilters returns elasticinference service filters.
func (filters NameValuesFilters) ElasticinferenceFilters() []*elasticinference.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*elasticinference.Filter, 0, len(m))

	for k, v := range m {
		filter := &elasticinference.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// ElasticsearchserviceFilters returns elasticsearchservice service filters.
func (filters NameValuesFilters) ElasticsearchserviceFilters() []*elasticsearchservice.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*elasticsearchservice.Filter, 0, len(m))

	for k, v := range m {
		filter := &elasticsearchservice.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// FSxFilters returns fsx service filters.
func (filters NameValuesFilters) FSxFilters() []*fsx.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*fsx.Filter, 0, len(m))

	for k, v := range m {
		filter := &fsx.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// ImagebuilderFilters returns imagebuilder service filters.
func (filters NameValuesFilters) ImagebuilderFilters() []*imagebuilder.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*imagebuilder.Filter, 0, len(m))

	for k, v := range m {
		filter := &imagebuilder.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// LicensemanagerFilters returns licensemanager service filters.
func (filters NameValuesFilters) LicensemanagerFilters() []*licensemanager.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*licensemanager.Filter, 0, len(m))

	for k, v := range m {
		filter := &licensemanager.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// NeptuneFilters returns neptune service filters.
func (filters NameValuesFilters) NeptuneFilters() []*neptune.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*neptune.Filter, 0, len(m))

	for k, v := range m {
		filter := &neptune.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// RdsFilters returns rds service filters.
func (filters NameValuesFilters) RdsFilters() []*rds.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*rds.Filter, 0, len(m))

	for k, v := range m {
		filter := &rds.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// ResourcegroupstaggingapiFilters returns resourcegroupstaggingapi service filters.
func (filters NameValuesFilters) ResourcegroupstaggingapiFilters() []*resourcegroupstaggingapi.TagFilter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*resourcegroupstaggingapi.TagFilter, 0, len(m))

	for k, v := range m {
		filter := &resourcegroupstaggingapi.TagFilter{
			Key:    aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// Route53resolverFilters returns route53resolver service filters.
func (filters NameValuesFilters) Route53resolverFilters() []*route53resolver.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*route53resolver.Filter, 0, len(m))

	for k, v := range m {
		filter := &route53resolver.Filter{
			Name:   aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}

// SecretsmanagerFilters returns secretsmanager service filters.
func (filters NameValuesFilters) SecretsmanagerFilters() []*secretsmanager.Filter {
	m := filters.Map()

	if len(m) == 0 {
		return nil
	}

	result := make([]*secretsmanager.Filter, 0, len(m))

	for k, v := range m {
		filter := &secretsmanager.Filter{
			Key:    aws.String(k),
			Values: aws.StringSlice(v),
		}

		result = append(result, filter)
	}

	return result
}
