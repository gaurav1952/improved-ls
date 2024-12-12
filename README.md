# Improve-LS

**Improve-LS** is a tool designed to enhance the `ls` command functionality by providing additional features and customization options for listing directory contents.

## Features
- Enhanced directory listings with customizable output.
- Additional filters and sorting options.
- Compatibility with standard Unix-like environments.

## Requirements

- Golang should be installed 
- A Unix-like operating system (Linux, macOS, or similar).


## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/gaurav1952/improve-ls.git
   cd improve-ls
   ```

2. Compile the source code:
   ```bash
   go build  -o list 
   ```

3. Make the binary executable globally (optional):
   ```bash
   sudo mv list /usr/local/bin/
   ```

## Usage

Run the command:
```bash
./list [directory]
```


### Examples
 List all files in the current directory:
```bash
./list # By default it list file in current directory
```


## Contributing

Feel free to fork the repository and submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
