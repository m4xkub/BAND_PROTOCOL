
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



  
