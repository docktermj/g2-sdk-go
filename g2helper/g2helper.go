/*
Package helper ...
*/
package g2helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"

	errormsg "github.com/docktermj/go-json-log-message/message"
	//	"github.com/docktermj/go-logger/logger"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

var MessageLevelMap = map[int]string{
	1000:  "info",
	2000:  "warning",
	3000:  "error",
	4000:  "debug",
	5000:  "trace",
	6000:  "retryable",
	9000:  "reserved",
	10000: "fatal",
}

var SenzingErrorsMap = map[string]string{
	"0002E":  "error",
	"0007E":  "error",
	"0023E":  "error",
	"0024E":  "error",
	"0025E":  "error",
	"0026E":  "error",
	"0027E":  "error",
	"0032E":  "error",
	"0034E":  "error",
	"0035E":  "error",
	"0036E":  "error",
	"0048E":  "fatal",
	"0051E":  "error",
	"0053E":  "fatal",
	"0054E":  "error",
	"0061E":  "error",
	"0062E":  "error",
	"0064E":  "error",
	"1007E":  "error",
	"2134E":  "error",
	"30020":  "error",
	"30103E": "error",
	"30110E": "error",
	"30111E": "error",
	"30112E": "error",
	"30121E": "error",
	"30122E": "error",
	"30123E": "error",
	"9000E":  "error",
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getOsEnv(variableName string) (string, error) {
	var err error = nil
	result, isSet := os.LookupEnv(variableName)
	if !isSet {
		err = BuildError(MessageIdFormat, 2990, "Environment variable not set.", variableName)
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

// Build an error.
func BuildError(idTemplate string, errorNumber int, message string, details ...string) error {
	errorMessage := errormsg.BuildMessage(
		BuildMessageId(idTemplate, errorNumber),
		GetMessageLevel(errorNumber, message),
		message,
		details...,
	)
	return errors.New(errorMessage)
}

// Construct a unique message id.
func BuildMessageId(idTemplate string, errorNumber int) string {
	return fmt.Sprintf(idTemplate, errorNumber)
}

// Get a Senzing configuration for a "system install" with single database.
func BuildSimpleSystemConfigurationJson() (string, error) {
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

// Based on the errorNumber and Senzing error code, get the message level.
func GetMessageLevel(errorNumber int, message string) string {

	var result = "unknown"

	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(MessageLevelMap))
	for key := range MessageLevelMap {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Ints(messageLevelKeys)

	// Using the sorted message number, find the level.

	for _, messageLevelKey := range messageLevelKeys {
		if errorNumber < messageLevelKey {
			result = MessageLevelMap[messageLevelKey]
			break
		}
	}

	// Determine the level of a specific Senzing error.

	messageSplits := strings.Split(message, "|")
	for key, value := range SenzingErrorsMap {
		if messageSplits[0] == key {
			result = value
			break
		}
	}

	// Determine if message has error code.

	return result
}

// Based on the errorNumber and Senzing error code, get the message level.
func GetMessageLevelFromError(err error) string {

	// FIXME:

	return ""

}

// Inspect the error to see what the level is and log based on the level.
func LogMessageBasedOnLevel(err error) error {
	var result error = nil

	// FIXME:

	return result

}
