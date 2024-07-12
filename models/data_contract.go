package models

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type ContactPerson struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
	Phone string `yaml:"phone"`
}

type DataContract struct {
	Name                  string         `yaml:"name"`
	Date                  string         `yaml:"date"`
	DataSource            string         `yaml:"data_source"`
	TableName             string         `yaml:"table_name"`
	ExpectedSchema        []Field        `yaml:"expected_schema"`
	Frequency             string         `yaml:"frequency"`
	PIISensitive          bool           `yaml:"PII_sensitive"`
	Description           string         `yaml:"description"`
	ContactPerson         ContactPerson  `yaml:"contact_person"`
	DataRetentionPolicy   string         `yaml:"data_retention_policy"`
	SecurityClassification string        `yaml:"security_classification"`
	Version               string         `yaml:"version"`
}
