#!/bin/bash

set -e 

echo "$GOOGLE_SA_KEY" > sa.json
export GOOGLE_APPLICATION_CREDENTIALS="sa.json"

go run .