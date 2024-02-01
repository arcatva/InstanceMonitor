package handler

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var InfluxDB influxdb2.Client

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(".env loaded")

	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	InfluxDB = influxdb2.NewClient(url, token)

}
func TestGenerateStats(t *testing.T) {
	for _, stats := range generateStats(InfluxDB.QueryAPI(os.Getenv("INFLUXDB_ORG"))) {
		t.Log(stats.Hostname)
		t.Log(stats.CpuUsage)
		t.Log(stats.MemoryUsage)
		t.Log(stats.DiskUsage)
	}
	t.Log("PASS OK")
}
