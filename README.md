# Golang Stream Processing
This is a demo project of how to implement stream processing in Golang. Data to be read are stored in `data` directory, the server will read each line and process it accordingly.

## How to Run

1. Clone this repository

2. Run this command to run the app

```
make run
```

## How it Works

1. The server will open `data` directory that store the stock data.

2. It will open each file and read it line by line.

3. After each time the server read a line, it will pass the line to another function.

4. The function will spawn a goroutine to handle the processing.

5. The processing is done concurrently while the server read another line.

## Further Improvement
* Map might not be a good use to save data if the write operation is done more than read operation. Saving the stock detail to database will be better.
* Using microservice to task the reading to one service and the processing to another service can be done by using gRPC or Message Queue.