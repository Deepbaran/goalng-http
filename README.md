# Golang Web
This repository contains basic http request and response handling notes without frameworks.

---
## Links
- https://www.youtube.com/playlist?list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v

---

## FAQ
-  What is a Web App?
   -  A Web Application is a piece of software that can be accessed from a web browser such as Chrome, Firefox and Safari.
-  Web App and Web Client:
   -  We call the web app, the web server and the browser, the web client.
   -  A web client can be a mobile app, a desktop app, etc. Anything that is trying to access whatever web app we are trying to build.
   -  The Client is trying to talk to the server in some form of communication, like
      -  HTTP Request [most common]
         -  GET: Retrieve info from the server
         -  PUT: Create/Update resource in the server
         -  POST: Update a resource in the server
         -  DELETE: Delete resource from the server
      -  FTP Request
      -  WebSocket
-  HTTP Request
   -  HTTP Request and HTTPS request
      -  The difference between the two is that, in HTTPS, we are making the data that we are sending along with the data secure(encrypted) before sending the request to the destination server.
      -  To get the data in the original form, the server must have the private key.
      -  And hopefully, any middleman, like the router, cellular career, internet provider does not have access to the same key in order to unencryp the data.
   -  DNS Server
      -  The domain in the request gets translated by the DNS server to get the server IP address for the original request to go to.
   -  Filename
      -  Once into the server, we will need the specific path in the server where the file in the server that we are specifically looking for.
-  Response
   -  After we send a Request, hopefully we will get a response.
   -  As Web Application developers, constructing the response depending on the Request is our job.
   -  In the response we can first see the status.
   -  HTTP Status Codes
      -  1xx - Informational
      -  2xx - Success
      -  3xx - Redirection
      -  4xx - Client Error
      -  5xx - Server Error
   -  Body
      -  This holds the info we want to send back as response. Such as:
         -  HTML
         -  JSON