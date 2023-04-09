<h1>GoLang Discord Bot</h1>

This GoLang Discord Bot project is designed to help users interact with a Discord server. It uses several dependencies such as aws-lambda-go, discordgo, uuid, wire, testify, and exp.

<h2>Getting Started</h2>

To get started with this project, follow these steps:

<ol>
  <li>Clone the repository</li>
  <li>Ensure that all dependencies are installed</li>
  <li>Build and run the code</li>
</ol><h3>Prerequisites</h3>
<ul>
  <li>Go v1.13 or later</li>
  <li>Discord server and API key</li>
  <li>AWS account (optional)</li>
</ul><h3>Installation</h3>
To install the dependencies, run the following command:

``` shell
go mod download
```
<h3>Building and Running</h3>
To build and run the code, execute the following command:

``` shell
go build -o bot ./main.go
./bot
```
<h2>Features</h2>
<ul>
  <li>Moderation of messages</li>
  <li>Responding to certain commands</li>
  <li>Automatic generation of unique user IDs using uuid</li>
  <li>Dependency injection using wire</li>
</ul><h2>License</h2>
This project is licensed under the MIT License.
<h1>Infrastructure</h1>
The infrastructure directory contains the necessary files to deploy the GoLang Discord Bot on the AWS cloud platform using Pulumi. It has several dependencies such as aws, aws-native, awsx, command, and pulumi.
<h2>Getting Started</h2>
To get started with this project, follow these steps:
<ol>
  <li>Clone the repository</li>
  <li>Ensure that all dependencies are installed</li>
  <li>Configure the AWS account credentials</li>
  <li>Deploy the infrastructure</li>
</ol><h3>Prerequisites</h3>
<ul>
  <li>Pulumi v3.0 or later</li>
  <li>AWS account and API key</li>
</ul><h3>Installation</h3>
To install the dependencies, run the following command:

``` shell
npm install
```

<h3>Configuration</h3>

To configure the AWS account credentials, run the following command:

``` shell
pulumi config set aws:accessKeyId <ACCESS_KEY_ID>
pulumi config set aws:secretAccessKey <SECRET_ACCESS_KEY> --secret
```
    
<h3>Deployment</h3>

To deploy the infrastructure, execute the following command:

``` shell
pulumi up
```
<h2>License</h2>

This project is licensed under the MIT License.

<h2>Contributing</h2>
This project is on hold.