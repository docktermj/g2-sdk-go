/*
Package helper ...
*/
package g2helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getOsEnv(variableName string) (string, error) {
	var err error = nil
	result, isSet := os.LookupEnv(variableName)
	if !isSet {
		err = errors.New(fmt.Sprintf(
			"%s - %s environment variable not set",
			fmt.Sprintf(MessageIdFormat, 1111),
			variableName))
	}
	return result, err
}

func getDatabaseUrl() (string, error) {
	result := ""
	databaseUrl, err := getOsEnv("XYZZY_DATABASE_URL")
	if err != nil {
		return result, err
	}

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
