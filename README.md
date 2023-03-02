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
1). [TERRAFORM](https://github.com/trilogy-group/cloudfix-linter/blob/ReadmeUpdate/Docs/Readme.Terraform.md)    
2). [CLOUDFORMATION]()   
3). [CDK](https://github.com/trilogy-group/cloudfix-linter/blob/ReadmeUpdate/Docs/Readme.CDK.md)     


## Contributing

The project uses a custom ruleset written for [TfLint](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/architecture.md) to flag reccomendations from cloudfix. The github repo for the ruleset can be accessed [here] (https://github.com/trilogy-group/tflint-ruleset-template)

### Local debugging
```bash
TODO The commands
```

### Building and publising
 
 See the Github action file for details.
