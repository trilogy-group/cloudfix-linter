# Cloudfix-linter Cloudformation

## Prerequisite

1. [pip](https://pip.pypa.io/en/stable/installation/) should be installed on the user's system

## Command summary
Use the command `cloudfix-linter cfn`
```
This tool bring Cloudfix recommendations for your AWS Cloudformation template, to your terminal

Usage:
  cloudfix-linter cfn [flags]
  cloudfix-linter cfn [command]

Available Commands:
  update      Update recommendation for cloudformation template from cloudfix. Give deployed stack names in space separated format
  reco        List template based recommendations. Run update command before running this.

Flags:
  -h, --help   help for cfn

Use "cloudfix-linter cfn [command] --help" for more information about a command.
```

## Usage guide
1. Login to AWS using terminal (in default profile) by any of the following options:
    - `aws configure` to login with permanent credentials (using `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
    - To login with temporary credentials (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`, `AWS_SECURITY_TOKEN`), follow [this](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html#using-temp-creds-sdk-cli)
    - Using `saml2aws`. For user guide visit [here](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html#using-temp-creds-sdk-cli)
    - Note: Setting `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN` as enviroment variables from terminal won't work because they are just available in the terminal instance in which have set them and not available globally.
2. The region for default profile should also be set to the region where the stack exists. Use the command `aws configure`. Following is an example of setting the region to `us-east-1`
      ```
      AWS Access Key ID [****************H44M]: 
      AWS Secret Access Key [****************9jFj]: 
      Default region name [None]: us-east-1
      Default output format [None]:
      ```
3. Deploy a stack in your AWS account using the template (No need to do this again if deployed once and that stack isn't deleted). There are two ways to do this:
    - Use AWS [console](https://us-east-1.console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks)
    - Login to aws through CLI, then run this command
      ```
      aws cloudformation deploy --template-file <locationOfTemplateFile> --stack-name <stackName>
      ```
4. Run the following command the first time you're fetching recommendations, and run it again if changes are made to the cloudformation stack
    ```
    cloudfix-linter/cloudfix-linter cfn update
    ```
5. Run the following command to view recommendations
    ```
    cloudfix-linter/cloudfix-linter cfn reco
    ```
    For json output
    ```
    cloudfix-linter/cloudfix-linter cfn reco -j
    ```

Note: If you use the given template to deploy multiple stacks(in the same env -> Account+Region), recommendations from all the stacks will be linted along with the stack name for the recommendation