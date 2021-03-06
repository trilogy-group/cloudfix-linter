package main

import (
	"os/exec"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

type TerraformManager struct {
	//No data types required
}

func (t *TerraformManager) getTagToID(TfLintOutData []byte) (map[string]string, error) {
	tagToID := make(map[string]string)
	var tfState tfjson.State
	errU := tfState.UnmarshalJSON(TfLintOutData)
	if errU != nil {
		return tagToID, errU
	}
	if tfState.Values == nil {
		//log that no resources have been deployed
		return tagToID, nil
	}
	//for root module resources
	for _, rootResource := range tfState.Values.RootModule.Resources {
		if rootResource != nil {
			t.addPairToTagMap(rootResource, tagToID)
		}
	}
	// for all the resources present in child modules under the root module
	for _, childModule := range tfState.Values.RootModule.ChildModules {
		for _, childResource := range childModule.Resources {
			if childResource != nil {
				t.addPairToTagMap(childResource, tagToID)
			}
		}
	}
	return tagToID, nil
}

func (t *TerraformManager) addPairToTagMap(resource *tfjson.StateResource, tagToID map[string]string) {
	AWSResourceIDRaw, ok := resource.AttributeValues["id"]
	if !ok {
		//log that id is not present
		return
	}
	AWSResourceID := AWSResourceIDRaw.(string)
	tagsRaw, ok := resource.AttributeValues["tags"]
	if !ok {
		//log that tags are not present
		return
	}
	tags := tagsRaw.(map[string]interface{})
	yorTagRaw, ok := tags["yor_trace"]
	if !ok {
		//log that yor_trace is not present
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

func (t *TerraformManager) getTagToIDMapping() (map[string]string, error) {
	tagToID := make(map[string]string)
	TfLintOutData, errT := exec.Command("terraform", "show", "-json").Output()
	if errT != nil {
		//Add Log
		return tagToID, errT
	}
	tagToID, err := t.getTagToID(TfLintOutData)
	if err != nil {
		//Add Log
		return tagToID, err
	}
	return tagToID, nil
}
