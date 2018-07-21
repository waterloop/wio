package config

type meta struct {
    Name                 string
    Version              string
    EnableBashCompletion bool
    Copyright            string
    UsageText            string
}

var ProjectMeta = meta{
    Name:                 "wio",
    Version:              "0.4.0",
    EnableBashCompletion: true,
    Copyright:            "Copyright (c) 2018 Waterloop",
    UsageText:            "An Iot development environment",
}
