#!/usr/bin/env bash
USAGE="\
$0 is a convenience script for managing Docker Compose-based deployment.

Usage:
    $0 [OPTION] ...

Options:
    -u Run all containers
    -p Update the frontend container
    -d Terminate all containers
    -z Terminate all containers and its volumes. Dangerous!
"

# Exit immediately on any failure
set -e
set -o pipefail

function help() {
    echo "$USAGE" >&2
    exit 1
}

function _check_docker_compose() {
    printf "Checking docker-compose.yml..."
    if [[ ! -f ./docker-compose.yml ]]; then
        echo "Error: docker-compose.yml is missing." >&2
        echo "The file must be placed at the same level as $0" >&2
        exit 1
    fi
    echo "OK. ✔️"

    #printf "Checking docker-compose.override.yml..."
    #if [[ ! -f ./docker-compose.override.yml ]]; then
    #    echo "Error: docker-compose.override.yml is missing." >&2
    #    echo "The file must be placed at the same level as $0" >&2
    #    exit 1
    #fi
    #echo "OK. ✔️"
}

function run() {
    _check_docker_compose

    echo "Running all containers..."
    docker-compose up --detach
    echo "OK. ✔️"
}

function update() {
    # Possible values: develop, production
    local _IMAGE_TAG=${BACKEND_IMAGE_TAG:-develop}

    echo "$_IMAGE_TAG"
    echo "Pulling latest image if available..."
    docker pull "registry.pusilkom.com/bapenda/bapenda-pkb-bbnkb-backend:$_IMAGE_TAG"
    echo "OK. ✔️"

    _check_docker_compose

    echo "Replacing frontend container..."
    docker-compose stop backend
    docker-compose rm --force backend
    docker-compose up --detach backend
    echo "OK. ✔️"
}

function terminate() {
    _check_docker_compose

    echo "Stopping all containers..."
    docker-compose down
    echo "OK. ✔️"
}

function destroy() {
    _check_docker_compose

    echo "Stopping all containers and volumes..."
    docker-compose down --volumes
    echo "OK. ✔️"
}

while getopts ":updz" option; do
    case "$option" in
        u) run ;;
        p) update ;;
        d) terminate ;;
        z) destroy ;;
        ?)
            echo "Error: invalid option"
            help ;;
    esac
done

exit 0
