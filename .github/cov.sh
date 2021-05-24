#!/bin/bash

set -e

./.github/test.sh

go tool cover -html=coverage.txt