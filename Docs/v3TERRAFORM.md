## Command summary
Use the command `cloudfix-linter tf`
```
This tool brings Cloudfix recommendations for your terraform code, to your terminal

Usage:
cloudfix-linter tf [flags]
cloudfix-linter tf [command]

Available Commands:
init        To initialise the directory. Run this before asking for recommendations
addTags     Add tags to your terraform code to trace them back to the cloud
reco        To flag recommendations

Flags:
-h, --help   help for tf

Use "cloudfix-linter tf [command] --help" for more information about a command.
```
## Usage guide
1. Run command 
	```bash
	wget -O - https://github.com/trilogy-group/cloudfix-linter/releases/latest/download/install.sh | bash
	```

2. Ensure that terraform can access your AWS account. You can user one of the following

    1. Devconnect with [saml2aws](https://github.com/Versent/saml2aws)
    2. Set the access key and the secret key inside of the provider "aws" block eg: in the main.tf file provider "aws" { region = "us-east-1" access_key = "my-access-key" secret_key = "my-secret-key" } 
    3. Set and export AWS_ACCESS_KEY_ID , AWS_SECRET_ACCESS_KEY , AWS_SESSION_TOKEN as enviroment variables. More information on how to give access can be found [here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

3. This version works with CloudFix v3 so make sure you have credentials to https://app.cloudfix.com/

4. Open the folder with your terraform files as the rootfolder. 
   Eg - 
   Wrong setup :The rootfolder does not contain the terrafrom files, instead terraform-dev-setup/s3 has the terraform files, extension does not supports such file heirarchy
   image.png

   Correct Setup :The rootfolder S3-Buckets has the terraform files directly at his next level
   image.png

5. From your terraform code working directory do "cloudfix-linter init".
	```bash
	cd my-terraform-project
	cloudfix-linter/cloudfix-linter tf init
	cloudfix-linter/cloudfix-linter tf --help
	```


6. Run "terraform apply" to deploy the resources from your terraform code working directory. Currently the CLI uses terraform version 1.2.6 which it downloads itself, this maybe not be compatible with your terraform version. To avoid conflicts with your local terraform version we recommend using the bundled terraform. If your current working directory is the project root then run the follow commands to use the bundled terraform. [Video demo](https://www.loom.com/share/f27c295e251b452696516055b65323f1)
	```bash
	cloudfix-linter/terraform init
	cloudfix-linter/terraform apply
	```

7. To get recommendations from cloudfix and see them through CLI run command 
    ```
    cloudfix-linter/cloudfix-linter tf reco
    ```

Note :- If you make any changes to your terraform code, You first have to deploy them using `cloudfix-linter/terraform apply` and then run `cloudfix-linter tf reco` command again through working directory of your terraform code to see reccomendations being flagged according to the recent changes. 

Note:- If you do not have terraform code template to test this tool. You can use [this](https://github.com/trilogy-group/cloudfixLinter-demo) demo


## Tflint

The project uses a custom ruleset written for [TfLint](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/architecture.md) to flag reccomendations from cloudfix. The github repo for the ruleset releases can be accessed [here] (https://github.com/trilogy-group/tflint-ruleset-template)
