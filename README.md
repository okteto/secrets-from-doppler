# Integrate Doppler and Okteto pipelines together

This demo shows you how bring secrets defined in [Doppler](https://doppler.com) into your Okteto development environment. 

## Setup
1. Create a Doppler Project, and add two secrets with the name `MY_NAME` and `MY_COLOR`
1. Create a Doppler CLI Service Token ([instructions available here](https://docs.doppler.com/docs/enclave-service-tokens#cli-create-service-token))
1. Once created, add it to your Okteto account as a secret, using the name "DOPPLER_TOKEN"

## Launch your Development Environment

Once the project and secrets have been configured, deploy your development environment in Okteto Cloud (or your own Okteto Enteprise instance).

