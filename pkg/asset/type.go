// Copyright 2022 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package asset

import (
	"github.com/google/syzkaller/dashboard/dashapi"
	"github.com/google/syzkaller/sys/targets"
)

type TypeDescription struct {
	AllowMultiple     bool
	GetTitle          QueryTypeTitle
	ContentType       string
	ReportingPrio     int // the smaller, the higher the asset is on the list during reporting
	NoReporting       bool
	customCompressor  Compressor
	preserveExtension bool
}

var assetTypes = map[dashapi.AssetType]*TypeDescription{
	dashapi.BootableDisk: {
		GetTitle:      constTitle("disk image"),
		ReportingPrio: 1,
	},
	dashapi.NonBootableDisk: {
		GetTitle:      constTitle("disk image (non-bootable)"),
		ReportingPrio: 2,
	},
	dashapi.KernelObject: {
		GetTitle: func(target *targets.Target) string {
			if target != nil && target.KernelObject != "" {
				return target.KernelObject
			}
			return "kernel object"
		},
		ReportingPrio: 3,
	},
	dashapi.KernelImage: {
		GetTitle:      constTitle("kernel image"),
		ReportingPrio: 4,
	},
	dashapi.HTMLCoverageReport: {
		GetTitle:          constTitle("coverage report(html)"),
		AllowMultiple:     true,
		ContentType:       "text/html",
		NoReporting:       true,
		customCompressor:  gzipCompressor,
		preserveExtension: true,
	},
}

type QueryTypeTitle func(*targets.Target) string

func constTitle(title string) QueryTypeTitle {
	return func(*targets.Target) string {
		return title
	}
}

func GetTypeDescription(assetType dashapi.AssetType) *TypeDescription {
	return assetTypes[assetType]
}
