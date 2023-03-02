# cloudfix-linter

## Who is the product for
# cloudfix-linter

## Who is the product for

The product is for anyone who uses Iac(Terraform, Cloudformation, CDK) to manage their AWS infrastructure and would like to know how best they can optimise their cloud infrastucture in order to save costs.

## What is the product ?

It is a command line tool that flags optimisation opportunities detected by Cloudfix for the resources that have been deployed using Iac. 
1. For Terraform and Cloudformation - it'll either flag the specific attribute within the resource that needs to be changed (along with what it needs to be changed to), or in the case that such an attribute does not exist, describe the oppurtunity against the name of the resource about which the oppurtunity is present. 

2. For CDK - It will prepare a report in JSON format. Report will be created for the account CDK code is deployed in. It will have all the resources created by the CDK code along with the recommendations for those resources.


## Pre-requisites to use the product.  
 1. An active cloudfix account at https://app.cloudfix.com/ 
 2. An AWS account logged in to deploy resources via IaC

## Usage guides - 
- [TERRAFORM](./Docs/v3TERRAFORM.md)    
- [CLOUDFORMATION](./Docs/CLOUDFORMATION.md)   
- [CDK](./Docs/CDK.md)     

## Commands summary
- `cloudfix-linter`

    ```
    A CLI tool to fetch cloudfix recommendations

    Usage:
    cloudfix-linter [command]

    Available Commands:
    tf          Cloudfix Linter for terraform
    cfn         Cloudfix Linter for AWS Cloudformation
    cdk         Cloudfix Linter for CDK
    help        Help about any command
    completion  Generate the autocompletion script for the specified shell

    Flags:
    -h, --help     help for cloudfix-linter
    -t, --toggle   Help message for toggle

    Use "cloudfix-linter [command] --help" for more information about a command.
    ```

## Contributing

The project uses a custom ruleset written for [TfLint](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/architecture.md) to flag reccomendations from cloudfix. The github repo for the ruleset releases can be accessed [here] (https://github.com/trilogy-group/tflint-ruleset-template)


## Building and publising
 
 See the Github action file for details.
