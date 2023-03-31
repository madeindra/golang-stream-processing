# Golang Stream Processing
This is a demo project of how to implement stream processing in Golang.

## Further Improvement
* Map might not be a good use to save data if the write operation is done more than read operation. Saving the stock detail to database will be better.
* Using microservice to task the reading to one service and the processing to another service can be done by using gRPC or Message Queue.