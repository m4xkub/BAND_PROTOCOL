
# BAND PROTOCOL

This repository contains programs for **BAND PROTOCOL** company.

## Problem 1: Boss Baby's Revenge

- **Time Complexity**: O(n)
- **Space Complexity**: O(n)

### Instructions:
1. Compile the program using `g++`:
   ```bash
   g++ FirstQuestion.cpp -o FirstQuestion
   ```
2. Run the compiled program:
   ```bash
   ./FirstQuestion
   ```

## Problem 2: Superman's Chicken Rescue

- **Time Complexity**: O(n)
- **Space Complexity**: O(n)

### Instructions:
1. Compile the program using `g++`:
   ```bash
   g++ SecondQuestion.cpp -o SecondQuestion
   ```
2. Run the compiled program:
   ```bash
   ./SecondQuestion
   ```

## Problem 3: Transaction Broadcasting and Monitoring Client

1. **Navigate to the `ThirdQuestion` folder**:
   ```bash
   cd ThirdQuestion
   ```

2. **Run the Go program**:
   ```bash
   go run main.go
   ```

3. **Server Configuration**:
   - The server runs on port **8080** by default. You can change this in `main.go` (line 38).

4. **Endpoints**:

   - `POST localhost:8080/CreateTransaction`: Creates a new transaction.

     **Example body**:
     ```json
     {
       "symbol": "BTC",
       "price": 4500
     }
5. **Explaination**:
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

6. **Discussion**:
   - I can create this program completly using a microservice architecture for scalability
   - In a real-world scenario, I will divide this program into two parts: one for transaction creation and the other for transaction monitoring. If the load on transaction creation increases, I can easily scale up the    replica from 1 to 2 in Kubernetes and scale it down when necessary
   - In Performance aspect : This program is still have improvements in this part.Because we want to only monitor the transaction that still in pending state but this program is monitoring all of them.
   - In Reliable aspect : As I said that this program is use memory to store data so data will lost if program is shutting down but we can fix this by using Nosql database like mongodb for read/write performance
     and reliable instead of using array. But as i said before this program is a simple program for test so i think that there is no need to use real database to store data.
   - this program is for test not for production so that i decided to use monolith architecture for easily develop, no need for complex set up and deployment



  
