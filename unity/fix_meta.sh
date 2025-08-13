#!/bin/bash

./tools/unity-meta-check ./gen/csharp | ./tools/unity-meta-autofix -root-dir ./gen/csharp .

./tools/unity-meta-check ./gen/csharp/GRPC/ | ./tools/unity-meta-autofix -root-dir ./gen/csharp/GRPC/ .