# letgofur

letgofur (Lieutenant Gofurr) is a CapRover (Captain Rover) subordinate.

Simple CLI for CapRover API, built with Go. Help you to define your infrastructure hosted with CapRover in a declarative way. It is designed to be a simple and efficient command-line interface for managing your CapRover applications and resources.
It allows you to easily interact with your CapRover server, manage applications, and automate tasks without the need for a web interface.

## About

letgofur provides a simple and efficient way to interact with your CapRover server from the command line. It allows you to:

- List all applications deployed on your CapRover instance
- Connect to your CapRover server using authentication credentials
- Manage your CapRover applications without using the web interface

Built with Go and the forked [GoCaproverAPI](https://github.com/ErSauravAdhikari/GoCaproverAPI) as its base, letgofur makes it easy to automate and streamline your CapRover management tasks on your terminal.

## Installation

### Prerequisites

- Go 1.20 or higher

### Building from source

```bash
# Clone the repository
git clone https://github.com/pararang/letgofur.git
cd letgofur

# Build the application
go build -o letgofur

# Make it executable (optional)
chmod +x letgofur

# Move to a directory in your PATH (optional)
mv letgofur /usr/local/bin/
```

## Configuration

### Command-line Flags

You have provide credentials directly via command-line flags:

```bash
letgofur --host https://captain.your.domain --passwd yourpassword
```

## Usage

### List all applications

```bash
letgofur --host https://captain.your.domain --passwd yourpassword ls
```

### Create a workspace
Initialize a workspace for infrastructure as code configuration:

```bash
letgofur --host https://captain.your.domain --passwd yourpassword init
```

This command will create a directory named based on the hostname of your CapRover instance. Inside this directory, you will find all the current apps config. Currently this only supports the instance and app resource configurations.

```yaml
# Example of the generated YAML file
# captain.your.domain/app-name.yml
AppName: app-name
Instances: 3
Resources:
    Limits:
        MemoryBytes: 16777216
        NanoCPUs: 1000000
    Reservations:
        MemoryBytes: 1122323
        NanoCPUs: 1000000
```

### Update app configuration

Apply configuration changes to an existing app using a YAML file. Lets say you are inside the generated workspace directory:

```bash
letgofur --host https://captain.your.domain --passwd yourpassword apply app-name.yml
```

This command updates app resources and instance count based on the configuration file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Attribution

The `crapi` directory contains code from the [GoCaproverAPI](https://github.com/ErSauravAdhikari/GoCaproverAPI) project, which is licensed under the Apache License 2.0. The original code has been incorporated into this project with minimal modifications to support the letgofur CLI functionality.

Changes made to the original code include:
- Minor adaptations for integration with the letgofur command structure
- Remove some printed messages for cleaner CLI output

All copyright notices and license terms from the original project have been preserved in the source files.

## To Do

The following features are planned for future releases, based on the capabilities of the GoCaproverAPI:

- [ ] **App Management**
  - [x] List all applications
  - [x] Generate workspace for infra as code configuration
  - [x] Update application details and configurations
  - [ ] Create new applications 
  - [ ] Remove/delete applications 
  - [ ] Force build applications 

- [ ] **Domain Management**
  - [ ] Add custom domains to applications 
  - [ ] Enable SSL for base domains 
  - [ ] Enable SSL for custom domains 
  - [ ] Enable force redirect to the custom domain

- [x] **Resource Management**
  - [x] Update resource constraints (memory, CPU) for applications
  - [x] Scale application instances

- [ ] **Deployment Options**
  - [ ] Configure environment variables
  - [ ] Set up port mappings

- [ ] **User Interface Improvements**
  - [ ] Interactive mode for commands
  - [ ] Progress indicators for long-running operations
  - [ ] Colorized output for better readability
  - [ ] Interective mode with session

## License

This project is licensed under the Apache License, Version 2.0 (the "License"). You may obtain a copy of the License at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](LICENSE) file for the specific terms and conditions of the license.
