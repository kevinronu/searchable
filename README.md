# Email Indexer

Email Indexer is a powerful email processing tool designed to seamlessly integrate with ZincSearch, a robust alternative to Elasticsearch. This project exemplifies the versatility of modern web development technologies, featuring a backend crafted in `Go` and `Chi`, and a frontend powered by `TypeScript`, `Vue 3`, `Pinia` for state management, and `Router 4` for navigation.

## Screenshot

![Searchable Screenshot](https://github.com/kevinronu/email-indexer/blob/main/screenshot.webp?raw=true)

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

6. **Insomnia Testing**:
   To explore and test the server and ZincSearch, import the [server_insomnia.json](https://github.com/kevinronu/email-indexer/blob/main/server-insomnia.json) and [zinc-insomnia.json](https://github.com/kevinronu/email-indexer/blob/main/zinc-insomnia.json) files into your Insomnia workspace.

With these steps completed, Email Indexer is now running on your local environment.

## Project Overview

Email Indexer is your solution for efficiently managing and querying email data. Here's what this project offers:

- **Effortless Email Processing**: Streamline the handling of email files and effortlessly send them to ZincSearch for indexing.

- **Modern Backend Infrastructure**: Built with Go and Chi, the backend ensures robust and efficient data processing, making it a reliable choice for email management.

- **Sleek Frontend Experience**: The user interface, developed using TypeScript and Vue 3, delivers a seamless and intuitive experience for interacting with your email data.

- **State-of-the-Art State Management**: Utilizing Pinia for state management, Email Indexer ensures your application remains responsive and performant.

- **Smooth Navigation**: Router 4 facilitates smooth navigation throughout the application, enhancing the user experience.

## Deploy with Terraform

You can easily deploy Email Indexer to AWS using Terraform. Follow these steps to set up and deploy the application:

1. **Install Terraform:**
   If you haven't already, you'll need to install Terraform. Visit the [Terraform](https://www.terraform.io) page and download the appropriate version for your operating system.

2. **Verify Terraform Installation:**
   To ensure Terraform is installed correctly, open your terminal and run:

   ```sh
   terraform --version
   ```

3. **AWS Credentials:**
   Make sure you have AWS credentials configured on your local machine. You can set up your AWS credentials by using the AWS CLI or AWS Toolkit
   extension for VSCode to creating a `~/.aws/credentials` file. Ensure your credentials have the necessary permissions to create the required AWS resources, I recommend `AdministratorAccess` policies.

4. **Customize Terraform Variables:**
   Open the `variables.tf` file in the Terraform directory and adjust the following variables to match your preferences:

   - **`host_os`**: Set to `"windows"` or `"unix"` depending on your host OS preference.
   - **`my_ip`**: Modify this to specify your public IP address or address range if needed. The default is "`0.0.0.0/0`" for unrestricted access.

   `Windows Users without WSL`: If you're using Windows without Windows Subsystem for Linux (WSL), ensure you customize the SSH configuration in windows-ssh-config.tpl. Replace `username` with your Windows username in the SSH configuration template.

5. **Deploy:**
   After customizing your Terraform variables, you can deploy Email Indexer to AWS by running:

   ```sh
   terraform apply
   ```

   Terraform will display a plan showing the resources it will create. Review the plan and type `yes` to confirm the deployment. Terraform will provision the necessary AWS resources.

6. **Access the Deployed Application:**
   Once the deployment is complete, Terraform will provide information about the resources it created, including the public IP address or DNS name of your deployed Email Indexer application and you can view the output `dev-ip`.

7. **Accessing the EC2 Instance:**
   Open your terminal and run the following command, replacing `190.190.190.190` with the public IP address or DNS name of your EC2 instance (`dev-ip`):

   ```sh
   sh -i ~/.ssh/mtckey ubuntu@190.190.190.190
   ```

   Now you are in the EC2 Instance terminal.

8. **GitHub Credentials Configuration:**
   Follow the steps mentioned here [Generating a new SSH key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent) then the steps mentioned here [Adding a new SSH](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account).

9. **Installation and Setup:**
   With GitHub credentials configured, you can proceed with the installation and setup of Email Indexer following the steps outlined in the [Installation and Setup](#installation-and-setup) section, but after `step 3` change the .env configuration according to the computational capacity of the EC2 instance (you can do it with the nano command), is mandatory to change the value of `DEPLOY_IP` to the value of `dev-ip`.

   ```sh
   nano .env
   ```

   Then continue with step 4. Ensure that you've completed all the necessary steps to set up and run Email Indexer successfully.

10. **Cleanup Resources:**
    If you want to tear down the deployed resources, you can run:

    ```sh
    terraform destroy
    ```

    Be cautious with this command, as it will remove all resources created by Terraform.

With these steps, you can easily deploy Email Indexer to AWS using Terraform. Enjoy using your email processing tool in the cloud!

## Connect with Me

LinkedIn: [Kevin Robles](https://www.linkedin.com/in/kevinronu/)
