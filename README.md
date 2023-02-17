# cloudfix-linter

## Who is the product for

The product is for anyone who uses Iac(Terraform, Cloudformation, CDK) to manage their AWS infrastructure and would like to know how best they can optimise their cloud infrastucture in order to save costs.

## What is the product ?

It is a command line tool that flags optimisation opportunities detected by Cloudfix for the resources that have been deployed using Iac. 
1. For Terraform and Cloudformatino - it'll either flag the specific attribute within the resource that needs to be changed (along with what it needs to be changed to), or in the case that such an attribute does not exist, describe the oppurtunity against the name of the resource about which the oppurtunity is present. 

2. For CDK - It will prepare a report which can be viewed in the VSCode. Report will be created for the account CDK code is deployed in. It will have all the resources created by the CDK code along with the recommendations for those resources.


## Pre-requisites to use the product.  
 1. An active cloudfix account at https://app.cloudfix.com/ 
 2. An AWS account logged in to deploy resources via IaC

## Usage guides - 
1). [TERRAFORM](https://github.com/trilogy-group/cloudfix-linter/blob/ReadmeUpdate/Readme.Terraform.md)    
2). [CLOUDFORMATION]()   
3). [CDK](https://github.com/trilogy-group/cloudfix-linter/blob/ReadmeUpdate/Readme.CDK.md)     

## Guide on how to add support for new Cloudfix Oppurtunity Types:

***New cloudfix Opportunity can be added to CLI's in case of terraform and cloudformation only.***     

A pure json mapping has been made so that support for new insights can be added easily.
Sample mapping json used incase of Terraform:

```
{
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
	}
```

For each new oppurtunity type, create a new block in the json by its name. If the opportunity type targets an attribute in specific, put in the name of the attribute for the Attribute Type. If it does not target any attribute, put in "NoAttributeMarker" instead. For the Attribute Value, if that needs to picked up from the parameters field of the cloudfix recommendation, set that as parameters.{Name of field within parameters block} (for reference take a look at the block for Ec2IntelToAmd). In case the value for the attribute is static and need not be picked up from the parameters field, it can be hardcoded directly in the json (for reference take a look at the block for Gp2Gp3). If the oppurtunity type does not target any attribute in specific, for the attribute value, put in the message that you want displayed to the user (for reference see the block for EfsInfrequentAccess)

This mapping is currently part of the code itself, So change must be done in code. These mapping could be hosted later on for public access.  


 ***Above mapping is for terraform only mapping is different in case of cloudformation***

## Contributing

The project uses a custom ruleset written for [TfLint](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/architecture.md) to flag reccomendations from cloudfix. The github repo for the ruleset can be accessed [here] (https://github.com/trilogy-group/tflint-ruleset-template)

### Local debugging
```bash
TODO The commands
```

### Building and publising
 
 See the Github action file for details.
