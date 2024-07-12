
# Data Contract Generator (YAML -> SVG)

A Go-based tool to convert YAML data contracts into visually appealing SVG files.


## Data Contracts

A data contract is a formal agreement between different parties about the structure, format, and semantics of the data being exchanged. It defines how data should be produced, processed, and consumed, ensuring consistency and quality across different systems and teams.

**Key Components of a Data Contract**
- **Name**: The name of the data contract.
- **Date**: The date the data contract was created or last updated.
- **Data Source**: The source system or application where the data originates.
- **Table Name**: The name of the table or dataset where the data is stored.
- **Expected Schema**: A detailed definition of the schema, including field names and data types.
- **Frequency**: How often the data is expected to be updated or refreshed.
- **PII Sensitive**: Indicates whether the data contains Personally Identifiable Information (PII).
- **Description**: A detailed description of the data and its purpose.
- **Contact Person**: Information about the person responsible for the data.
- **Data Retention Policy**: The policy for how long the data should be retained.
- **Security Classification**: The security classification of the data (e.g., confidential, public).
- **Version**: The version of the data contract.

**Benefits of Data Contracts**
- **Data Quality**: Data contracts enforce a standardized schema and data format, reducing errors and inconsistencies. This ensures that data is clean, accurate, and reliable.
- **Pipeline Development**: With clearly defined data structures and expectations, development of data pipelines becomes more straightforward. Developers can build and maintain ETL (Extract, Transform, Load) processes with confidence that the data will adhere to the agreed-upon contract.
- **Clear Communication**: Data contracts serve as a single source of truth, providing clear documentation and communication between different teams (e.g., data producers, data consumers, and data engineers).
- **Compliance and Security**: By defining data retention policies and security classifications, data contracts help organizations comply with legal and regulatory requirements, and protect sensitive information.
- **Change Management**: When data contracts include versioning, it becomes easier to manage changes over time. Teams can track updates and ensure backward compatibility.

**Using Data Contracts in a Codebase**
- **Schema Validation**: Implement automated schema validation to ensure incoming data adheres to the defined contract. This can be done using tools or libraries that compare incoming data to the expected schema and raise alerts or errors for any discrepancies.
- **Data Ingestion**: Use data contracts to guide the ingestion process. When setting up ETL pipelines, reference the data contract to understand the schema, data types, and transformation rules required.
- **Data Quality Checks**: Incorporate data quality checks as part of your data processing pipelines. Use the data contract to define and enforce rules for data completeness, accuracy, and consistency.
- **Documentation and Onboarding**: Data contracts serve as comprehensive documentation for your datasets. New team members can refer to data contracts to quickly understand the structure, purpose, and usage of the data.
- **Monitoring and Alerting**: Set up monitoring and alerting based on the data contract. For example, if the data is expected to be updated daily, set up alerts for missing or delayed updates.
- **Version Control**: Maintain version control for data contracts. When a data contract is updated, ensure that all dependent systems and pipelines are compatible with the new version. This can help in managing breaking changes and ensuring smooth transitions.


[Additional Resource: Atlan - Data Contracts](https://atlan.com/data-contracts/)

## Tool
### Features
- Parses YAML data contracts.
- Generates SVG files with formatted and readable content.
- Supports section headers, bullets, and text wrapping for long descriptions.

### Purpose
Enables data contracts to be defined in .yaml for easy usability within a codebase, while allowing for the conversion of those contracts to more client-friendly for shareout, validation, and buy-in.

### Directory Structure
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

### Build

1. **Clone the repository**:
    ```sh
    git clone https://github.com/et-do/data-contract-generator.git
    cd data-contract-generator
    ```

2. **Open in Devcontainer (if Go is not installed locally)**:
    - Open the project in Visual Studio Code.
    - Use the "Remote - Containers" extension to open the project in a container.

3. **Build the binary**:
    ```sh
    go build -o data-contract-generator main.go
    ```
    
### Use

1. **Source your YAML files**:
   - Place your YAML files in any directory of your choice.

2. **Run the tool**:
    ```sh
    ./data-contract-generator ./directory_with_yaml_files
    ```

   - This will generate SVG files in the same directory as the YAML files.

### Example
#### Input YAML
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

#### SVG Output:
This will generate an SVG file named with the content formatted for readability

![Example Output](/static/contract.png)

### Development Environment

This project includes a development environment configuration for Visual Studio Code with Docker.

1. **Setup Dev Container**:
   - Open the project in Visual Studio Code.
   - Use the "Remote - Containers" extension to open the project in a container.

### Contributing

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

### License

This project is licensed under the MIT License.
