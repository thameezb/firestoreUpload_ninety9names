#!/bin/bash

set -e 

echo "$FIREBASE_SA_KEY" > sa.json
export GOOGLE_APPLICATION_CREDENTIALS="sa.json"

go run .