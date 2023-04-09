import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsnative from "@pulumi/aws-native";
import {local} from "@pulumi/command";
import {Input, Output, OutputInstance} from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

export function deployLambdas(
    // vpcConfig: Output<FunctionVpcConfig>
): OutputInstance<string> {

  const repo = new aws.ecr.Repository("fsn-discord-bot");

  const role = new aws.iam.Role("fsnDiscordBotRole", {
    assumeRolePolicy: aws.iam.assumeRolePolicyForPrincipal({Service: "lambda.amazonaws.com"}),
    inlinePolicies: [
      {
        name: "my_inline_policy",
        policy: JSON.stringify({
          Version: "2012-10-17",
          Statement: [{
            Action: [
              "ec2:DescribeInstances",
              "ec2:CreateNetworkInterface",
              "ec2:AttachNetworkInterface",
              "ec2:DescribeNetworkInterfaces",
              "ec2:DeleteNetworkInterface"
            ],
            Effect: "Allow",
            Resource: "*",
          }],
        }),
      },

    ],
  });
  new aws.iam.RolePolicyAttachment("lambdaFullAccessDiscordBot", {
    role: role.name,
    policyArn: aws.iam.ManagedPolicy.AWSLambdaExecute,
  });

  // Get the repository credentials we use to push to the repository
  const repoCreds = repo.registryId.apply(async (registryId) => {
    const credentials = await aws.ecr.getCredentials({
      registryId: registryId,
    });
    const decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
    const [username, password] = decodedCredentials.split(":");
    return { server: credentials.proxyEndpoint, username, password };
  });

  const image = new docker.Image("fsn-discord-bot-image", {
    imageName: repo.repositoryUrl,
    build: "../",
    registry: repoCreds,
  })

  const lambda = new aws.lambda.Function("fsn-discord-bot-lambda", {
    packageType: "Image",
    imageUri: image.imageName,
    role: role.arn,
    timeout: 10,
    memorySize: 64,
    reservedConcurrentExecutions: 1,
  });

  const lambdaUrl = new awsnative.lambda.Url("fsn-discord-bot-url", {
    targetFunctionArn: lambda.arn,
    authType: awsnative.lambda.UrlAuthType.None,
  });

  const awsCommand = new local.Command("aws-command", {
    create: pulumi.interpolate`aws lambda add-permission --function-name ${lambda.name} --action lambda:InvokeFunctionUrl --principal '*' --function-url-auth-type NONE --statement-id FunctionURLAllowPublicAccess`
  }, {deleteBeforeReplace: true, dependsOn: [lambda]})

  let functionUrl: OutputInstance<string> = lambdaUrl.functionUrl;

  return functionUrl;
}
