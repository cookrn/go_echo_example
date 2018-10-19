# Go Echo Example

A tiny repository to play with Go, its standard libraries, the Echo framework,
and testing.

## Endpoints

### GET /

The index page is found in the `public/` directory and creates a simple form
that allows one to upload a file.

### POST /uploads

This endpoint generates a random id for the upload as well as calculating the
sha1 digest of the uploaded file and returns said data to the user.

## Tests

One test uses the example set by the Echo framework tests to create a fake
upload and checks the return data.
