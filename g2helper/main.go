// The G2helper Package is a set of method to help with common tasks.
//
// The purpose of a g2helper object is:
//   • ...
//   • ...
//   • ...
package g2helper

import ()

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6015%04d"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type SenzingConfigurationPipeline struct {
	ConfigPath   string `json:"CONFIGPATH"`
	ResourcePath string `json:"RESOURCEPATH"`
	SupportPath  string `json:"SUPPORTPATH"`
}

type SenzingConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type SenzingConfiguration struct {
	Pipeline SenzingConfigurationPipeline `json:"PIPELINE"`
	Sql      SenzingConfigurationSql      `json:"SQL"`
}

type G2helperImpl struct{}
