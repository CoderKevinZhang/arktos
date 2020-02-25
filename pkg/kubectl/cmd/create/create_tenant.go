/*
Copyright 2020 Authors of Arktos.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package create

import (
	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/generate"
	generateversioned "k8s.io/kubernetes/pkg/kubectl/generate/versioned"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
	"k8s.io/kubernetes/pkg/kubectl/util/templates"
)

var (
	tenantLong = templates.LongDesc(i18n.T(`
		Create a tenant with the specified name.`))

	tenantExample = templates.Examples(i18n.T(`
	  # Create a new tenant named my-tenant
	  kubectl create tenant my-tenant`))
)

// TenantOpts is the options for 'create namespace' sub command
type TenantOpts struct {
	CreateSubcommandOptions *CreateSubcommandOptions
}

// NewCmdCreateTenant is a macro command to create a new tenant
func NewCmdCreateTenant(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	options := &TenantOpts{
		CreateSubcommandOptions: NewCreateSubcommandOptions(ioStreams),
	}

	cmd := &cobra.Command{
		Use:                   "tenant NAME [--dry-run]",
		DisableFlagsInUseLine: true,
		Aliases:               []string{"te"},
		Short:                 i18n.T("Create a tenant with the specified name"),
		Long:                  tenantLong,
		Example:               tenantExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(options.Complete(f, cmd, args))
			cmdutil.CheckErr(options.Run())
		},
	}

	options.CreateSubcommandOptions.PrintFlags.AddFlags(cmd)

	cmdutil.AddApplyAnnotationFlags(cmd)
	cmdutil.AddValidateFlags(cmd)
	cmdutil.AddGeneratorFlags(cmd, generateversioned.TenantV1GeneratorName)

	return cmd
}

// Complete completes all the required options
func (o *TenantOpts) Complete(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	name, err := NameFromCommandArgs(cmd, args)
	if err != nil {
		return err
	}

	var generator generate.StructuredGenerator
	switch generatorName := cmdutil.GetFlagString(cmd, "generator"); generatorName {
	case generateversioned.TenantV1GeneratorName:
		generator = &generateversioned.TenantGeneratorV1{Name: name}
	default:
		return errUnsupportedGenerator(cmd, generatorName)
	}

	return o.CreateSubcommandOptions.Complete(f, cmd, args, generator)
}

// Run calls the CreateSubcommandOptions.Run in TenantOpts instance
func (o *TenantOpts) Run() error {
	return o.CreateSubcommandOptions.Run()
}
