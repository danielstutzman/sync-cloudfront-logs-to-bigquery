package influxdb

import (
	"time"

	"github.com/danielstutzman/sync-log-files-to-db/src/log"
	clientPkg "github.com/influxdata/influxdb/client/v2"
)

func (conn *InfluxdbConnection) InsertMaps(tagsSet map[string]bool,
	maps []map[string]interface{}) {

	// Create a batch
	points, err := clientPkg.NewBatchPoints(clientPkg.BatchPointsConfig{
		Database: conn.databaseName,
	})
	if err != nil {
		log.Fatalw("Error from NewBatchPoints", "err", err)
	}

	for _, mapUnfiltered := range maps {
		tags := map[string]string{}
		fields := map[string]interface{}{}
		for key, value := range mapUnfiltered {
			if key == "timestamp" {
				// skip
			} else if tagsSet[key] {
				tags[key] = value.(string)
			} else {
				fields[key] = value
			}
		}

		point, err := clientPkg.NewPoint(conn.measurementName, tags,
			fields, mapUnfiltered["timestamp"].(time.Time))
		if err != nil {
			log.Fatalw("Error from NewPoint", "err", err)
		}
		points.AddPoint(point)
	}

	if err := conn.client.Write(points); err != nil {
		log.Fatalw("Error from Write", "err", err)
	}
}
