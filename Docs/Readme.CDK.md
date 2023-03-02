# Cloudfix-linter CDK


## Additional Pre-requisites for CDK
1). Resources deployed on AWS using CDK for which you would like to see reccomendations.    
2). Ensure that our tool can access your AWS account. You can use one of the following options:
    a) Devconnect with [saml2aws](https://github.com/Versent/saml2aws) and mention the profile name in config file.   
    b) If the enviroment is configured in your CDK code, log into the account mentioned in environment for CDK code via aws-cli. 
    
    **Note**: Seting and exporting `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN` as enviroment variables is 'not supported for usage of cli with extension'.

<!-- 3). You must deploy cdk from your CDK project directory, with cdk.json at root level of cdk project.    -->

<!-- ## Using the VS Code [Cloudfix-linter extension](https://open-vsx.trilogy.devspaces.com/extension/devfactory/cloudfix-linter)

1. Deploy your cdk stacks with `--ouptut .cdkout` added at last of cdk deploy command. This creates a .cdkout directory with cdk deploy outputs(manifest.json,tree.json and stack templates etc.)

2. Open command pallete by pressing `ctrl+shift+P` and `Cloudfix-linter: cdk reco`.   

3. Input box for profile comes .Type in the profile that you want to have the result for, if no profile provided default aws-profile will be choosen.   

4. A prompt for aws profile verified and getting resources is shown, meaning all the stack resources are being fetched.   

5. A new view will open in your VSCode with the recommendation report, showing recommendations from Cloudfix corresponding to the resoruces created by the stack. -->

## Using CLI independently 
1. Installation
    ```bash
    wget -O - https://github.com/trilogy-group/cloudfix-linter-cdk/releases/latest/download/install.sh | bash
    ```
       

2. Deploy your cdk stacks as you do and add `--output .cdkout` at end of cdk deploy command.   
    - Note: If you have already deployed stack with `--output .cdkout` command to get recomendations before, you can skip this step to get report for the same resources as earlier

    Sample command: 
    ```bash
    cdk deploy MyStack --parameters uploadBucketName=uploadbucket --output .cdkout
    ```
    - Note: After this step a .cdkout folder will be created with at the same level from where you ran your cdk deployment command.
 
3. Run recco command to fetch recommendations from cloudfix.
    ```bash
    cloudfix-linter cdk reco --aws-profile XYZ
    ```
    If non profile is mentioned cli will fetch recommendations for default aws-profile

4. CLI will fetch all the recommendations from cloudfix for the resources you deployed at step 1 and print it in console.

5. User can enable a json mode to get the output of CLI in a json format.

### Building and publising
See the Github action file for details.