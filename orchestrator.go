package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

type Orchestrator struct {
	// No Data Fields are required for this class
}

// Memeber functions for the Orchestrator class follow:

func (o *Orchestrator) extractModulePaths(jsonString []byte) ([]string, error) {
	var modulePaths []string
	//byteValue := []byte(jsonString)
	/*
		Initialising a variable result which stores the data in the format of defined structure.
		Structure: "map(key->string,value->(array of map(key->string,value->interface))"
	*/
	var result map[string][]map[string]interface{}
	if len(jsonString) == 0 {
		Log.Warn("Empty Json string has been sent. No modules present")
		return modulePaths, nil
	}
	err := json.Unmarshal(jsonString, &result)
	if err != nil {
		Log.Error("Failed to unmarshall module paths from json string")
		return modulePaths, err
	}
	Log.Info("Unmarshalled module paths succesfully!")
	noOfModules := len(result["issues"])
	modulePaths = make([]string, noOfModules)
	for key, element := range result["issues"] {
		modulePaths[key] = fmt.Sprint(element["message"])
	}
	Log.Info("Extracted module paths succesfully!")
	return modulePaths, nil
}

func (o *Orchestrator) runReccos() {

	var persist Persistance
	var cloudfixMan CloudfixManager
	var terraMan TerraformManager
	reccosFileName := "recos.txt"
	reccosMapping := cloudfixMan.parseReccos()
	if len(reccosMapping) == 0 {
		Log.Warn("No recommendations could be received")
		//exit gracefully
	}
	errP := persist.store_reccos(reccosMapping, reccosFileName)
	if errP != nil {
		Log.WithFields(logrus.Fields{
			"Error": errP,
		}).Error("Storing Reccos to persistance failed!")
		panic(errP)
	}
	os.Setenv("ReccosMapFile", reccosFileName)
	tagFileName := "tagsID.txt"
	tagToIDMap, errG := terraMan.getTagToIDMapping()
	if errG != nil {
		Log.WithFields(logrus.Fields{
			"Error": errG,
		}).Error("Failed to create tag to ID mapping")
		panic(errG)
	}
	errT := persist.store_tagMap(tagToIDMap, tagFileName)
	if errT != nil {
		Log.WithFields(logrus.Fields{
			"Error": errT,
		}).Error("Storing Tag Id map to persistance failed!")
		panic(errT)
	}
	os.Setenv("TagsMapFile", tagFileName)
	modulesJson, _ := exec.Command("tflint", "--only=module_source", "-f=json").Output()
	modulePaths, errM := o.extractModulePaths(modulesJson)

	if errM != nil {
		Log.WithFields(logrus.Fields{
			"Error": errM,
		}).Error("Failed to extract module paths from Json output")
		return
	}
	output, _ := exec.Command("tflint", "--module", "--disable-rule=module_source").Output()
	fmt.Print(string(output))
	for _, module := range modulePaths {
		outputM, _ := exec.Command("tflint", module, "--module", "--disable-rule=module_source").Output()
		fmt.Print(string(outputM))
	}
	Log.Info("Orchestrator run successful!")
}
