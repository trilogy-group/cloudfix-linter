# cloudfix-linter

## Who is the product for

The product is for anyone who uses terraform to manage their AWS infrastructure and would like to know how best they can optimise their cloud infrastucture in order to save costs.

## What is the product

It is a command line tool that flags optimisation oppurtunities detected by Cloudfix for the resources that have been deployed using terraform. It'll either flag the specific attribute within the resource that needs to be changed (along with what it needs to be changed to), or in the case that such an attribute does not exist, describe the oppurtunity against the name of the resource about which the oppurtunity is present. 

## Pre-requisites to use the product

1. An active cloudfix account at https://app.cloudfix.com/
2. Resources deployed on AWS using terraform for which you would like to see reccomendations.

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

## Usage guide
- [Terraform](./TERRAFORM.md)
- [Cloudformation](./CLOUDFORMATION.md)
- [CDK](./CDK.md)