#!/bin/bash

mkdir -p "output"
go get github.com/anaskhan96/soup
go get golang.org/x/exp

platforms=(
  # NOT SUPPORTED "darwin/amd64"
  # NOT SUPPORTED "darwin/arm64"
  "linux/amd64"
  # NOT SUPPORTED "linux/arm"
  # NOT SUPPORTED "linux/arm64"
  # NOT SUPPORTED "windows/amd64"
  # NOT SUPPORTED "windows/arm"
)

for platform in "${platforms[@]}"; do
  GOOS="${platform%/*}"
  GOARCH="${platform#*/}"
  
  executableName=""
  fixedArch=""
  fixedOs=""

  if [ "$GOARCH" = "amd64" ]; then
    fixedArch="x86_64";
  elif [ "$GOARCH" = "arm" ]; then
    fixedArch="arm-v7a"
  elif [ "$GOARCH" = "arm64" ]; then
    fixedArch="arm64-v8a"
  fi

  if [ "$GOOS" = "windows" ]; then
    fixedOs="Windows"
  elif [ "$GOOS" = "darwin" ]; then
    fixedOs="macOS"
  elif [ "$GOOS" = "linux" ]; then
    fixedOs="Linux"
  fi

  if [ "$GOOS" = "windows" ]; then
    executableName="$fixedOs-$fixedArch.exe"
  else
    executableName="$fixedOs-$fixedArch"
  fi

  if [ "$GOARCH" = "arm" ] && [ ! "$GOOS" = "linux" ]; then
    continue
  fi

  echo "Building $fixedOs-$fixedArch..."
  GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -o "output/GoTgCalls-$executableName" &
done
wait
echo "Done building all executables."