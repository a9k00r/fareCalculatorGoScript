package main

import (
	"Beat/computeAlgorithm"
	constant "Beat/constants"
	"Beat/datamodel"
	util "Beat/utils"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	tupleBatch := make(chan datamodel.TupleBatch, constant.BufferSize)
	rideReports := make(chan datamodel.RideReport, constant.BufferSize)

	reader := util.GetCsvReader()

	// use a wait group to manage synchronization
	var wg sync.WaitGroup

	// declare the workers
	for i := 0; i < constant.NoOfConcurrentRequest; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			process(tupleBatch, rideReports)
		}()
	}

	go readCsvFile(reader, tupleBatch)

	// wait for process group to finish and close rideReports
	go func() {
		wg.Wait()
		close(rideReports)
	}()

	generateRideReports(rideReports)
}

// async read row from csv and push to tupleBatch channel
func readCsvFile(reader *csv.Reader, tupleBatch chan datamodel.TupleBatch) {
	id := int64(-1)
	var dataBatch datamodel.TupleBatch
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else if tup := getTuple(record); id == -1 || id == tup.RiderId {
			id = tup.RiderId
			dataBatch.RiderId = tup.RiderId
			dataBatch.Positions = append(dataBatch.Positions, tup.Position)
		} else if id != tup.RiderId {
			tupleBatch <- dataBatch
			id = tup.RiderId
			dataBatch = datamodel.TupleBatch{}
			dataBatch.RiderId = tup.RiderId
			dataBatch.Positions = append(dataBatch.Positions, tup.Position)
		}

	}
	close(tupleBatch) // close tupleBatch to signal workers that no more job are incoming.
}

// read tuples batch from dataBatch channel and push to rideReports channel
func process(dataBatch chan datamodel.TupleBatch, rideReports chan datamodel.RideReport) {
	for {
		select {
		case tupleBatch, ok := <-dataBatch:
			if !ok {
				log.Printf("invalid tupleBatch %v", tupleBatch)
				return
			}
			rideReports <- computeAlgorithm.CalculateRideFare(tupleBatch)
		}
	}
}

// write RideReport to csv
func generateRideReports(out chan datamodel.RideReport) {
	file, err := os.Create(constant.OutputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for rideData := range out {
		err := writer.Write([]string{strconv.FormatInt(rideData.Id, 10), strconv.FormatFloat(rideData.TotalFare, 'f', 6, 64)})
		checkError("Cannot write to file", err)
	}
}

// map slice of string from csv to Tuple dataModel
func getTuple(record []string) datamodel.Tuple {
	tuple := datamodel.Tuple{}

	if id, err := strconv.ParseInt(record[0], 10, 64); err == nil {
		tuple.RiderId = id
	} else {
		log.Fatal("riderId is incorrect", record)
	}

	if lat, err := strconv.ParseFloat(record[1], 64); err == nil && util.IsValidLat(lat) {
		tuple.Position.LatLong.Lat = lat
	} else {
		log.Fatal("latitude is incorrect", record)
	}

	if lon, err := strconv.ParseFloat(record[2], 64); err == nil && util.IsValidLon(lon) {
		tuple.Position.LatLong.Lon = lon
	} else {
		log.Fatal("longitude is incorrect", record)
	}

	if timestamp, err := strconv.ParseInt(record[3], 10, 64); err == nil {
		tuple.Position.TimeStamp = timestamp
	} else {
		log.Fatal("timestamp is incorrect", record)
	}
	return tuple
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
