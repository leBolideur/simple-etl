# Simple-ETL

**Simple-ETL** is a lightweight command-line tool written in Go for applying transformations on CSV files. It supports filtering and modifying data.

## Features

- **Filtering rows** based on column values:
  - Numeric filters: `gt`, `gte`, `lt`, `lte`, `eq`, `ne`
  - Boolean filters: `eq`, `ne`
  - String length filters: `len_gt`, `len_gte`, `len_lt`, `len_lte`, `len_eq`
- **Modifying column values**:
  - Convert strings to uppercase or lowercase.

## Usage

```bash
go run main.go -in=<input_file> -filter=<filters> -modifier=<modifiers>
```

### Arguments

- `-in` (required): The path to the input CSV file.
- `-filter`: A comma-separated list of filters to apply.
  - Format: `<column>:<operation>:<value>`
- `-modifier`: A comma-separated list of modifiers to apply.
  - Format: `<column>:<operation>`

### Examples

#### Example 1: Filter rows based on numeric and boolean values

```bash
go run main.go -in=samples/sample.csv -filter=salary:gt:70000,active:eq:true
```

This command filters rows where the `salary` is greater than 70,000 and `active` is `true`.

#### Example 2: Filter rows with string length conditions

```bash
go run main.go -in=samples/sample.csv -filter=department:len_gt:5
```

This command filters rows where the length of the `department` column is greater than 5 characters.

#### Example 3: Apply multiple modifiers

```bash
go run main.go -in=samples/sample.csv -modifier=name:uppercase,department:lowercase
```

This command converts the `name` column to uppercase and the `department` column to lowercase.

#### Example 4: Combine filters and modifiers

```bash
go run main.go -in=samples/sample_large.csv -filter=age:gte:30,salary:lt:100000 -modifier=name:uppercase
```

This command filters rows where `age` is greater than or equal to 30 and `salary` is less than 100,000, and then converts the `name` column to uppercase.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.
