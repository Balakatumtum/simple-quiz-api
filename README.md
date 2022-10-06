# simple-quiz-api
﻿Backend - Golang Gin Framework
﻿Database - just in memory, no database

﻿# To run API server
﻿CD exec: go run main.go

﻿# Quiz
﻿A simple questions with a set of 3 answer choices. To request the questions and its answer choices run
﻿exec: curl localhost:3000/exampaper

﻿# Input Answers
﻿Input your answers in accordance to the current number of available questions
﻿exec: curl "localhost:3000/input?answe1r=..&answer2=.." --request "POST"

﻿an example using the default database without extra questions would be like
﻿exec: curl "localhost:3000/input?answe1r=7&answer2=39&answer3=VARITAS!&answer4=UI" --request "POST"

﻿# AddQuestion
﻿Add question to the database, create a json file and
﻿exec: curl localhost:3000/exampaper --include --header "Content-Type: application/json" -d @file_name.json --request "POST"

﻿example using json file within the /database/extraquestion.json would be:
﻿exec: curl localhost:3000/exampaper --include --header "Content-Type: application/json" -d @file_name.json --request "POST"
