#! /bin/sh

#Installing terraform 
sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl 
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=$(dpkg --print-architecture)] https://apt.releases.hashicorp.com $(lsb_release -cs) main" 
sudo apt update
sudo apt install terraform

#Installing zip/unzip to unzip tflint and other commands ahead
sudo apt install zip
sudo apt install unzip

#Installing yor_trace 
brew tap bridgecrewio/tap
brew install bridgecrewio/tap/yor

#Installing tflint 
curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash

#Adding ".tflint.hcl" file and adding commands to it
echo "plugin \"template\"{" >> .tflint.hcl 
echo "    enabled = true"  >> .tflint.hcl 
echo "    version = \"0.1.4\"" >> .tflint.hcl 
echo "    source = \"github.com/trilogy-group/tflint-ruleset-template\" " >> .tflint.hcl 
echo "}" >> .tflint.hcl 

#initializing tflint
tflint --init 


#Checking if terraform is installed
terraform 

