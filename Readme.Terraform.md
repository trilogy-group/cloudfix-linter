# CLoudfix linter - Terraform

## Usage Guide - 
1) Run command 
```bash
wget -O - https://github.com/trilogy-group/cloudfix-linter/releases/latest/download/install.sh | bash
 ```

2). Ensure that terraform can access your AWS account. You can user one of the following
    a) Devconnect with [saml2aws](https://github.com/Versent/saml2aws)
    b) Set the access key and the secret key inside of the provider "aws" block eg: in the main.tf file provider "aws" { region = "us-east-1" access_key = "my-access-key" secret_key = "my-secret-key" } 
    c) Set and export AWS_ACCESS_KEY_ID , AWS_SECRET_ACCESS_KEY , AWS_SESSION_TOKEN as enviroment variables. More information on how to give access can be found [here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

3) This version works with CloudFix v3 so make sure you have credentials to https://app.cloudfix.com/

4). From your terraform code working directory do "cloudfix-linter init".
```bash
cd my-terraform-project
cloudfix-linter init
cloudfix-linter --help
```

5). Run "terraform apply" to deploy the resources from your terraform code working directory.
```bash
terraform apply
```

6). To get recommendations from cloudfix and see them through CLI run command "cloudfix-linter flagRecco" 

Note :- If you make any changes to your terraform code, You first have to deploy them using “terraform apply” and then run “cloudfix-linter” command again through working directory of your terraform code to see reccomendations being flagged according to recent changes. 

Note:- If you do not have terraform code template to test this tool. You can use [this](https://github.com/trilogy-group/cloudfixLinter-demo) demo
