package global

import (
	"time"
)

var (
	Version               = "0.8.6"
	Commit                = ""
	BuildTime             = "none"
	Debug                 = false
	ProfilingEnabled      = false
	StartupTimeout        = 5 * time.Minute
	ConfigFilename        = "cog.yaml"
	ReplicateRegistryHost = "r8.im"
	ReplicateWebsiteHost  = "replicate.com"
	LabelNamespace        = "run.cog."
)
