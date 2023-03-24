package program

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/organizations"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func PulumiProgram(content string) pulumi.RunFunc {
	return func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		conf := config.New(ctx, "")
		prefix := conf.Require("bucket-name")

		account, err := organizations.NewAccount(ctx, prefix, &organizations.AccountArgs{
			CloseOnDeletion: pulumi.Bool(true),
			Name:            pulumi.String(prefix),
			RoleName:        pulumi.String("pulumi-admin-role"),
			Email:           pulumi.String("samuel.ant.neves+" + prefix + "@gmail.com"),
			ParentId:        pulumi.String("ou-uocx-ptrfzl2q"),
		})
		if err != nil {
			return err
		}

		ctx.Export("account_name", account.Name)
		ctx.Export("account_email", account.Email)
		ctx.Export("account_parent", account.ParentId)
		ctx.Export("account_Arn", account.Arn)
		ctx.Export("assume_role", account.RoleName)
		ctx.Export("account_id", account.ID())

		roleToAssume := pulumi.Sprintf("arn:aws:iam::%s:role/%v", account.ID(), account.RoleName.Elem())

		childAccountProvider, err := aws.NewProvider(ctx, "child-acc-provider", &aws.ProviderArgs{

			AssumeRole: aws.ProviderAssumeRoleArgs{
				RoleArn:     roleToAssume,
				SessionName: pulumi.String("pulumi-child-account-role"),
				Tags:        nil,
			},
			Region: pulumi.String("us-east-1"),
		})

		bucket, err := s3.NewBucket(ctx, "teste", nil, pulumi.Provider(childAccountProvider))
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		//ctx.Export("assumeId", opa.ID())
		return nil
	}
}
