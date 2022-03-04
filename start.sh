#!/usr/bin/env bash

set -euo pipefail

main() {
    source ./load_dotenv.sh && netlify dev
}

main

