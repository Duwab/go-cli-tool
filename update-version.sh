#!/bin/bash
VERSION=$(git tag)
BUILD_TIME=$(date)

sed -i "s/BuildVersion string = \"[^\"]*\"/BuildVersion string = \"${VERSION}\"/" version.go
sed -i "s/BuildTime    string = \"[^\"]*\"/BuildTime    string = \"${BUILD_TIME}\"/" version.go
