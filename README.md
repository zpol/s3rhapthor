# S3Rapthor - AWS S3 Bucket Analyzer

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

> **S3Rapthor** is an AWS S3 bucket analysis tool that allows you to explore, analyze, and download content from public S3 buckets efficiently.

## Features

- **Bucket Analysis**: Explore public S3 buckets and extract detailed information
- **File Statistics**: Generate comprehensive statistics about file types and extensions
- **Bulk Download**: Automatic download of all bucket files
- **Data Organization**: Save all data in an organized structure
- **Colored Output**: Color-coded output for better readability
- **Fast and Efficient**: Written in Go for maximum performance

## Installation

### Prerequisites

- Go 1.18 or higher
- Git

### Installation from Repository

```bash
# Clone the repository
git clone https://github.com/zpol/s3rhapthor.git
cd s3rhapthor

# Install dependencies
go mod tidy

# Build the tool
go build -o s3Rapthor main.go
```

### Direct Installation with Go

```bash
go install github.com/zpol/s3rhapthor@latest
```

## Usage

### Basic Usage

```bash
./s3Rapthor <S3_BUCKET_URL>
```

### Examples

```bash
# Analyze a public S3 bucket
./s3Rapthor https://my-bucket.s3.amazonaws.com/

# Analyze bucket with specific content
./s3Rapthor https://public-data.s3.us-east-1.amazonaws.com/
```

### Parameters

- `<S3_BUCKET_URL>`: Complete URL of the S3 bucket to analyze (required)

## Output Structure

The tool automatically creates a `data/` directory with the following files:

```
data/
├── bucket-name-s3_bucket.txt    # Basic bucket information
├── bucket-name-s3_bucket.xml    # Raw XML bucket data
├── bucket-name-s3_bucket.brf    # List of file URLs
└── downloaded_files/            # Downloaded files (if enabled)
```

## Functionality

### 1. Bucket Analysis
- Verifies bucket accessibility
- Extracts bucket metadata
- Lists all contained objects

### 2. File Statistics
- Total file count
- Analysis by file extensions
- Percentages of each file type
- Highlights potentially interesting files (PDF, DOC, etc.)

### 3. File Search
```bash
# Search files by extension
grep '\.pdf' data/bucket-name-s3_bucket.brf
grep '\.doc' data/bucket-name-s3_bucket.brf
grep '\.zip' data/bucket-name-s3_bucket.brf
```

### 4. File Download
The download functionality is commented out by default. To enable it, uncomment the line:
```go
download_all_files(brief_file)
```

## Example Output

```
S3Rapthor v0.0.9

·	An AWS S3 bucket analyzer

>> [ 200 ] : https://my-bucket.s3.amazonaws.com/
>> Bucket Name: my-bucket
>> Bucket URL: https://my-bucket.s3.amazonaws.com/
>> Total files: 1250
>> Extension filetypes: 15
>> Extension filetypes expanded:
	.pdf: 45 (3.60%)
	.jpg: 234 (18.72%)
	.png: 156 (12.48%)
	.zip: 23 (1.84%)
	.doc: 12 (0.96%)
	...

>> Type: grep '\.ext' data/my-bucket-s3_bucket.brf to search for an specific file extension
>> Example: grep '\.pdf' data/my-bucket-s3_bucket.brf
```

## Configuration

### Environment Variables (Optional)

```bash
export AWS_REGION=us-east-1
export AWS_PROFILE=default
```

## Security Considerations

- **Public Buckets Only**: This tool is designed to analyze public S3 buckets
- **Data Respect**: Use this tool ethically and responsibly
- **Compliance**: Ensure compliance with data usage policies
- **Permissions**: Verify you have authorization to access the buckets

## Contributing

Contributions are welcome. Please:

1. Fork the project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Disclaimer

This tool is designed solely for educational and research purposes. Users are responsible for complying with all applicable laws and regulations when using this tool.

## Support

If you have issues or suggestions:

- Open an issue on GitHub
- Report bugs with detailed problem information
- Suggest new features

---

**Developed with Go**
