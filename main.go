package main

import (
	"context"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	program "github.com/samuelneves/aws-ephemeral-accounts/program"
	"github.com/samuelneves/aws-ephemeral-accounts/pulumi/wrapper"
	"github.com/samuelneves/aws-ephemeral-accounts/utils"
	"os"
)

func main() {
	ensurePlugins()
	cfg := utils.GetConfig()
	ctx := context.Background()

	project := cfg.ProgramName
	fmt.Println("%+v", cfg)
	stackName := "bucket"
	pgr := program.PulumiProgram("programa")

	handleRequest(ctx, stackName, project, pgr)
}

func handleRequest(ctx context.Context, stackName string, project string, pgr pulumi.RunFunc) {
	s, err := auto.NewStackInlineSource(ctx, stackName, project, pgr)
	if err != nil {
		// if stack already exists, 409
		if auto.IsCreateStack409Error(err) {
			fmt.Printf("stack %q already exists", stackName)

		}
	}
	s.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-east-1"})
	s.SetConfig(ctx, "bucket-name", auto.ConfigValue{Value: "ephemeral-aws-account1"})

	res, err := wrapper.Up(ctx, s)
	fmt.Printf("resposta:", res)
}

// ensure plugins runs once before the server boots up
// making sure the proper pulumi plugins are installed
func ensurePlugins() {
	ctx := context.Background()
	w, err := auto.NewLocalWorkspace(ctx)
	if err != nil {

		fmt.Printf("Failed to setup and run http server: %v\n", err)
		os.Exit(1)
	}
	err = w.InstallPlugin(ctx, "aws", "v3.2.1")
	if err != nil {
		fmt.Printf("Failed to install program plugins: %v\n", err)
		os.Exit(1)
	}
}
