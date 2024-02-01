package handler

import (
	"backend/model"
	"backend/repository"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"log"
	"os"
	"time"
)

func GetStats(ctx *fiber.Ctx) error {
	org := os.Getenv("INFLUXDB_ORG")
	stats := generateStats(repository.InfluxDB.QueryAPI(org))
	log.Println(time.Now())
	return ctx.Status(fiber.StatusOK).JSON(stats)
}

func generateStats(queryAPI api.QueryAPI) []*model.Stats {

	var resultList []*model.Stats
	hostQuery := `from(bucket: "default_bucket")
				|> range(start: -5s)
				|> last()
				|> keep(columns: ["hostname"])
				|> distinct(column: "hostname")`
	hostResult, err := queryAPI.Query(context.Background(), hostQuery)
	if err != nil {
		panic(err)
	}

	var hostnames []string
	for hostResult.Next() {
		hostnames = append(hostnames, hostResult.Record().ValueByKey("hostname").(string))
	}

	for _, hostname := range hostnames {
		currentHost := &model.Stats{Hostname: hostname}
		resultList = append(resultList, currentHost)
		query := fmt.Sprintf(
			`from(bucket: "default_bucket")
						|> range(start: -5s)
						|> filter(fn: (r) => r["hostname"] == "%s")
						|> last()`, hostname)
		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			panic(err)
		}
		for result.Next() {
			switch result.Record().Measurement() {
			case "cpu_usage":
				currentHost.CpuUsage = result.Record().ValueByKey("_value").(float64)
			case "disk_usage":
				currentHost.DiskUsage = result.Record().ValueByKey("_value").(float64)
			case "memory_usage":
				currentHost.MemoryUsage = result.Record().ValueByKey("_value").(float64)
			}
		}

		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}

	}

	return resultList

}
