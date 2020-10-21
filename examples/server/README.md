## Server example

This example will run a server on port 3000 with the given .env configuration. It will give you a quick view about how to handle the OAuth2 process as well an overview of all the different endpoints you can consume from Factorial app.

## How it works

We recommend you to start by reading the official documentation from Factorial https://docs.factorialhr.com/docs/getting-started and then start creating your own OAuth App in Factorial (https://docs.factorialhr.com/docs/create-a-new-oauth-application).

During the creation process you will be asked to introduce a callback url, that it should be https. In our case we use ngrok in order to expose our local server to the internet. If you want to download and play with ngrok please refer to the official documentation https://ngrok.com/.

Once you have you OAuth application correctly created in Factorial, you will need to create and fill a new **.env** file with the following information.

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

Then open your browser and navigate to http://localhost:3000, you should see the factorial button, if you click it the OAuth2 handling process will start.
