## Beat Ride Fare Calculator

## Setup

### Requirements
- Go

### MacOS
```
$ go mod download
$ go run fareReportGenerator.go
```

## Run

- specify input and output file path inside /constant/constant.go
- place input file to the mentioned input path
```
$ go run fareReportGenerator.go

```

### Algorithm
```
1. read csv file from the input path through csv reader
2. create two channel *tupleBatch and *rideReports with some bufferSize
3. generate tuplesBatch (list of data point of a perticular ride) concurrently using goroutine ans push to tupleBatch channel.
4. again concurrenly read dataBatch from tuplesBatch channel and compute fare.
5. push fareReport to rideReports channel
6. concurretly read fare from rideReports channel and write to output file
```

#### Ride Rule
```
. Based on the time day/night or zone it will compute and validate
. default zone = "Europe/Athens"

```