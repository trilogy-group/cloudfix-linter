# cloudfix-linter
A linting tool for HashiCorp's Terraform to flag recommendations from Cloudfix

## Usage guide
1) Install golang (https://go.dev/doc/install)

2) Download “install.sh” and add it to the directory of terraform code for which you need to see recommendations. Now run command “./install.sh”  from inside working directory of your terraform code through command line. All the dependencies would get installed.
(We have to decide the location from where user will download install.sh file).

3). Run "terraform" on the command line to verify that it is recognised as a command.

4). Ensure that terraform can access your AWS account. You can either put in the access key and the secret key inside of the provider "aws" block eg: in the main.tf file
provider "aws" { region = "us-east-1" access_key = "my-access-key" secret_key = "my-secret-key" }
or you could export AWS_ACCESS_KEY_ID , AWS_SECRET_ACCESS_KEY , AWS_SESSION_TOKEN as enviroment variables (this approach is preffered). More information on how to give access can be found [here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
5). Run "terraform init" to initialise terraform, and then "terraform apply" to deploy the resources.
6). Run command “./cloudfixlinter” from directory in which your terraform code is present. You should see reccomendations being flagged in your console.

Note :- If you make any changes to your terraform code, You first have to deploy them using “terraform apply” and then run “./cloudfixlinter” command again through working directory of your terraform code to see reccomendations being flagged according to recent changes. 

Note:- If you do not have terraform code template to test our tool. You can use [this](https://github.com/trilogy-group/cloudfixLinter-demo) demo