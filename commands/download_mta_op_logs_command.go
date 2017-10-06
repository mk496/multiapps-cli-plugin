package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/SAP/cf-mta-plugin/log"
	"github.com/SAP/cf-mta-plugin/ui"
)

const (
	defaultDownloadDirPrefix string = "mta-op-"
)

// DownloadMtaOperationLogsCommand is a command for retrieving the logs of an MTA operation
type DownloadMtaOperationLogsCommand struct {
	BaseCommand
}

// GetPluginCommand returns the plugin command details
func (c *DownloadMtaOperationLogsCommand) GetPluginCommand() plugin.Command {
	return plugin.Command{
		Name:     "download-mta-op-logs",
		Alias:    "dmol",
		HelpText: "Download logs of multi-target app operation",
		UsageDetails: plugin.Usage{
			Usage: "cf download-mta-op-logs -i OPERATION_ID [-d DIRECTORY] [-u URL]",
			Options: map[string]string{
				"i": "Operation id",
				"d": "Directory to download logs, by default '" + defaultDownloadDirPrefix + "<OPERATION_ID>/'",
				"u": "Deploy service URL, by default 'deploy-service.<system-domain>'",
			},
		},
	}
}

// Execute executes the command
func (c *DownloadMtaOperationLogsCommand) Execute(args []string) ExecutionStatus {
	log.Tracef("Executing command '"+c.name+"': args: '%v'\n", args)

	var host string
	var operationID string
	var downloadDirName string

	// Parse command arguments and check for required options
	flags, err := c.CreateFlags(&host)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}
	flags.StringVar(&operationID, "i", "", "")
	flags.StringVar(&downloadDirName, "d", "", "")
	err = c.ParseFlags(args, nil, flags, map[string]bool{"i": true})
	if err != nil {
		c.Usage(err.Error())
		return Failure
	}

	// Set the download directory if not specified
	if downloadDirName == "" {
		downloadDirName = defaultDownloadDirPrefix + operationID + "/"
	}

	context, err := c.GetContext()
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	// Print initial message
	ui.Say("Downloading logs of multi-target app operation with id %s in org %s / space %s as %s...",
		terminal.EntityNameColor(operationID), terminal.EntityNameColor(context.Org),
		terminal.EntityNameColor(context.Space), terminal.EntityNameColor(context.Username))

	// Create new SLMP client
	slmpClient, err := c.NewSlmpClient(host)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	// Check SLMP metadata
	err = CheckSlmpMetadata(slmpClient)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	// Get the service ID for the specified operation ID
	serviceID, err := GetServiceID(operationID, slmpClient)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	// Create new SLPP client
	slppClient, err := c.NewSlppClient(host, serviceID, operationID)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	// Check SLPP metadata
	err = CheckSlppMetadata(slppClient)
	if err != nil {
		ui.Failed(err.Error())
		return Failure
	}

	// Download all logs
	downloadedLogs := make(map[string]*string)
	logs, err := slppClient.GetLogs()
	if err != nil {
		ui.Failed("Could not get process logs: %s", err)
		return Failure
	}
	for _, logx := range logs.Logs {
		content, err := slppClient.GetLogContent(*logx.ID)
		if err != nil {
			ui.Failed("Could not get content of log %s: %s", terminal.EntityNameColor(*logx.ID), err)
			return Failure
		}
		downloadedLogs[*logx.ID] = &content
	}
	ui.Ok()

	// Create the download directory
	downloadDir, err := createDownloadDirectory(downloadDirName)
	if err != nil {
		ui.Failed("Could not create download directory %s: %s", terminal.EntityNameColor(downloadDirName), err)
		return Failure
	}

	// Get all logs and save their contents to the download directory
	ui.Say("Saving logs to %s...", terminal.EntityNameColor(downloadDir))
	for logID, content := range downloadedLogs {
		err = saveLogContent(downloadDir, logID, content)
		if err != nil {
			ui.Failed("Could not save log %s: %s", terminal.EntityNameColor(logID), err)
			return Failure
		}
	}
	ui.Ok()
	return Success
}

func createDownloadDirectory(downloadDirName string) (string, error) {
	// Check if directory name ends with the os specific path separator
	if !strings.HasSuffix(downloadDirName, string(os.PathSeparator)) {
		//If there is no os specific path separator, put it at the end of the direcotry name
		downloadDirName = downloadDirName + string(os.PathSeparator)
	}

	// Check if the directory already exists
	if stat, _ := os.Stat(downloadDirName); stat != nil {
		return "", fmt.Errorf("File or directory already exists.")
	}

	// Create the directory
	err := os.MkdirAll(downloadDirName, 0755)
	if err != nil {
		return "", nil
	}

	// Return the absolute path of the directory
	return filepath.Abs(filepath.Dir(downloadDirName))
}

func saveLogContent(downloadDir, logID string, content *string) error {
	ui.Say("  %s", logID)
	return ioutil.WriteFile(downloadDir+"/"+logID, []byte(*content), 0644)
}