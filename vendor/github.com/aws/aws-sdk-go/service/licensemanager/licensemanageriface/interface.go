// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package licensemanageriface provides an interface to enable mocking the AWS License Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package licensemanageriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/licensemanager"
)

// LicenseManagerAPI provides an interface to enable mocking the
// licensemanager.LicenseManager service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS License Manager.
//    func myFunc(svc licensemanageriface.LicenseManagerAPI) bool {
//        // Make svc.CreateLicenseConfiguration request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := licensemanager.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLicenseManagerClient struct {
//        licensemanageriface.LicenseManagerAPI
//    }
//    func (m *mockLicenseManagerClient) CreateLicenseConfiguration(input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLicenseManagerClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LicenseManagerAPI interface {
	CreateLicenseConfiguration(*licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error)
	CreateLicenseConfigurationWithContext(aws.Context, *licensemanager.CreateLicenseConfigurationInput, ...request.Option) (*licensemanager.CreateLicenseConfigurationOutput, error)
	CreateLicenseConfigurationRequest(*licensemanager.CreateLicenseConfigurationInput) (*request.Request, *licensemanager.CreateLicenseConfigurationOutput)

	DeleteLicenseConfiguration(*licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error)
	DeleteLicenseConfigurationWithContext(aws.Context, *licensemanager.DeleteLicenseConfigurationInput, ...request.Option) (*licensemanager.DeleteLicenseConfigurationOutput, error)
	DeleteLicenseConfigurationRequest(*licensemanager.DeleteLicenseConfigurationInput) (*request.Request, *licensemanager.DeleteLicenseConfigurationOutput)

	GetLicenseConfiguration(*licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error)
	GetLicenseConfigurationWithContext(aws.Context, *licensemanager.GetLicenseConfigurationInput, ...request.Option) (*licensemanager.GetLicenseConfigurationOutput, error)
	GetLicenseConfigurationRequest(*licensemanager.GetLicenseConfigurationInput) (*request.Request, *licensemanager.GetLicenseConfigurationOutput)

	GetServiceSettings(*licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error)
	GetServiceSettingsWithContext(aws.Context, *licensemanager.GetServiceSettingsInput, ...request.Option) (*licensemanager.GetServiceSettingsOutput, error)
	GetServiceSettingsRequest(*licensemanager.GetServiceSettingsInput) (*request.Request, *licensemanager.GetServiceSettingsOutput)

	ListAssociationsForLicenseConfiguration(*licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
	ListAssociationsForLicenseConfigurationWithContext(aws.Context, *licensemanager.ListAssociationsForLicenseConfigurationInput, ...request.Option) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
	ListAssociationsForLicenseConfigurationRequest(*licensemanager.ListAssociationsForLicenseConfigurationInput) (*request.Request, *licensemanager.ListAssociationsForLicenseConfigurationOutput)

	ListFailuresForLicenseConfigurationOperations(*licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
	ListFailuresForLicenseConfigurationOperationsWithContext(aws.Context, *licensemanager.ListFailuresForLicenseConfigurationOperationsInput, ...request.Option) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
	ListFailuresForLicenseConfigurationOperationsRequest(*licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*request.Request, *licensemanager.ListFailuresForLicenseConfigurationOperationsOutput)

	ListLicenseConfigurations(*licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error)
	ListLicenseConfigurationsWithContext(aws.Context, *licensemanager.ListLicenseConfigurationsInput, ...request.Option) (*licensemanager.ListLicenseConfigurationsOutput, error)
	ListLicenseConfigurationsRequest(*licensemanager.ListLicenseConfigurationsInput) (*request.Request, *licensemanager.ListLicenseConfigurationsOutput)

	ListLicenseSpecificationsForResource(*licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
	ListLicenseSpecificationsForResourceWithContext(aws.Context, *licensemanager.ListLicenseSpecificationsForResourceInput, ...request.Option) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
	ListLicenseSpecificationsForResourceRequest(*licensemanager.ListLicenseSpecificationsForResourceInput) (*request.Request, *licensemanager.ListLicenseSpecificationsForResourceOutput)

	ListResourceInventory(*licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error)
	ListResourceInventoryWithContext(aws.Context, *licensemanager.ListResourceInventoryInput, ...request.Option) (*licensemanager.ListResourceInventoryOutput, error)
	ListResourceInventoryRequest(*licensemanager.ListResourceInventoryInput) (*request.Request, *licensemanager.ListResourceInventoryOutput)

	ListTagsForResource(*licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *licensemanager.ListTagsForResourceInput, ...request.Option) (*licensemanager.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*licensemanager.ListTagsForResourceInput) (*request.Request, *licensemanager.ListTagsForResourceOutput)

	ListUsageForLicenseConfiguration(*licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
	ListUsageForLicenseConfigurationWithContext(aws.Context, *licensemanager.ListUsageForLicenseConfigurationInput, ...request.Option) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
	ListUsageForLicenseConfigurationRequest(*licensemanager.ListUsageForLicenseConfigurationInput) (*request.Request, *licensemanager.ListUsageForLicenseConfigurationOutput)

	TagResource(*licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *licensemanager.TagResourceInput, ...request.Option) (*licensemanager.TagResourceOutput, error)
	TagResourceRequest(*licensemanager.TagResourceInput) (*request.Request, *licensemanager.TagResourceOutput)

	UntagResource(*licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *licensemanager.UntagResourceInput, ...request.Option) (*licensemanager.UntagResourceOutput, error)
	UntagResourceRequest(*licensemanager.UntagResourceInput) (*request.Request, *licensemanager.UntagResourceOutput)

	UpdateLicenseConfiguration(*licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error)
	UpdateLicenseConfigurationWithContext(aws.Context, *licensemanager.UpdateLicenseConfigurationInput, ...request.Option) (*licensemanager.UpdateLicenseConfigurationOutput, error)
	UpdateLicenseConfigurationRequest(*licensemanager.UpdateLicenseConfigurationInput) (*request.Request, *licensemanager.UpdateLicenseConfigurationOutput)

	UpdateLicenseSpecificationsForResource(*licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
	UpdateLicenseSpecificationsForResourceWithContext(aws.Context, *licensemanager.UpdateLicenseSpecificationsForResourceInput, ...request.Option) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
	UpdateLicenseSpecificationsForResourceRequest(*licensemanager.UpdateLicenseSpecificationsForResourceInput) (*request.Request, *licensemanager.UpdateLicenseSpecificationsForResourceOutput)

	UpdateServiceSettings(*licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error)
	UpdateServiceSettingsWithContext(aws.Context, *licensemanager.UpdateServiceSettingsInput, ...request.Option) (*licensemanager.UpdateServiceSettingsOutput, error)
	UpdateServiceSettingsRequest(*licensemanager.UpdateServiceSettingsInput) (*request.Request, *licensemanager.UpdateServiceSettingsOutput)
}

var _ LicenseManagerAPI = (*licensemanager.LicenseManager)(nil)
