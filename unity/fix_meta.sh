#!/bin/bash

./tools/unity-meta-check-darwin-amd64 ./gen/csharp | ./tools/unity-meta-autofix-darwin-amd64 -root-dir ./gen/csharp .