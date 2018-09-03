# Serverless development with Go and AWS SAM

A workshop.


## brief

The goal is to create a serverless implementaion of the unix command `wc`, using
Go and AWS SAM. Test locally and deploy to AWS. Count words and lines to match
the equivalent output of `wc -w` and `wc -l`. Build in multiple iterations,
testing and deploying at each phase.


## disclaimer

If you're using Windows you're gonna have to figure out a bunch of stuff for
yourself. The deployment scripts are written in bash.


## what you need

- [go](https://golang.org/doc/install), prefer `1.11` for module support.
- [docker](https://www.docker.com/get-started), required for `sam local`.
- [aws cli](https://aws.amazon.com/cli/), installed and [configured](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).
- [sam cli](https://docs.aws.amazon.com/lambda/latest/dg/sam-cli-requirements.html), for hopefully obvious reasons.
- [curl](https://curl.haxx.se/), or an alternative for testing.


## prep

Create an S3 bucket to store the deployment package.

```
aws s3 mb s3://some-bucket
```

Clone this repo probably.

```
git clone https://github.com/dedelala/go-sam-workshop.git
cd go-sam-workshop
```

If you have Go 1.11 you can clone anywhere, otherwise you will need to use the
Go workspace. Version 1.11 is preferable.


## template

Create an infrastructure template, `template.yml`, with a *serverless function*!
The function needs an *API event* with a root endpoint (`/`) that accepts *POST* requests.
The template needs a lambda execution *IAM role* that allows logging.  It is helpful
for the template to have an output for the resulting API endpoint.

**Tip**: searching for `AWS::Serverless::Function` will get you going.

**Skip**: for the completed template, `git checkout step-template`.


## handler

Write a handler, `main.go`, that responds with a 200 status (and a message if
you like).  The `main` method needs to call `lambda.Start` with a handler
method.  The handler method accepts an `events.APIGatewayProxyRequest` and
returns an `events.APIGatewayProxyResponse` and an `error`.

**Tip**: `lambda` and `events` are sub-packages of `github.com/aws/aws-lambda-go`.

**Skip**: for the completed handler, `git checkout step-handler`.


## local

Write a script that builds the handler and runs it locally.  Build a
*static*, *linux*, *amd64* binary named `wc` to match the template.  Use of Go modules
is preferred.  The binary must be zipped and named `wc.zip` to match the
template.  Then run `sam local start-api`

**Skip**: for the completed script, `git checkout step-local`.

Run it!

```
./local.sh
```

Test with another prompt.

```
curl -X POST http://localhost:3000
```

Stop `local.sh` with `Ctrl+C`.


## deploy

Write a script that builds the handler, packages and deploys it using sam cli.
Accept a bucket name and stack name as parameters.  For bonus points the script
will check the system for its dependencies. Finally, if the stack template
outputs the API endpoint the script should get and print it.

**Tip**: the build and zip steps are the same as `local.sh`.

**Tip**: the relevant `sam` commands are `package` and `deploy`.

**Skip**: for the completed script, `git checkout step-deploy`.

Run it!

```
./deploy.sh <some-bucket> <stack-name>
```

Test it!

```
curl -X POST <url-goes-here>
```


## words

Modify the handler to count the number of words in the request body.  The
response must be json encoded.

```
# example
curl -X POST --data-binary 'a b c d e f g' http://localhost:3000
# output: {"words":7}
```

**Tip**: package `unicode` is your friend.

**Tip**: package `json` is used for json encoding. Producing the result in any
other way will be frowned upon.

**Skip**: for the completed handler, `git checkout step-words`

Test the changes locally.

```
./local.sh
```

Count the words in `main.go`.

```
curl -X POST --data-binary @main.go http://localhost:3000
```

Check the result!

```
wc -w main.go
```

Deploy the modified handler and test it live!


## lines

Modify the handler to count lines as well as words. Add endpoints to the
handler, `/words` for the word count, `/lines` for the line count. Have `/`
return both!

**Tip**: adding endpoints requires modifications to `template.yml`.

**Skip**: for the completed handler, `git checkout step-lines`

Local, test, deploy, test, drop down, increase spead and reverse direction!


## congratulations

You made a thing!


## clean up

The web console will be sufficient for a quick clean up.

- Delete any Cloudformation stacks that were created.
- Delete the S3 bucket.

