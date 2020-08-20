# CloudRun + Google ko

This is a companion repo to the [Article Written on Medium](https://medium.com/@peter.malina/noops-go-on-cloud-run-689d92215c5c?sk=1b5e8f716686ddffa1b73c4a652b84d1).

## Installing Tools

There are two main prerequisites (not counting Go itself) when running this setup:
- Google ko (follow the command below)
- [Google Cloud SDK](https://cloud.google.com/sdk) (follow their instructions for setup)

```
# install Google ko
GO111MODULE=on go get github.com/google/ko/cmd/ko
```

## Getting Started

First, build the container using Google `ko`:
```shell script
KO_DOCKER_REPO=eu.gcr.io/<your-project> ko publish .

# or you can capture the image name by getting the last line of the output
APP_IMAGE=$(KO_DOCKER_REPO=europe-west3-docker.pkg.dev/petermalina/test ko publish . | tail -1)
```

Deploy the image created and published by `ko`:
```shell script
gcloud run deploy cloudrun-ko-example \
                  --image=$APP_IMAGE \
                  --region=europe-west1 \
                  --platform=managed \
                  --allow-unauthenticated
```