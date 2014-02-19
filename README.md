logshifter [![Build Status](https://travis-ci.org/ironcladlou/logshifter.png?branch=master)](https://travis-ci.org/ironcladlou/logshifter)
=====

A simple log pipe designed to maintain consistently high input consumption rates, preferring to
drop old messages rather than block the input producer.

The primary design goals are:

* Minimal blocking of the input producer
* Asynchronous dispatch of log events to an output
* Sacrifice delivery rate for input processing consistency
* Fixed maximum memory use as a factor of configurable buffer sizes
* Pluggable and configurable outputs

logshifter is useful for capturing and redirecting output of heterogenous applications which
emit logs to stdout rather than to a downstream aggregator (e.g. syslog) directly.


Configuration
---
An optional configuration file can be supplied to logshifter with the `-config` flag. The file
is a list of key value pairs in the format `k = v`, one per line. The possible options and their
defaults are described in the [Config type](lib/config.go). Keys are case insensitive.


Statistics
---
Periodic stats can be emitted to a file using the `-statsfilename` argument.  The stats are written
in JSON format on an interval specified by `-statsinterval` (which is a string in a format expected
by the golang [time.ParseDuration](http://golang.org/pkg/time/#ParseDuration) function).

Note that enabling statistics will introduce extra processing overhead.
