# TimeTracker API

The TimeTracker API is an application built using [GoFiber](https://github.com/gofiber/fiber), a lightning-fast web framework for Go programming language. Designed to cater to individuals, freelancers, and businesses, the TimeTracker API empowers users to efficiently manage and track their time across various projects and clients.

This project can be useful for developers who are new to GoFiber and interested in understanding how to architect and develop a web API using [GoFiber](https://github.com/gofiber/fiber), [Docker](https://www.docker.com/), and [Kubernetes](https://kubernetes.io/).

> Notice: This project is currently in the development phase and is not yet suitable for production use.

## Features

1. User Authentication with JWT.
2. Organization Management.
3. Project and Client Management.
4. Time Logging.

## Getting Started

### Dependencies

Make sure you have the following dependencies installed and configured on your machine.

1. [Go](https://go.dev/doc/install) version 1.20 or higher.
2. [Docker](https://www.docker.com/).
3. [Kind](https://kind.sigs.k8s.io/) in order to run the application in Kubernetes cluster.
4. [pgAdmin](https://www.pgadmin.org/) (optional).

### Download

To obtain your copy of the project, you have several options:

1. Download .zip File:
   Click on the "Download .zip" option provided in the repository. This will download the entire project as a compressed zip file to your local machine. You can then extract the contents and explore the project.

2. Clone the Repository:
   Clone the repository to your local machine using Git. Open your terminal or command prompt, navigate to your desired location, and execute the following command:

   ```
   git clone <repository_url>
   ```

   Replace `<repository_url>` with the URL of the repository. This will create a local copy of the project on your machine, allowing you to make changes and contribute back to the repository if desired.

3. Fork the Repository:
   Forking the repository will create a personal copy of the project in your GitHub account. Click on the "Fork" button at the top right corner of the repository page. This will duplicate the entire repository under your GitHub account.

Choose the option that best suits your needs and start exploring, modifying, or contributing to the project as you see fit. Happy coding!

### Run

Once you have the project downloaded on your machine you can start and test the application. The application comes with a `makefile` which contains many helpful commands to build and run the containerized application in a Kubernetes cluster with Kind.

1. Create the cluster

   make kind-up

2. Apply the Kubernetes configurations

   make kind-apply

3. Build the Docker image and load it to the cluster

   make build
   make kind-load

4. See the logs

   make kind-logs

5. Check the status of your cluster

   make kind-status

> Please refer to the [makefile](https://github.com/mahdi-mk/time-tracker-api/blob/main/makefile) to learn more about the useful commands available.

> If you are on a UNIX based OS you may need to change some of the commands in the `makefile` because they have been tested on a Windows machine.

## Resources

After successfully installing and setting up the project, you can test the API endpoints using [Postman](https://www.postman.com/downloads/), a popular API development and testing tool. To get started, follow these steps:

1. Download the Postman Collection, a comprehensive set of pre-configured API endpoints, by navigating to docs directory of this project.

2. Import the downloaded Postman Collection into your Postman application. This will configure all the necessary API endpoints, saving you valuable time and effort.

With the Postman Collection readily available, you can explore and interact with all the API endpoints, enabling you to thoroughly test the functionality.

## Contributions

Contributions are more than welcome. If you have any feature suggestions, bug fixes, or new ideas to enhance the project, please follow these simple instructions:

1. Fork this repository by clicking on the "Fork" button at the top right corner of the page. This will create a personal copy of the project in your GitHub account.
2. Make the necessary changes or additions to your forked repository.
3. Once you are satisfied with your changes, open a pull request (PR) by navigating to the original repository and clicking on the "New Pull Request" button. Provide a clear and detailed description of the changes you made in your PR.
