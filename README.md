# Change Operator

Thanks for checking out Change Operator! This is a Golang application stack that is useful for providing 
Change Awareness in many ways. Within this stack is an API; a Postgres DB; and a Slackbot to help
you get started. 

There's many ways to do such a setup, but this one is light and simple! If the Slackbot
is not needed, then just hit the Change Operator endpoint. 

Lastly, I recommended using something like `Grafana` to view these changes. The `postgres` database
stores a timestamp of the Change which can be plotted and searched in many interesting ways!


# Setup

This setup is primarily for local testing, but can easily be configured in a cluster, web-service, etc.

## If you want to use the Slackbot...

For local usage, you will need a service to expose local ports for Slack to hit -- `ngrok` is a great option!

https://ngrok.com/

The port you expose will need to be used for the `changebot` port. The address of the exposed port
will need to be used for the Slack App setup.

Next, you'll need to configure your own Slack App. There's many resources to help with this, for example:

https://api.slack.com/authentication/basics

Specifically what you'll need to configure under your Slack App is:

**Event Subscriptions** something like...

`<your-exposed-port-addres>/events`

**Slash Commands** something like...

`<your-exposed-port-addres>/<command-name>`

**Interactivity & Shortcuts** something like...

`<your-exposed-port-addres>/interactive-endpoint`

## I don't want to use the Slackbot...

Simply delete or comment out the `changebot` code-block in the `docker-compose.yaml` file. This means 
you will use something like `Postman` to send a raw JSON.

## Docker-Compose file...

Open and edit the `docker-compose.yaml` file and fill out the necessary lines. Comments were left behind
and will need to be replaced for this to work.

Once that is complete, you can run:

`docker-compose -f docker-compose.yaml up`

## Postgres Table

The `create_tables.sql` can be used, removed or edited. The purpose of it is to  help setup the Database and Table quickly.

