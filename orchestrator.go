package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/trilogy-group/cloudfix-linter/logger"
)

func main() {

	//reccosJson:=output(get_reccos(cloudfix credentials)->json)                                    //API call to cloudfix for getting recommendations from cloudfix, has to be decided

	reccosMapping := parseReccos(reccosJson) // Parser generates a map after parsing through json file in map[string]map[string]string format

	store_reccos(reccosMapping) // storing reccos mapping inside the persistor

	IDtoTagsMapping := getIDtoTags() // function returns id to tag mapping in map[string]string format

	store_tagMap(IDtoTagsMapping) //storing id to tag mapping in the persistor

	modulesJsonByte, err := exec.Command("bash", "-c", "tflint only=get_modules -f json").Output() //moduleJsonByte stores the json returned by getmodules rule in tflint in byte format
	moduleJsonString := string(modulesJsonByte[:])

	modulePaths := extractModulePaths(moduleJsonString) // Function call to extract module paths and store it in an array

	// A file or a data type(array of string/string/ map) has to be initialised here, where all the tflint module calls outputs would be consolidated and then passed to the formatter.

	// Call to tflint root module first,  command:- "tflint . and output stored"                      // Since main.tf file path is not included in module paths

	for _, currentModulePath := range modulePaths {
		/*
		   calling tflint for every module path separately
		   Storing all the outputs in a central place for all the module which will be formatted later and returned to the user
		*/
	}

	//finalOutputFormatter()                                                                            // Formates the consolidated output from all the tflint calls to present it to user, input of formatter has to be decided(file or any data structure)

	return
}

func extractModulePaths(jsonString string) ([]string, error) {
	appLogger := logger.New()
	var modulePaths []string
	byteValue := []byte(jsonString)
	/*
		Initialising a variable result which stores the data in the format of defined structure.
		Structure: "map(key->string,value->(array of map(key->string,value->interface))"
	*/
	var result map[string][]map[string]interface{}
	err := json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		appLogger.Error().Println("Failed to unmarshall module paths from json string")
		return modulePaths, err
	}
	appLogger.Info().Println("Unmarshalled module paths succesfully!")
	noOfModules := len(result["issues"])
	modulePaths = make([]string, noOfModules)
	for key, element := range result["issues"] {
		modulePaths[key] = fmt.Sprint(element["message"])
	}
	appLogger.Info().Println("Extracted module paths succesfully!")
	return modulePaths, nil
}
