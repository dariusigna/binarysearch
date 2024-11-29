# Index API

## Overview

This application offers a REST API to find the index of a given value in a sorted array. If the exact value is not found, the index of the closest value within 10% is returned.

## Features

- Find the index of a given value in a sorted array
- If the exact value is not found, return the index of the closest value within 10%
- if the value is not found and the closest value is at the beginning of the array, return an error
- The array is read from a file

## API Endpoints

- `GET /index/:number` - Find the index of a given value in a sorted array. `number` - The value must be between 0 and 10000000.

## Installation

1. Clone the repository:
   ```sh
   git clone <repository-url>
   
2. Add/Update .env file in the root directory and add the following environment variables:
   ```sh
   PORT=3000
   LOG_LEVEL=info
   INPUT_FILE=./input.txt
   ```
   
3. Start the server
    ```sh
    make start
    ```
   
4. Test the endpoint
    ```sh
    curl http://localhost:3003/index/5
    ```
   

## Frontend

The frontend is a simple React application that allows the user to input a number and find the index of that number in the sorted array.
For more information, please refer to the [frontend README](./web-app/README.md).

