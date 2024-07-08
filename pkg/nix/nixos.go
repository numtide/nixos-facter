package nix

import (
	"embed"
	"fmt"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"io"
	"slices"
	"text/template"
)

var (
	//go:embed templates
	templatesFS embed.FS
)

type ModuleGenerator struct {
	Report *hwinfo.Report

	Attrs   []string
	Imports []string

	KernelModules                []string
	ModulePackages               []string
	InitrdKernelModules          []string
	InitrdAvailableKernelModules []string

	NetworkingInterfaces []string
}

func (mg *ModuleGenerator) Generate(writer io.Writer) error {

	for _, item := range mg.Report.Items {
		pci(mg, item)
		usb(mg, item)
		networking(mg, item)
	}

	// sort first
	slices.Sort(mg.Attrs)
	slices.Sort(mg.Imports)
	slices.Sort(mg.KernelModules)
	slices.Sort(mg.ModulePackages)
	slices.Sort(mg.InitrdKernelModules)
	slices.Sort(mg.InitrdAvailableKernelModules)

	// remove duplicates
	mg.Attrs = slices.Compact(mg.Attrs)
	mg.Imports = slices.Compact(mg.Imports)
	mg.KernelModules = slices.Compact(mg.KernelModules)
	mg.ModulePackages = slices.Compact(mg.ModulePackages)
	mg.InitrdKernelModules = slices.Compact(mg.InitrdKernelModules)
	mg.InitrdAvailableKernelModules = slices.Compact(mg.InitrdAvailableKernelModules)

	funcMap := template.FuncMap{
		"nixList":       ToNixList,
		"nixStringList": ToNixStringList,
		"multiLineList": MultiLineList,
	}

	tmpl, err := template.
		New("default.tmpl").
		Funcs(funcMap).
		ParseFS(templatesFS, "**/*.tmpl")

	if err != nil {
		return fmt.Errorf("failed to load default template: %w", err)
	}

	return tmpl.Funcs(funcMap).Execute(writer, mg)
}
