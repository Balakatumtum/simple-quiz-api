# simple-quiz-api
Backend - Golang Gin Framework  
Database - just in memory, no database

# user stories
User should be able to get questions with answers  
User should be able to answer all the questions and then post his/hers answers and get back how many correct answers he/she able to get.  

# to run API server
exec: go run main.go

# quiz
A simple questions with a set of 3 answer choices. To request the questions and its answer choices do  
exec: curl localhost:3000/exampaper

# input Answers
Input your answers in accordance to the current number of available questions  
exec: curl "localhost:3000/input?answer1=..&answer2=.." --request "POST"

an example using the default database without extra questions would be like  
exec: curl "localhost:3000/input?answer1=7&answer2=39&answer3=HYGGE!&answer4=Advancing+Humanity" --request "POST"

# addquestion
Add question to the database, create a json file do  
exec: curl localhost:3000/exampaper --include --header "Content-Type: application/json" -d @file_name.json --request "POST"

example using json file within the /database/extraquestion.json would be:  
exec: curl localhost:3000/exampaper --include --header "Content-Type: application/json" -d @extraquestion.json --request "POST"
