#!/bin/bash
cat > install.sh <<EOF1
#! /bin/bash

if [[ "\$OSTYPE" =~ ^darwin ]]; then
  OS=darwin
  brew install wget
else
  OS=linux
fi

Arch=\$(uname -m)

if [[ "\$Arch" == "x86_64" || "\$Arch" == "amd64" ]]; then
    ARCH=amd64
elif [[ "\$Arch" == "aarch64" || "\$Arch" == "arm64" ]]; then
    ARCH=arm64
elif [[ "\$Arch" == "i686" || "\$Arch" == "i386" ]]; then
    ARCH=386
elif [ "\$Arch" = "armhf" ]; then
    ARCH=arm
else 
    echo "Unsupported platform"
    exit 1
fi

PLATFORM=\$OS
PLATFORM+="_"
PLATFORM+=\$ARCH

mkdir -p cloudfix-linter
cd cloudfix-linter

#Installing terraform 
echo "Installing Terraform"
TERRAFORM_VERSION=1.2.6
file_name=terraform_\${TERRAFORM_VERSION}_\${PLATFORM}.zip
( wget https://releases.hashicorp.com/terraform/\${TERRAFORM_VERSION}/\${file_name} -O \${file_name} --no-check-certificate \\
  && unzip -u \${file_name} \\
  && rm \${file_name})
path=\$(pwd)
path+="/terraform"
alias terraform=\$path
chmod +x terraform
echo "Terraform installed successfully"

#Installing yor
echo "Installing Yor"
YOR_VERSION=0.1.150
file_name=yor_\${YOR_VERSION}_\${PLATFORM}.tar.gz
wget https://github.com/bridgecrewio/yor/releases/download/\${YOR_VERSION}/\${file_name} -O \${file_name} --no-check-certificate \\
&& tar -xvzf \${file_name} \\
  && rm \${file_name}               
path=\$(pwd)
path+="/yor"
alias yor=\$path
chmod +x yor
echo "Yor installed successfully"

#Installing tflint 
# plugin updated for compatibility with tflint v0.44.1
echo "Installing Tflint"
export TFLINT_VERSION=v0.44.1
file_name=tflint_\${PLATFORM}.zip
(wget https://github.com/terraform-linters/tflint/releases/download/\${TFLINT_VERSION}/\${file_name} -O \${file_name} --no-check-certificate  \\
  && unzip -u \${file_name} \\
  && rm \${file_name})
# Setting alias for tflint so that it can be used via command line without referencing the binary path
path=\$(pwd)
path+="/tflint"
alias tflint=\$path
chmod +x tflint
echo "Tflint installed successfully"

VERSION_TAG=$(git describe --tags --abbrev=0)
# Install cloudfix-linter
echo "Installing cloudfix-linter"
file_name=cloudfix-linter-developer_\${PLATFORM}
(wget https://github.com/trilogy-group/cloudfix-linter/releases/download/\${VERSION_TAG}/\${file_name} -O \${file_name} --no-check-certificate \\
  && mv \${file_name} cloudfix-linter)
# Setting alias for cloudfix-linter so that it can be used via command line without referencing the binary path
path=\$(pwd)
path+="/cloudfix-linter"
alias cloudfix-linter=\$path
chmod +x cloudfix-linter
echo "Cloudfix-linter installed successfully"
EOF1

cat >install.ps1 <<EOF2
# Finding OS architecture

\$is64Bit = Test-Path 'Env:ProgramFiles(x86)'
\$PLATFORM="Unidentified Operating System"
# Identifying the Operting system Architecture
if(\$is64Bit){
    \$PLATFORM="windows_amd64"
}else {
    \$PLATFORM="windows_386"
}


\$OUT_PATH= \$(Get-Item .).FullName+"\cloudfix-linter\"
if (-Not (Get-Item \$OUT_PATH)) { New-Item -Path \$OUT_PATH -ItemType Directory }

\$VERSION_TAG=$(git describe --tags --abbrev=0)
# Install cloudfix-linter
Write-Output "Installing cloudfix-linter-cloudformation........"
\$OUT_PATH_CFT=\$OUT_PATH+"cloudfix-linter-cloudformation.exe"
\$DOWNLOAD_ADDRESS="https://github.com/trilogy-group/Cloudfix-linter-Cloudformation-Release/releases/download/"+\$VERSION_TAG
Invoke-WebRequest -URI \${DOWNLOAD_ADDRESS}/cloudfix-linter-cloudformation_\${PLATFORM}.exe -OutFile \$OUT_PATH_CFT
\$TEMP=\$OUT_PATH+"cloudfix-linter-cloudformation.exe"
Set-Alias -Name cloudfix-linter-cloudformation -Value \$TEMP -Scope Global
Write-Output "Cloudfix-linter installed successfully"


Write-Output "Installing cloudfix-linter-cloudformation........"
\$OUT_PATH_CFT=\$OUT_PATH+"mynewrule.py"
Invoke-WebRequest -URI \${DOWNLOAD_ADDRESS}/mynewrule.py -OutFile \$OUT_PATH_CFT


# Installing CFN-Lint
pip install cfn-lint
EOF2