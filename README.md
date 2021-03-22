# uploadToFirestore

A simple golang application which reads in names stored in a CSV format and uploads it to selected FireStore DB.

Presumes `GOOGLE_APPLICATION_CREDENTIALS` is set and points to Service Account Key in JSON format

## Deploy

Deployments can be run via script ./deploy.sh. Expects `GOOGLE_SA_KEY` to be set
