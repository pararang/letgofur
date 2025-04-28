# letgofur

letgofur (Leutenant Gofurr) is Captain Rover subordinat.

## About

letgofur provides a simple and efficient way to interact with your CapRover server from the command line. It allows you to:

- List all applications deployed on your CapRover instance
- Connect to your CapRover server using authentication credentials
- Manage your CapRover applications without using the web interface

Built with Go and the forked [GoCaproverAPI](https://github.com/ErSauravAdhikari/GoCaproverAPI) as its base, letgofur makes it easy to automate and streamline your CapRover management tasks on your terminal.

## Installation

### Prerequisites

- Go 1.23 or higher

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

You can configure letgofur using environment variables or command-line flags.

### Environment Variables

Create a `.env` file in the project root with the following variables:

```
URL=https://captain.your.domain
PASSWORD=yourpassword
```

You can use the provided `.env.example` as a template:

```bash
cp .env.example .env
# Edit .env with your credentials
```

### Command-line Flags

Alternatively, you can provide credentials directly via command-line flags:

```bash
letgofur --host https://captain.your.domain --passwd yourpassword
```

## Usage

### List all applications

```bash
letgofur --host https://captain.your.domain --passwd yourpassword ls
```

Or if you've configured the `.env` file:

```bash
letgofur ls
```

### Generate YAML file for an app

Generate a YAML representation of an app's definition:

```bash
letgofur --host https://captain.your.domain --passwd yourpassword generate-yml app-name
```

You can specify a custom output file:

```bash
letgofur generate-yml app-name --output ./configs/app-name.yml
```

Aliases: `gen-yml`, `yml`

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

- **App Management**
  - Create new applications (`CreateApp`)
  - Remove/delete applications (`RemoveApp`)
  - Force build applications (`ForceBuild`)
  - Update application details and configurations

- **Domain Management**
  - Add custom domains to applications (`AddCustomDomain`)
  - Enable SSL for base domains (`EnableBaseDomainSSL`)
  - Enable SSL for custom domains (`EnableCustomDomainSSL`)

- **Resource Management**
  - Update resource constraints (memory, CPU) for applications
  - Scale application instances

- **Deployment Options**
  - Support for persistent data applications
  - Configure environment variables
  - Set up port mappings

- **User Interface Improvements**
  - Interactive mode for commands
  - Progress indicators for long-running operations
  - Colorized output for better readability

## License

This project is licensed under the Apache License, Version 2.0 (the "License"). You may obtain a copy of the License at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](LICENSE) file for the specific terms and conditions of the license.
