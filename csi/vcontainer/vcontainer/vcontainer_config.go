package vcontainer

import (
	"github.com/vngcloud/vcontainer-storage-interface/csi/client"
	"github.com/vngcloud/vcontainer-storage-interface/csi/utils/metadata"
)

type (
	Config struct {
		Global       client.AuthOpts
		Metadata     metadata.Opts
		BlockStorage BlockStorageOpts
	}
	BlockStorageOpts struct {
		NodeVolumeAttachLimit    int64 `gcfg:"node-volume-attach-limit"`
		RescanOnResize           bool  `gcfg:"rescan-on-resize"`
		IgnoreVolumeAZ           bool  `gcfg:"ignore-volume-az"` // ignore volume AZ (available zones) when attaching to node
		IgnoreVolumeMicroversion bool  `gcfg:"ignore-volume-microversion"`
	}
)