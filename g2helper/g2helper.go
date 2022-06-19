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
	"sort"
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

func GetErrorLevel(errorNumber int) string {

	// Create a map of the different levels. Map will be unsorted.

	errorLevelsMap := map[int]string{
		1000:  "I", // Informational
		2000:  "W", // Warning
		3000:  "E", // Error
		4000:  "D", // Debug
		5000:  "T", // Trace
		9000:  "R", // Reserved
		10000: "F", // Fatal
	}

	// Create a list of sorted keys.

	errorLevelsKeys := make([]int, 0, len(errorLevelsMap))
	for key := range errorLevelsMap {
		errorLevelsKeys = append(errorLevelsKeys, key)
	}
	sort.Ints(errorLevelsKeys)

	// Using the sorted key, find the level.

	for _, errorLevelsKey := range errorLevelsKeys {
		if errorNumber < errorLevelsKey {
			return errorLevelsMap[errorLevelsKey]
		}
	}
	return "" // Unknown
}

func GetMessageId(errorNumber int) string {
	return fmt.Sprintf(
		"%s%s",
		fmt.Sprintf(MessageIdFormat, errorNumber),
		GetErrorLevel(errorNumber))
}

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
