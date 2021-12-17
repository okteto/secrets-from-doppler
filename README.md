# Integrate Doppler and Okteto pipelines together

This demo shows you how bring secrets defined in [Doppler](https://doppler.com) into your Okteto development environment. 

## Setup
1. Create a Doppler Project, and add two secrets with the name `MY_NAME` and `MY_COLOR`
1. Create a Doppler CLI Service Token ([instructions available here](https://docs.doppler.com/docs/enclave-service-tokens#cli-create-service-token))
1. Once created, add it to your Okteto account as a secret, using the name "DOPPLER_TOKEN"

## Launch your Development Environment

Once the project and secrets have been configured, deploy your development environment in Okteto Cloud (or your own Okteto Enteprise instance).

Go to the endpoints in the dashboard, and click on it. You'll see something like this:

```
Hi, my name is Cindy, and my favorite color is Blue
```

Go to the [pipeline manifest](.okteto/okteto-pipeline.yml) and notice how the secrets are being pulled from your Doppler project via the  `doppler secret get` command and injected directly into the helm chart deployment.

 If you go to Doppler, change the values of the secrets, and redeploy your development environment, you'll see the changes immediately.

## Access the secrets while coding

From your local machine, run `okteto up` after deploying your development enviroment. Now run `env` to dump all the environment variables. Notice how `MY_NAME` and  `MY_COLOR` are available? This is one of the benefits of using Okteto's development environments: the ability to use the same configuration and tools that you do in production (like secret vaults) while you develop.