## Persistence example

This example will run a server on port 3000 with the given .env configuration. It will give you a quick view about how to handle the OAuth2 process and how to persist the given tokens using a **tokenRepository** implemented in memory.

## How it works

We recommend you to start by reading the official documentation from Factorial https://docs.factorialhr.com/docs/getting-started and then start creating your own OAuth App in Factorial (https://docs.factorialhr.com/docs/create-a-new-oauth-application).

During the creation process you will be asked to introduce a callback url, that it should be https. In our case we use ngrok in order to expose our local server to the internet. If you want to download and play with ngrok please refer to the official documentation https://ngrok.com/.

Once you have your OAuth application correctly created in Factorial, you will need to create and fill a new **.env** file with the following information.

```
CLIENT_ID="---- Your client ID ----"
CLIENT_SECRET="--- Your client secret ---"
SCOPES="read,write"
REDIRECT_URL="--- Your redirect url, sample (https://7cad0b374498.ngrok.io/auth/factorial/callback) ----"
```

Once you have this set up, you can start the server using

```
go run *.go
```

Then open your browser and navigate to http://localhost:3000, you should see the Factorial button, if you click it the OAuth2 handling process will start.

Following the documentation, the given token has an expiration time of 2 hours, so after two hours our persisted client will refresh the token and update the given token information into our in memory implementation. This is a proposal of how you can persist this token, but you can build it on your own way and use the provider and client as we use on the **server example**