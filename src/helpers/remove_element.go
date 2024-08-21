package helpers

import "github.com/lib/pq"

func RemoveElement(indexDeleted int, originalarr pq.Int64Array) pq.Int64Array {
	j := 0
	for index := range originalarr {
		if indexDeleted != index {
			originalarr[j] = originalarr[index]
			j++
		}
	}

	return originalarr[:j]
}
