#!/bin/bash

# Navigate to the directory
cd ~/my-react-app/src

# Files to keep (space-separated list)
KEEP_FILES=("App.js" "App.css" "index.js" "index.css" "reportWebVitals.js")

# Convert the array to a pattern
KEEP_PATTERN=$(printf "|%s" "${KEEP_FILES[@]}")
KEEP_PATTERN=${KEEP_PATTERN:1}

# Find and delete files not in the keep list
find "$DIR" -type f ! -name "$KEEP_PATTERN" -exec rm -f {} +

echo "Cleanup complete. Kept files: ${KEEP_FILES[*]}"
