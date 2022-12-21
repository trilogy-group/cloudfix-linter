#! /bin/bash


Arch=$(uname -m)

if [[ "$Arch" == "x86_64" || "$Arch" == "amd64" ]]; then
    PLATFORM=linux_amd64
elif [[ "$Arch" == "aarch64" || "$Arch" == "arm64" ]]; then
    PLATFORM=linux_arm64
elif [[ "$Arch" == "i686" || "$Arch" == "i386" ]]; then
    PLATFORM=linux_386
elif [ "$Arch" = "armhf" ]; then
    PLATFORM=linux_arm
else 
    echo "Unsupported platform"
    exit 1
fi

mkdir cloudfix-linter
cd cloudfix-linter
shopt -s expand_aliases
#Installing terraform 
TERRAFORM_VERSION=1.2.6
(wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${PLATFORM}.zip \
  && unzip terraform_${TERRAFORM_VERSION}_${PLATFORM}.zip \
  && rm terraform_${TERRAFORM_VERSION}_${PLATFORM}.zip)
path=$(pwd)
path+="/terraform"
alias terraform=$path
chmod +x terraform

#Installing yor_trace 
YOR_VERSION=0.1.158
wget https://github.com/bridgecrewio/yor/releases/download/${YOR_VERSION}/yor_${YOR_VERSION}_${PLATFORM}.tar.gz \
&& tar -xvzf yor_${YOR_VERSION}_${PLATFORM}.tar.gz \
&& rm yor_${YOR_VERSION}_${PLATFORM}.tar.gz               
path=$(pwd)
path+="/yor"
alias yor=$path
chmod +x yor

#Installing tflint 
# higher version have breaking changes to the plugin system and hence we can't install them without changing the plugin
export TFLINT_VERSION=v0.39.3
(wget https://github.com/terraform-linters/tflint/releases/download/${TFLINT_VERSION}/tflint_${PLATFORM}.zip \
  && unzip tflint_${PLATFORM}.zip \
  && rm tflint_${PLATFORM}.zip)
# Setting alias for tflint so that it can be used via command line without referencing the binary path
path=$(pwd)
path+="/tflint"
alias tflint=$path
chmod +x tflint

# Install cloudfix-linter
echo "Installing cloudfix-linter"
(wget https://github.com/trilogy-group/cloudfix-linter/releases/latest/download/cloudfix-linter_${PLATFORM} \
  && mv cloudfix-linter_${PLATFORM} cloudfix-linter)
# Setting alias for cloudfix-linter so that it can be used via command line without referencing the binary path
path=$(pwd)
path+="/cloudfix-linter"
alias cloudfix-linter=$path
chmod +x cloudfix-linter