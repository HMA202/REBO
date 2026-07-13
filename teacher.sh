#!/bin/sh
export INTERVIEW_NUMBER=$(grep -o 'SEE INTERVIEW #[0-9]*' mystery/streets/Buckingham_Place | grep -o '[0-9]*')
echo "$INTERVIEW_NUMBER"
cat "mystery/interviews/interview-$INTERVIEW_NUMBER"
echo "$MAIN_SUSPECT"
