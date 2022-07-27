package main

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"strings"
)

//Structure for unmarshalling the oppurtunityType to Attributes mapping (the mapping is present in "mappingAttributes.json")
type IdealAttributes struct {
	AttributeType  string `json:"Attribute Type"`
	AttributeValue string `json:"Attribute Value"`
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
	Parameters             map[string]interface{}
	TemplateApproved       bool
	CustomerId             int
	AccountId              string
	AccountNickname        string
	OpportunityType        string
	OpportunityDescription string
	GeneratedDate          string
	LastUpdatedDate        string
}

type CloudfixManager struct {
}

//Member functions follow:

func (c *CloudfixManager) createMap(reccos []byte, attrMapping []byte) map[string]map[string]string {
	mapping := map[string]map[string]string{} //this is the map that has to be returned in the end
	var responses []ResponseReccos
	if len(reccos) == 0 {
		Log.Warn("No reccommendations has been received")
		return mapping
	}
	errU := json.Unmarshal(reccos, &responses) //the reccomendations from cloudfix are being unmarshalled
	if errU != nil {
		Log.Error("Failed to unmarshall recommendations from cloudfix")
		return mapping
	}
	var attrMap map[string]IdealAttributes
	errM := json.Unmarshal(attrMapping, &attrMap) //the mapping that defines how to parse an oppurtunity type is being unmarshalled here
	if errM != nil {
		Log.Error("Failed to unmarshall attribute mapping")
		return mapping
	}
	for _, recco := range responses { //iterating through the recommendations one by one
		awsID := recco.ResourceId
		oppurType := recco.OpportunityType
		attributeTypeToValue := map[string]string{}
		attributes, ok := attrMap[oppurType]
		if ok {
			//known oppurtunity type has been encountered
			atrValueByPeriod := strings.Split(attributes.AttributeValue, ".")
			if atrValueByPeriod[0] == "parameters" {
				//the ideal value needs to be picked up from cloudfix reccomendations
				valueFromReccos, ok := recco.Parameters[atrValueByPeriod[1]]
				if !ok {
					Log.Warn("Attribute is not present")
					//if the code reaches here, then this means that the strategy for parsing has not been made correctly.
					// So we are resorting to showing the reccomendation against the resource name with the description for the oppurtunity
					attributeTypeToValue["NoAttributeMarker"] = recco.OpportunityDescription
				} else {
					idealAtrValue := valueFromReccos.(string) //extracting the ideal value as a string from cloudfix reccomendations
					attributeTypeToValue[attributes.AttributeType] = idealAtrValue
				}
			} else {
				//the ideal value is static and can be directly added
				attributeTypeToValue[attributes.AttributeType] = attributes.AttributeValue
			}
		} else {
			//unknown oppurtunity type has been encountered
			//So we are resorting to showing the reccomendation against the resource name with the description for the oppurtunity
			attributeTypeToValue["NoAttributeMarker"] = recco.OpportunityDescription
		}
		mapping[awsID] = attributeTypeToValue
	}
	Log.Info("Tag to ID map created successfully!")
	return mapping
}

func (c *CloudfixManager) parseReccos() map[string]map[string]string {
	//function to process the reccomendations from cloudfix and turn that into a map
	//the structure of the map is resourceID -> Attribute type that needs to be targetted -> Ideal Attribute Value
	// If there is no attribute that has to be targetted, attribute type would be filled with "NoAttributeMarker" and
	//Attribute Value would be filled with any message that in the end has to be displayed to the user
	currPWD, _ := exec.Command("pwd").Output()
	currPWDStr := string(currPWD[:])
	currPWDStrip := strings.Trim(currPWDStr, "\n")
	currPWDStrip += "/reccos.json"
	reccos, errR := ioutil.ReadFile(currPWDStrip)
	if errR != nil {
		Log.Error("Failed to read reccos json file")
		panic(errR)
	}
	attrMapping := []byte(`{
		"Gp2Gp3": {
			"Attribute Type": "type",
			"Attribute Value": "gp3"
		},
		"Ec2IntelToAmd": {
			"Attribute Type": "instance_type",
			"Attribute Value": "parameters.Migrating to instance type"
		},
		"StandardToSIT": {
			"Attribute Type": "NoAttributeMarker",
			"Attribute Value": "Enable Intelligent Tiering for this S3 Block by writing a aws_s3_bucket_intelligent_tiering_configuration resource block"
		},
		"EfsInfrequentAccess": {
			"Attribute Type": "NoAttributeMarker",
			"Attribute Value": "Enable Intelligent Tiering for EFS File by declaring a sub-block called lifecycle_policy within this resource block"
		}
	}`)
	mapping := c.createMap(reccos, attrMapping)
	Log.Info("Recommendations Parsed successfully!")
	return mapping
}
