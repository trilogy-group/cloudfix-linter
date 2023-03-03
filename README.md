# cloudfix-linter

## Who is the product for

The product is for anyone who uses terraform to manage their AWS infrastructure and would like to know how best they can optimise their cloud infrastucture in order to save costs.

## What is the product

It is a command line tool that flags optimisation oppurtunities detected by Cloudfix for the resources that have been deployed using terraform. It'll either flag the specific attribute within the resource that needs to be changed (along with what it needs to be changed to), or in the case that such an attribute does not exist, describe the oppurtunity against the name of the resource about which the oppurtunity is present. 

## Pre-requisites to use the product

1. An active cloudfix account at https://app.cloudfix.com/
2. Resources deployed on AWS using terraform for which you would like to see reccomendations.

## Usage guide

Note: For CLI version v3.0.0 and above refer to this [readme](./v3README.md)
1. Run command 
	```bash
	wget -O - https://github.com/trilogy-group/cloudfix-linter/releases/download/v2.0.3/install.sh | bash
	```

2. Ensure that terraform can access your AWS account. You can user one of the following

    1. Devconnect with [saml2aws](https://github.com/Versent/saml2aws)
    2. Set the access key and the secret key inside of the provider "aws" block eg: in the main.tf file provider "aws" { region = "us-east-1" access_key = "my-access-key" secret_key = "my-secret-key" } 
    3. Set and export AWS_ACCESS_KEY_ID , AWS_SECRET_ACCESS_KEY , AWS_SESSION_TOKEN as enviroment variables. More information on how to give access can be found [here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

3. This version works with CloudFix v3 so make sure you have credentials to https://app.cloudfix.com/                       
4. Open the folder with your terraform files as the rootfolder.    
            
   Wrong setup :The rootfolder does not contain the terrafrom files, instead terraform-dev-setup/s3 has the terraform files, extension does not supports such file heirarchy.    
   <img width="700" alt="image" src="https://user-images.githubusercontent.com/110278052/222429097-b9788278-1ac1-41ae-96c7-f803e57a9643.png">

   Correct Setup :The rootfolder S3-Buckets has the terraform files directly under it.     
   <img width="690" alt="image" src="https://user-images.githubusercontent.com/110278052/222428861-0485684d-9f20-4270-b51f-7596cccccb04.png">    
   
5. From your terraform code working directory do "cloudfix-linter init".
	```bash
	cd my-terraform-project
	cloudfix-linter init
	cloudfix-linter --help
	```

6. Run "terraform apply" to deploy the resources from your terraform code working directory. Currently he CLI uses terraform version 1.2.6 which it downloads itself, this maybe not be compatible with your terraform version. To avoid conflicts with your local terraform version we recommend using the bundled terraform. If your current working directory is the project root then run the follow commands to use the bundled terraform. [Video demo](https://www.loom.com/share/f27c295e251b452696516055b65323f1)
	```bash
	cloudfix-linter/terraform init
	cloudfix-linter/terraform apply
	```

7. To get recommendations from cloudfix and see them through CLI run command "cloudfix-linter flagRecco" 

Note :- If you make any changes to your terraform code, You first have to deploy them using `terraform apply` and then run “cloudfix-linter” command again through working directory of your terraform code to see reccomendations being flagged according to recent changes. 

Note:- If you do not have terraform code template to test this tool. You can use [this](https://github.com/trilogy-group/cloudfixLinter-demo) demo


## Tflint

The project uses a custom ruleset written for [TfLint](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/architecture.md) to flag reccomendations from cloudfix. The github repo for the ruleset releases can be accessed [here] (https://github.com/trilogy-group/tflint-ruleset-template)

## Building and publising
 
 See the Github action file for details.
