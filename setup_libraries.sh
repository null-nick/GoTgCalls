#!/bin/bash

VERSION="v1.2.3"
ARCHITECTURE=$(uname -m)
DEBUG_MODE=false

parse_arguments() {
  for arg in "$@"; do
    case $arg in
      --debug)
        DEBUG_MODE=true
        shift
        ;;
      *)
        ;;
    esac
  done
}

download_library() {
  local url=$1
  local output_file=$2

  echo -e "\e[34mDownloading library from: $url\e[0m"
  wget -q -O "$output_file" "$url"

  # shellcheck disable=SC2181
  if [ $? -eq 0 ]; then
    echo -e "\e[32mDownload successful: $output_file\e[0m"
  else
    echo -e "\e[31m[ERROR] Download failed!\e[0m"
    exit 1
  fi
}

unzip_library() {
  local file=$1
  local destination=$2

  echo -e "\e[34mUnzipping $file to $destination\e[0m"
  unzip -q "$file" -d "$destination"

  # shellcheck disable=SC2181
  if [ $? -eq 0 ]; then
    echo -e "\e[32mUnzip successful!\e[0m"
  else
    echo -e "\e[31m[ERROR] Unzip failed!\e[0m"
    exit 1
  fi
}

clear_workspace() {
  local file=$1
  local destination=$2
  if [ "$DEBUG_MODE" == true ]; then
    echo -e "\e[34m[DEBUG] Clearing workspace: Removing $file and $destination\e[0m"
  fi
  rm -rf "$file" "$destination"
}

copy_files() {
  local source_folder=$1

  local header_file="$source_folder/include/ntgcalls.h"
  if [ -f "$header_file" ]; then
    cp -r "$header_file" "ntgcalls/bindings/"
    echo -e "\e[32mHeader file copied to ntgcalls/bindings/ntgcalls.h\e[0m"
  else
    echo -e "\e[31m[ERROR] Header file ntgcalls.h not found in $source_folder/include/\e[0m"
    exit 1
  fi

  local library_file="$source_folder/libntgcalls.so"
  if [ -f "$library_file" ]; then
    cp -r "$library_file" "ntgcalls/bindings/lib/"
    echo -e "\e[32mLibrary file copied to ntgcalls/bindings/lib/libntgcalls.so\e[0m"
  else
    echo -e "\e[31m[ERROR] Library file libntgcalls.so not found in $source_folder/\e[0m"
    exit 1
  fi
}

parse_arguments "$@"

OS=$(uname -s)
echo -e "\e[34mDetected OS: $OS ($ARCHITECTURE)\e[0m"

case $OS in
  Linux)
    if [ "$ARCHITECTURE" == "x86_64" ]; then
      URL="https://github.com/pytgcalls/ntgcalls/releases/download/$VERSION/ntgcalls.linux-x86_64-shared_libs.zip"
      FILE_NAME="ntgcalls.linux-x86_64-shared_libs.zip"
#   elif [ "$ARCHITECTURE" == "aarch64" ]; then
#     URL="https://github.com/pytgcalls/ntgcalls/releases/download/$VERSION/ntgcalls.linux-arm64-shared_libs.zip"
#     FILE_NAME="ntgcalls.linux-arm64-shared_libs.zip"
    else
      echo -e "\e[31m[ERROR] Unsupported architecture: $ARCHITECTURE\e[0m"
      exit 1
    fi
    ;;
#  Darwin)
#    if [ "$ARCHITECTURE" == "arm64" ]; then
#      URL="https://github.com/pytgcalls/ntgcalls/releases/download/$VERSION/ntgcalls.macos-arm64-shared_libs.zip"
#      FILE_NAME="ntgcalls.macos-arm64-shared_libs.zip"
#    else
#      echo -e "\e[31m[ERROR] Unsupported architecture: $ARCHITECTURE\e[0m"
#      exit 1
#    fi
#    ;;
#  MINGW*|MSYS*)
#    URL="https://github.com/pytgcalls/ntgcalls/releases/download/$VERSION/ntgcalls.windows-x86_64-shared_libs.zip"
#    FILE_NAME="ntgcalls.windows-x86_64-shared_libs.zip"
#    ;;
  *)
    echo -e "\e[31m[ERROR] Unsupported OS: $OS\e[0m"
    exit 1
    ;;
esac

echo -e "\e[34mEnter the folder to use for extraction ('release' or 'debug'). Default is 'release' if no input is provided:\e[0m"
read -r FOLDER_CHOICE

if [ -z "$FOLDER_CHOICE" ]; then
  FOLDER_CHOICE="release"
fi

EXTRACT_FOLDER="$FOLDER_CHOICE"
echo -e "\e[32mUsing the '$EXTRACT_FOLDER' folder for extraction.\e[0m"

DEST_DIR="./output_folder"
clear_workspace "$FILE_NAME" "$DEST_DIR"

download_library "$URL" "$FILE_NAME"
unzip_library "$FILE_NAME" "$DEST_DIR"
copy_files "$DEST_DIR/$EXTRACT_FOLDER"

echo -e "\e[32mSetup completed successfully!\e[0m"
clear_workspace "$FILE_NAME" "$DEST_DIR"
