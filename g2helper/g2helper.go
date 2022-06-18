/*
Package helper ...
*/
package g2helper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getDatabaseUrl() (string, error) {
	result := ""
	databaseUrl := os.Getenv("XYZZY_DATABASE_URL")
	parsedUrl, err := url.Parse(databaseUrl)

	switch parsedUrl.Scheme {
	case "db2":
		result = "FIXME: Not implemented"
	case "mssql":
		result = "FIXME: Not implemented"
	case "mysql":
		result = "FIXME: Not implemented"
	case "postgresql":
		result = fmt.Sprintf(
			"%s://%s@%s:%s/",
			parsedUrl.Scheme,
			parsedUrl.User,
			parsedUrl.Host,
			string(parsedUrl.Path[1:]),
		)
	case "sqlite3":
		result = "FIXME: Not implemented"
	default:
		result = ""
	}

	return result, err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func GetSimpleSystemConfigurationJson(ctx context.Context) (string, error) {
	var err error = nil

	databaseUrl, databaseUrlErr := getDatabaseUrl()
	if databaseUrlErr != nil {
		return "", databaseUrlErr
	}

	resultStruct := XyzzyConfiguration{
		Pipeline: XyzzyConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		Sql: XyzzyConfigurationSql{
			Connection: databaseUrl,
		},
	}

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes), err
}
