// The G2helper Package is a set of method to help with common tasks.
//
// The purpose of a g2helper object is:
//   • ...
//   • ...
//   • ...
package g2helper

import ()

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type XyzzyConfigurationPipeline struct {
	ConfigPath   string `json:"CONFIGPATH"`
	ResourcePath string `json:"RESOURCEPATH"`
	SupportPath  string `json:"SUPPORTPATH"`
}

type XyzzyConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type XyzzyConfiguration struct {
	Pipeline XyzzyConfigurationPipeline `json:"PIPELINE"`
	Sql      XyzzyConfigurationSql      `json:"SQL"`
}

type G2helperImpl struct{}
