package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
    "cloudfix-linter/logger"
	"github.com/sirupsen/logrus"
	tfjson "github.com/hashicorp/terraform-json"
)

var Log *logrus.Logger

func initializeLogger(){
	var flag string
	var filePath string
	fmt.Println("Enter Path to log debug files: (Y/N)")
	fmt.Scan(&flag)
	if flag !="Y" && flag !="N"{
		panic("Enter either of Y or N only")
	}
	if flag == "Y"{
		fmt.Scan(&filePath)
	}
	Log = logger.NewLogger(filePath)
}

type Parameter struct {
	IdealType      string `json:"Migrating to instance type"`
	RecentSnapshot bool   `json:"Has recent snapshot"`
}

//structure for unmarshalling the reccomendation json response from cloudfix
type ResponseReccos struct {
	Id                     string
	Region                 string
	PrimaryImpactedNodeId  string
	OtherImpactedNodeIds   []string
	ResourceId             string
	ResourceName           string
	Difficulty             int
	Risk                   int
	ApplicationEnvironment string
	AnnualSavings          float32
	AnnualCost             float32
	Status                 string
	Parameters             Parameter
	TemplateApproved       bool
	CustomerId             int
	AccountId              string
	AccountNickname        string
	OpportunityType        string
	OpportunityDescription string
	GeneratedDate          string
	LastUpdatedDate        string
}

type Orchestrator struct {
	// No Data Fields are required for this class
}

// Memeber functions for the Orchestrator class follow:

func (o *Orchestrator) parseReccos(reccos []byte) map[string]map[string]string {
	mapping := map[string]map[string]string{}
	var responses []ResponseReccos
	err := json.Unmarshal(reccos, &responses)
	if err != nil {
		Log.Error("Failed to unmarshall reccomendations")
		panic(err)
	}
	Log.Info("Reccomendations unmarrshalled succesfully!")
	for _, recco := range responses {
		awsID := recco.ResourceId
		oppurType := recco.OpportunityType
		temp := map[string]string{}
		switch oppurType {
		case "Gp2Gp3":
			temp["type"] = "gp3"
			mapping[awsID] = temp
		case "Ec2IntelToAmd":
			var idealType = recco.Parameters.IdealType
			temp["instance_type"] = idealType
			mapping[awsID] = temp
		default:
			Log.Warnf("Unknown Oppurtunity Type for resource ID: \"%s\"\n", awsID)
			temp["NoAttributeMarker"] = recco.OpportunityDescription
			mapping[awsID] = temp
		}
	}
	return mapping
}

func (o *Orchestrator) extractModulePaths(jsonString []byte) ([]string, error) {
	var modulePaths []string
	//byteValue := []byte(jsonString)
	/*
		Initialising a variable result which stores the data in the format of defined structure.
		Structure: "map(key->string,value->(array of map(key->string,value->interface))"
	*/
	var result map[string][]map[string]interface{}
	err := json.Unmarshal(jsonString, &result)
	if err != nil {
		Log.Error("Failed to unmarshall module paths from json string")
		return modulePaths, err
	}
	noOfModules := len(result["issues"])
	modulePaths = make([]string, noOfModules)
	for key, element := range result["issues"] {
		modulePaths[key] = fmt.Sprint(element["message"])
	}
	return modulePaths, nil
}

func (o *Orchestrator) getTagToID() (map[string]string, error) {
	tagToID := make(map[string]string)
	TfLintOutData, errT := exec.Command("terraform", "show", "-json").Output()
	if errT != nil {
		Log.Error("Failed to execute terraform show")
		return tagToID, errT
	}
	var tfState tfjson.State
	errU := tfState.UnmarshalJSON(TfLintOutData)
	if errU != nil {
		Log.Error("Failed to unmarshall byte array extracted from terraform show")
		return tagToID, errU
	}
	//for root module resources
	for _, rootResource := range tfState.Values.RootModule.Resources {
		o.addPairToTagMap(rootResource, tagToID)

	}
	// for all the resources present in child modules under the root module
	for _, childModule := range tfState.Values.RootModule.ChildModules {
		for _, childResource := range childModule.Resources {
			o.addPairToTagMap(childResource, tagToID)
		}
	}
	return tagToID, nil
}

func (o *Orchestrator) addPairToTagMap(resource *tfjson.StateResource, tagToID map[string]string) {
	AWSResourceIDRaw, ok := resource.AttributeValues["id"]
	if !ok {
		Log.Warn("ID not present")
		return
	}
	AWSResourceID := AWSResourceIDRaw.(string)
	tagsRaw, ok := resource.AttributeValues["tags"]
	if !ok {
		Log.Warn("tags are not present")
		return
	}
	tags := tagsRaw.(map[string]interface{})
	yorTagRaw, ok := tags["yor_trace"]
	if !ok {
		Log.Warn("yor_trace not present")
		return
	}
	yorTag := yorTagRaw.(string)
	AWSResourceIDStrip := strings.Trim(AWSResourceID, "\n")
	AWSResourceIDTrim := strings.Trim(AWSResourceIDStrip, `"`)
	yorTagStrip := strings.Trim(yorTag, "\n")
	yorTagTrim := strings.Trim(yorTagStrip, `"`)
	if yorTagTrim == "" || AWSResourceIDTrim == "" {
		return
	}
	tagToID[yorTagTrim] = AWSResourceIDTrim
}

func main() {
    initializeLogger()
	var orches Orchestrator
	var persist Persistance
	reccosFileName := "recos.txt"
	currPWD, _ := exec.Command("pwd").Output()
	currPWDStr := string(currPWD[:])
	currPWDStrip := strings.Trim(currPWDStr, "\n")
	currPWDStrip += "/reccos.json"
	fileR, errR := ioutil.ReadFile(currPWDStrip)
	if errR != nil {
		Log.Error("Reading from reccos.json failed!")
		panic(errR)
	}
	Log.Info("Reading from reccos.json Successful!")
	reccosMapping := orches.parseReccos(fileR)
	errP := persist.store_reccos(reccosMapping, reccosFileName)
	if errP != nil {
		Log.Error("Storing reccos mapping to persistance manager failed!")
		panic(errP)
	}
	Log.Info("Storing reccos mapping to persistance manager Successful!")
	os.Setenv("ReccosMapFile", reccosFileName)
	tagFileName := "tagsID.txt"
	tagToIDMap, errG := orches.getTagToID()
	if errG != nil {
		Log.Error("Storing reccos mapping to persistance manager failed!")
		panic(errG)
	}
	Log.Info("Storing tag to ID mapping to persistance manager Successful!")
	errT := persist.store_tagMap(tagToIDMap, tagFileName)
	if errT != nil {
		Log.Error("Storing tag to ID mapping to persistance manager failed!")
		panic(errT)
	}
	Log.Info("Storing tag to ID mapping to persistance manager Successful!")
	os.Setenv("TagsMapFile", tagFileName)
	modulesJson, _ := exec.Command("tflint", "--only=module_source", "-f=json").Output()
	modulePaths, errM := orches.extractModulePaths(modulesJson)
	if errM != nil {
		Log.Error("Extracting module paths from modules json failed!")
		return
	}
	Log.Info("Extracting module paths from modules json Successful!")
	output, _ := exec.Command("tflint", "--module", "--disable-rule=module_source").Output()
	fmt.Print(string(output))
	for _, module := range modulePaths {
		outputM, _ := exec.Command("tflint", module, "--module", "--disable-rule=module_source").Output()
		fmt.Print(string(outputM))
	}
	Log.Info("Recommendations display Successful!")
}
