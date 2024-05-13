#!/bin/sh

OUTPUT_FILE="schema/combined.sql"
INPUT_FILES="schema/*.m.sql"

sqlfluff format $INPUT_FILES --dialect postgres
cat > $OUTPUT_FILE <<- EOM
-- Copyright (©) 2024 - Shivam Kumar Jha - All Rights Reserved, Proprietary and confidential
-- Unauthorised copying of this file, via any medium is strictly prohibited
-- Auto generated, Don't edit by hand

EOM
tail -n +3 $INPUT_FILES >> $OUTPUT_FILE
sed -i 's/==>/-- ==>/g' $OUTPUT_FILE
