# Email Indexer

Email Indexer is a powerful email processing tool designed to seamlessly integrate with ZincSearch, a robust alternative to Elasticsearch. This project exemplifies the versatility of modern web development technologies, featuring a backend crafted in `Go` and `Chi`, and a frontend powered by `TypeScript`, `Vue 3`, `Pinia` for state management, and `Router 4` for navigation.

## Installation and Setup

1. **Clone this Repository:**
   Begin by cloning this repository to your local machine using Git:

   ```sh
   git clone git@github.com:kevinronu/email-indexer.git
   cd email-indexer
   ```

2. **Check Your Docker Version**:
   Make sure you have Docker installed. You can verify your Docker installation by running:

   ```shell
   docker -v
   ```

   If Docker is not installed, you can download it [<u>here</u>](https://www.docker.com/products/docker-desktop/).

3. **Copy Configuration File:**
   Duplicate the `*.env.example` file and rename it to `.env`:

   ```sh
   cp .env.example .env
   ```

4. **Download Email Data:**
   Download the Enron Corp email files and prepare them for processing. Create an `emails` folder and extract the data:

   ```sh
   curl -OL http://www.cs.cmu.edu/\~enron/enron_mail_20110402.tgz
   mkdir emails
   tar zxvf enron_mail_20110402.tgz -C emails
   ```

5. **Run the application:**
   Launch the Email Indexer application using Docker Compose:
   ```sh
   docker compose up
   ```

With these steps completed, Email Indexer is now running on your local environment..

## Project Overview

Email Indexer is your solution for efficiently managing and querying email data. Here's what this project offers:

- **Effortless Email Processing**: Streamline the handling of email files and effortlessly send them to ZincSearch for indexing.

- **Modern Backend Infrastructure**: Built with Go and Chi, the backend ensures robust and efficient data processing, making it a reliable choice for email management.

- **Sleek Frontend Experience**: The user interface, developed using TypeScript and Vue 3, delivers a seamless and intuitive experience for interacting with your email data.

- **State-of-the-Art State Management**: Utilizing Pinia for state management, Email Indexer ensures your application remains responsive and performant.

- **Smooth Navigation**: Router 4 facilitates smooth navigation throughout the application, enhancing the user experience.

## Connect with Me

LinkedIn: [Kevin Robles](https://www.linkedin.com/in/kevinronu/)
