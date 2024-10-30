# Caldera Docker image

Unfortunately,
[Caldera](https://caldera.readthedocs.io/en/latest/Installing-Caldera.html#docker-deployment)
does not provide a prebuilt Docker image. Installation instructions include cloning the entire
repository and building the Docker image locally.

Users of SOARCA may run `make caldera` to build a Caldera image locally, which can subsequently
be used within the Docker Compose setup.
