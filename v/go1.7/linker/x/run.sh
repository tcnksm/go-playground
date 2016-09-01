#!/bin/bash

go run -ldflags="-X main.Commit=GIT_COMMIT" main.go
