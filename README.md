# cloudfix-linter
A linting tool for HashiCorp's Terraform to flag recommendations from Cloudfix

## Usage guide
1) Run command “wget https://github.com/trilogy-group/cloudfix-linter/releases/latest/download/install.sh” and Installation sript would get installed. Give execution access to the script using command "chmod +x install.sh". Now run command “./install.sh”  from directory where you have run the above command. All the dependencies would get installed.

2). Ensure that terraform can access your AWS account. You can either put in the access key and the secret key inside of the provider "aws" block eg: in the main.tf file
provider "aws" { region = "us-east-1" access_key = "my-access-key" secret_key = "my-secret-key" } or you could export AWS_ACCESS_KEY_ID , AWS_SECRET_ACCESS_KEY , AWS_SESSION_TOKEN as enviroment variables (this approach is preffered). More information on how to give access can be found [here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

3). From your terraform code working directory do "cloudfix-linter init".

4). Run "terraform apply" to deploy the resources from your terraform code working directory.

5). Run command “cloudfix-linter” from directory in which your terraform code is present. This gives user usage guide on how to use our tool, the available commands and their functions. Also detailed description of these commands could be found by command "cloudfix-linter [command] --help".

6). For getting the recommendations from cloudfix and see them through CLI run command "cloudfix-linter flagRecco" and the recommendations would get displayed throug CLI.

Note :- If you make any changes to your terraform code, You first have to deploy them using “terraform apply” and then run “cloudfix-linter” command again through working directory of your terraform code to see reccomendations being flagged according to recent changes. 

Note:- If you do not have terraform code template to test our tool. You can use [this](https://github.com/trilogy-group/cloudfixLinter-demo) demo
