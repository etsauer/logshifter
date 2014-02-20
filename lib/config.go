package lib

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
	QueueSize             int    // size of the internal log message queue
	InputBufferSize       int    // input up to \n or this number of bytes is considered a line
	OutputType            string // one of syslog, file
	SyslogBufferSize      int    // lines bound for syslog lines are split at this size
	FileBufferSize        int    // lines bound for a file are split at this size
	OutputTypeFromEnviron bool   // allows outputtype to be overridden via LOGSHIFTER_OUTPUT_TYPE
}

func ParseConfig(file string) (*Config, error) {
	config := &Config{
		QueueSize:             100,
		InputBufferSize:       2048,
		OutputType:            "syslog",
		SyslogBufferSize:      2048,
		FileBufferSize:        2048,
		OutputTypeFromEnviron: true,
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
			config.QueueSize, _ = strconv.Atoi(v)
		case "inputbuffersize":
			config.InputBufferSize, _ = strconv.Atoi(v)
		case "outputtype":
			switch v {
			case "syslog":
				config.OutputType = Syslog
			case "file":
				config.OutputType = File
			}
		case "syslogbuffersize":
			config.SyslogBufferSize, _ = strconv.Atoi(v)
		case "filebuffersize":
			config.FileBufferSize, _ = strconv.Atoi(v)
		case "outputtypefromenviron":
			config.OutputTypeFromEnviron, _ = strconv.ParseBool(v)
		}
	}

	return config, nil
}
