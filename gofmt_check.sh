#!/bin/bash

! gofmt -l . | grep . | golangci-lint run --timeout 5m0s
