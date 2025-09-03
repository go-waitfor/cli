#!/bin/bash

# Copyright Tim Voronov 2020
version=$(curl -sI https://github.com/go-waitfor/cli/releases/latest | grep Location | awk -F"/" '{ printf "%s", $NF }' | tr -d '\r')
if [ ! "$version" ]; then
    echo "Failed while attempting to install waitfor. Please manually install:"
    echo ""
    echo "1. Open your web browser and go to https://github.com/go-waitfor/cli/releases"
    echo "2. Download the latest release for your platform."
    echo "3. chmod +x ./waitfor"
    echo "4. mv ./waitfor /usr/local/bin"
    exit 1
fi

hasCli() {
    if which waitfor >/dev/null 2>&1; then
        echo
        echo "You already have waitfor!"
        export n=3
        echo "Overwriting in $n seconds.. Press Control+C to cancel."
        echo
        sleep $n
    fi

    if ! which curl >/dev/null 2>&1; then
        echo "You need curl to use this script."
        exit 1
    fi

    if ! which tar >/dev/null 2>&1; then
        echo "You need tar to use this script."
        exit 1
    fi
}

checkHash(){
    sha_cmd="sha256sum"

    if [ ! -x "$(command -v "$sha_cmd")" ]; then
        sha_cmd="shasum -a 256"
    fi

    if [ -x "$(command -v "$sha_cmd")" ]; then
        if ! (cd "$targetDir" && curl -sSL "$baseUrl"/checksums.txt | "$sha_cmd" -c >/dev/null); then
            echo "Binary checksum didn't match. Exiting"
            exit 1
        fi
    fi
}

getPackage() {
    uname=$(uname)
    userid=$(id -u)

    platform=""
    case $uname in
    "Darwin")
    platform="_darwin"
    ;;
    "Linux")
    platform="_linux"
    ;;
    esac

    uname_arch=$(uname -m)
    arch=""
    case $uname_arch in
    "x86_64")
        arch="_amd64"
        ;;
    "aarch64"|"arm64")
        arch="_arm64"
        ;;
    *)
        echo "Architecture $uname_arch is not supported. Exiting"
        exit 1
        ;;
    esac

    suffix=$platform$arch
    targetDir="/tmp/waitfor$suffix"

    if [ "$userid" != "0" ]; then
        targetDir="$(pwd)/waitfor$suffix"
    fi

    if [ ! -d "$targetDir" ]; then
        mkdir "$targetDir"
    fi

    targetFile="$targetDir/waitfor"

    if [ -e "$targetFile" ]; then
        rm "$targetFile"
    fi

    baseUrl=https://github.com/go-waitfor/cli/releases/download/$version
    url=$baseUrl/waitfor$suffix.tar.gz
    echo "Downloading package $url as $targetFile"

    if ! curl -sSL "$url" | tar xz -C "$targetDir"; then
        echo "Failed to download or extract package"
        exit 1
    fi

    checkHash

    chmod +x "$targetFile"

    echo "Download complete."

    if [ "$userid" != "0" ]; then
        echo
        echo "========================================================="
        echo "==    As the script was run as a non-root user the     =="
        echo "==    following commands may need to be run manually   =="
        echo "========================================================="
        echo
        echo "  sudo cp \"$targetFile\" /usr/local/bin/waitfor"
        echo "  rm -rf \"$targetDir\""
        echo
    else
        echo
        echo "Running as root - Attempting to move \"$targetFile\" to /usr/local/bin"

        if mv "$targetFile" /usr/local/bin/waitfor; then
            echo "New version of waitfor installed to /usr/local/bin"
        fi

        if [ -d "$targetDir" ]; then
            rm -rf "$targetDir"
        fi

        waitfor version
    fi
}

hasCli
getPackage