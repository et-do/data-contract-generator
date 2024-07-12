
# Data Contract Generator (YAML -> SVG)

A Go-based tool to convert YAML data contracts into visually appealing SVG files.

## Features
- Parses YAML data contracts.
- Generates SVG files with formatted and readable content.
- Supports section headers, bullets, and text wrapping for long descriptions.

## Purpose
Enables data contracts to be defined in .yaml for easy usability within a codebase, while allowing for the conversion of those contracts to more client-friendly for shareout, validation, and buy-in.

## Directory Structure
```
data-contract-generator/
├── main.go
├── models/
│   ├── data_contract_test.go
│   └── data_contract.go
├── utils/
│   ├── file_processor.go
│   ├── svg_generator_test.go
│   ├── svg_generator.go
│   ├── yaml_parser_test.go
│   └── yaml_parser.go
├── testdata/
│   └── test_contract.yaml
└── .devcontainer/
    ├── Dockerfile
    └── devcontainer.json
```

## Usage

1. **Clone the repository**:
    ```sh
    git clone https://github.com/et-do/data-contract-generator.git
    cd data-contract-generator
    ```

2. **Build the binary**:
    ```sh
    go build -o data-contract-generator main.go
    ```

## Usage

1. **Source your YAML files**:
   - Place your YAML files in any directory of your choice.

2. **Run the tool**:
    ```sh
    ./data-contract-generator ./directory_with_yaml_files
    ```

   - This will generate SVG files in the same directory as the YAML files.

## Example
### Input YAML
Here is an example of a YAML data contract:

```yaml
name: Customer Data Contract
date: 2024-07-12
data_source: CRM System
table_name: customer_data
expected_schema:
  - name: customer_id
    type: STRING
  - name: first_name
    type: STRING
  - name: last_name
    type: STRING
  - name: email
    type: STRING
  - name: phone_number
    type: STRING
  - name: signup_date
    type: DATE
  - name: last_purchase_date
    type: DATE
  - name: total_spent
    type: FLOAT
frequency: daily
PII_sensitive: true
description: >
  This data contract specifies the schema and properties for the customer data
  collected from the CRM system. The data includes personal identifiable
  information (PII) and is updated daily.
contact_person:
  name: Jane Doe
  email: jane.doe@example.com
  phone: "+1-555-123-4567"
data_retention_policy: 3 years
security_classification: confidential
version: 1.0
```

### SVG Output:
This will generate an SVG file named with the content formatted for readability

![Example Output](/static/contract.png)

## Development Environment

This project includes a development environment configuration for Visual Studio Code with Docker.

1. **Setup Dev Container**:
   - Open the project in Visual Studio Code.
   - Use the "Remote - Containers" extension to open the project in a container.

## Contributing

1. **Fork the repository**.
2. **Create a new branch**:
    ```sh
    git checkout -b feature-branch
    ```
3. **Make your changes**.
4. **Commit your changes**:
    ```sh
    git commit -m "Description of your changes"
    ```
5. **Push to the branch**:
    ```sh
    git push origin feature-branch
    ```
6. **Create a Pull Request**.

## License

This project is licensed under the MIT License.
