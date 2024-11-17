# BAND_PROTOCOL
This repository's purpose is for BAND PROTOCOL company

for program 1,2 use g++ to comply the program
example command : g++ FirstQuestion.cpp -o FirstQuestion
then : ./FirstQuestion

Problem 1: Boss Baby's Revenge
  Time complexity : O(n)
  Space complexity : O(n)

Problem 2: Superman's Chicken Rescue
  Time complexity : O(n)
  Space complexity : O(n)

Problem 3: Transaction Broadcasting and Monitoring Client
  cd into ThridQuestion folder
  run go run main.go to start up

  server should run on port 8080 (you can change it in main.go line 38)

  localhost:8080//CreateTransaction for creating a transaction
    example of body
    {
      "symbol" : "BTC",
      "price" : 4500
    }

  - use main.go as a entry point for this program
  - controller folder is for receiving request from user
  - service folder is containing logic used in controller
  - config hold constant and config variable for this program
  - in deployment version it is better to use .env file instead of containing everything in config folder but this approach is focusing on simple program
  - using cronjob for intervaly monitoring transaction status
  - any status that is not "PENDING" is not in its final state yet
  - if all status is in final state, program will log "All transactions' status are in finalization state"
  - using array as a database for simple program
  - we should use a database because if the server shuts down, transaction data will be lost permanently
  - model hold request and response pattern for each endpoint
  - implement program based on the SOLID principles and other best practices for easy maintenance and debugging in the future



  
