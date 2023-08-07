// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vercel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel/internal"
)

func GetFile(ctx *pulumi.Context, args *GetFileArgs, opts ...pulumi.InvokeOption) (*GetFileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFileResult
	err := ctx.Invoke("vercel:index/getFile:getFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFile.
type GetFileArgs struct {
	Path string `pulumi:"path"`
}

// A collection of values returned by getFile.
type GetFileResult struct {
	File map[string]string `pulumi:"file"`
	Id   string            `pulumi:"id"`
	Path string            `pulumi:"path"`
}

func GetFileOutput(ctx *pulumi.Context, args GetFileOutputArgs, opts ...pulumi.InvokeOption) GetFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFileResult, error) {
			args := v.(GetFileArgs)
			r, err := GetFile(ctx, &args, opts...)
			var s GetFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFileResultOutput)
}

// A collection of arguments for invoking getFile.
type GetFileOutputArgs struct {
	Path pulumi.StringInput `pulumi:"path"`
}

func (GetFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileArgs)(nil)).Elem()
}

// A collection of values returned by getFile.
type GetFileResultOutput struct{ *pulumi.OutputState }

func (GetFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileResult)(nil)).Elem()
}

func (o GetFileResultOutput) ToGetFileResultOutput() GetFileResultOutput {
	return o
}

func (o GetFileResultOutput) ToGetFileResultOutputWithContext(ctx context.Context) GetFileResultOutput {
	return o
}

func (o GetFileResultOutput) File() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetFileResult) map[string]string { return v.File }).(pulumi.StringMapOutput)
}

func (o GetFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFileResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Path }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFileResultOutput{})
}
