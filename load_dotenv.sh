#!/usr/bin/env bash

main() {
    export $(sed 's/#.*//g' <.env | xargs)
}

main