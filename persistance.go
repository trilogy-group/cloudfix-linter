package main

import (
	"fmt"
	"os"
)

type Persistance struct {
	// No data fields required for this class. Data shall be persisted using the file system
}

//Member functions for the class follow:

func (p *Persistance) store_reccos(reccosMap map[string]map[string]string, fileNameForReccos string) error {
	file, err := os.Create(fileNameForReccos)
	if err != nil {
		Log.Error("Failed to create file for storing reccos")
		return err
	}
	for key, innerMap := range reccosMap {
		for innerKey, innerValue := range innerMap {
			toWrite := fmt.Sprintf("%s:%s:%s\n", key, innerKey, innerValue)
			_, err := file.WriteString(toWrite)
			if err != nil {
				Log.Error("Failed to write string to store reccos file")
				return err
			}
		}
	}
	Log.Info("Storing reccos to persistance successful!")
	return nil
}

func (p *Persistance) store_tagMap(tagToIDMap map[string]string, fileNameForTagMap string) error {
	file, err := os.Create(fileNameForTagMap)
	if err != nil {
		Log.Error("Failed to create file for storing tag to ID mapping")
		return err
	}
	for key, value := range tagToIDMap {
		toWrite := fmt.Sprintf("%s:%s\n", key, value)
		_, err := file.WriteString(toWrite)
		if err != nil {
			Log.Error("Failed to write string to tagID mapping file")
			return err
		}
	}
	Log.Info("Storing tag to ID to persistance successful!")
	return nil
}
