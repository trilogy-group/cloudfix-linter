# Cloudfix linter for terraform

## Who is the product for

The product is for anyone who uses terraform to manage their AWS infrastructure and would like to know how best they can optimise their cloud infrastucture in order to save costs.

## What is the product

It is a command line tool that flags off optimisation oppurtunities that exist for the resources that have been deployed using terraform. It'll either flag the specific attribute within the resource that needs to be changed (along with what it needs to be changed to), or in the case that such an attribute does not exist, describe the oppurtunity against the name of the resource about which the oppurtunity is present. 

Our linter is able to do this using reccomendations from [Cloudfix](https://cloudfix.com/). To  use your product, you will need to have an active Cloudfix account. 

## Pre-requisites to use the product

1. An active cloudfix account.
2. Resources deployed on AWS using terraform for which you would like to see reccomendations.

## How to install and run the product

Usage guide is a work in progress. The current version can be found [here] (https://github.com/trilogy-group/cloudfix-linter/pull/12)

## Contributing to cloudfix-linter

Thank you for showing interest in contriburing to cloudfix-linter. Everyone from inside the trilogy-group is welcome to contribute to the project.

The project uses a custom ruleset written for [TfLint](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/architecture.md) to flag off reccomendations from cloudfix. The github repo for the ruleset can be accessed [here] (https://github.com/trilogy-group/tflint-ruleset-template)

If you find any bugs with the product, please open and issue. Within the issue, describe the bug with exact steps on how to reproduce it. The team would be happy to look into it. If you wish to contribute to the product, the team is willing to review PRs.

Building the product:

1. Terraform installation. Commands:

    sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl 
    curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
    sudo apt-add-repository "deb [arch=$(dpkg --print-architecture)] https://apt.releases.hashicorp.com $(lsb_release -cs) main" 
    sudo apt update
    sudo apt install terraform

2. AWS Environment Configuration. Commands:

    export AWS_ACCESS_KEY_ID=NOTCONFIGURED
    export AWS_SECRET_ACCESS_KEY=NOTCONFIGURED
    export AWS_SESSION_TOKEN=NOTCONFIGURED
    export AWS_DEFAULT_REGION=NOTCONFIGURED

The above two steps can be skipped if you already have Terraform up and running (which is likely the case since you need to have resourced deployed already to use the product)

3. TfLint Installation. Command:

    curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash

4. In the root directory of your terraform template code, make a .tflint.hcl file and add the following lines.

    plugin "template"{
    enabled = true
    version = "0.1.4"
    source  = "github.com/trilogy-group/tflint-ruleset-template"
    }

5. Initialise tflint with the ruleset.

    tflint --init

6. Get the latest release of the cloudfix-linter binary for your system's architecture

7. Install yor_trace. Commands:

    brew tap bridgecrewio/tap
    brew install bridgecrewio/tap/yor

8. Add tags. Command:

    yor tag -d . yor_trace

9. Run terraform apply to apply the tags. Command:

    terraform apply

10. Run the cloudfix-linter binary from the directory of the root module of your terraform code. 

   



