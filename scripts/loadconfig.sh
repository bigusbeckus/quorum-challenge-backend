#!/usr/bin/env bash

if [ -z $CONFIG_YAML_B64 ]; then
  echo Config not found in environment
	exit 1
fi

echo $CONFIG_YAML_B64 | base64 -d >internal/pkg/config/config.yaml
