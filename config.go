package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	// output types
	Syslog = "syslog"
	File   = "file"

	DefaultConfigFile = "/etc/openshift/logshifter.conf"
)

type Config struct {
	queueSize             int    // size of the internal log message queue
	inputBufferSize       int    // input up to \n or this number of bytes is considered a line
	outputType            string // one of syslog, file
	syslogBufferSize      int    // lines bound for syslog lines are split at this size
	fileBufferSize        int    // lines bound for a file are split at this size
	fileWriterDir         string // base dir for the file writer output's file
	outputTypeFromEnviron bool   // allows outputtype to be overridden via LOGSHIFTER_OUTPUT_TYPE
}

func ParseConfig(file string) (*Config, error) {
	config := &Config{
		queueSize:             100,
		inputBufferSize:       2048,
		outputType:            "syslog",
		syslogBufferSize:      2048,
		fileBufferSize:        2048,
		outputTypeFromEnviron: true,
	}

	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || len(line) == 0 {
			break
		}

		c := strings.SplitN(line, "=", 2)

		if len(c) != 2 {
			break
		}

		k := strings.Trim(c[0], "\n ")
		v := strings.Trim(c[1], "\n ")

		switch strings.ToLower(k) {
		case "queuesize":
			config.queueSize, _ = strconv.Atoi(v)
		case "inputbuffersize":
			config.inputBufferSize, _ = strconv.Atoi(v)
		case "outputtype":
			switch v {
			case "syslog":
				config.outputType = Syslog
			case "file":
				config.outputType = File
			}
		case "syslogbuffersize":
			config.syslogBufferSize, _ = strconv.Atoi(v)
		case "filebuffersize":
			config.fileBufferSize, _ = strconv.Atoi(v)
		case "outputtypefromenviron":
			config.outputTypeFromEnviron, _ = strconv.ParseBool(v)
		case "filewriterdir":
			config.fileWriterDir = v
		}
	}

	return config, nil
}
