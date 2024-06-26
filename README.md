# infrastructure-hw

## Your mission

For this assignment, you're going to create the infrastructure for an application with a small set of services.

- One service needs to broadcast `Hello world` at random intervals. Make the interval anywhere from 1 to 10 seconds, with each the time until the next broadcast each chosen randomly.

- Another service needs to receive the `Hello world` broadcasts.

- Then a user should be able to view the `Hello world` broadcasts, as they arrive, from a web browser.

### Other requirements

- Use whatever languages and frameworks you want to create the services.
- We're aiming to just run this application on an engineer's local machine, not the cloud; design your solution for `minikube`
- Your solution should have the minimum number of manual setup steps necessary.
- Use any adjacent infrastructure tools you think make for a more elegant solution.

## Submission

- Fork this repository on GitHub. Develop a solution on your fork. Extra points for good git hygiene.
- Include specific instructions in your README about pre-requisites and setup steps. Another engineer should be able to go from zero to running your solution on their local machine.
- Either send us the link to your repository (if you make it public) or email us a zipped-up folder.

## Implementation
### Minikube

Ensure Minikube cluster is running with `minikube start` and Docker Hub credentials are cached in `~/.docker/config.json` via `docker login`

To deploy the Kubernetes manifests, port-forward, and view the output in the browser, run:
 `make deploy-k8s`

To kill the port forward running in the background: `make stop-portforward`


### Local development with Docker Compose

Build broadcast and browser service images:

`make build-broadcast`

`make build-browser`

Run services, NGINX, and Redis:

`make up`

Open browser and navigate to [localhost:8000](http://localhost:8000)

To gracefully shutdown services run `make down`

