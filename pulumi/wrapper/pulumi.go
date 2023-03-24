package wrapper

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optpreview"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
)

type Response struct {
	ID       string
	Response string
	StdOut   string
}

func Preview(ctx context.Context, stack auto.Stack) (*Response, error) {
	buf := new(bytes.Buffer)

	previewResult, err := stack.Preview(ctx, optpreview.ProgressStreams(buf))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response := &Response{
		ID:       stack.Name(),
		Response: fmt.Sprintf("%+v", previewResult),
		StdOut:   buf.String(),
	}
	return response, nil
}

func Up(ctx context.Context, stack auto.Stack) (*Response, error) {
	buf := new(bytes.Buffer)
	upRes, err := stack.Up(ctx, optup.ProgressStreams(buf))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response := &Response{
		ID:       stack.Name(),
		Response: fmt.Sprintf("%+v", upRes),
		StdOut:   buf.String(),
	}
	return response, nil
}

func Destroy(ctx context.Context, stack auto.Stack) (*Response, error) {
	buf := new(bytes.Buffer)

	destroyResult, err := stack.Destroy(ctx, optdestroy.ProgressStreams(buf))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response := &Response{
		ID:       stack.Name(),
		Response: fmt.Sprintf("%+v", destroyResult),
		StdOut:   buf.String(),
	}
	return response, nil
}
