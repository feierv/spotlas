#!/bin/bash

# Define variables
DATABASE_NAME="initial_database"
FUNCTION_FILE="$1"
OUTPUT_DIR="output_task_1_query"
OUTPUT_FILE="$OUTPUT_DIR/output_task_1_query.txt"

# Check if a function file is provided as an argument
if [ -z "$FUNCTION_FILE" ]; then
    echo "Usage: $0 <function_file>"
    exit 1
fi

# Create the output directory if it doesn't exist
if [ ! -d $OUTPUT_DIR ]; then
    mkdir -p $OUTPUT_DIR
fi

# Connect to PostgreSQL and execute the function, redirect output to the output file
psql -d $DATABASE_NAME -f $FUNCTION_FILE > $OUTPUT_FILE

# Display the contents of the output file
cat $OUTPUT_FILE
