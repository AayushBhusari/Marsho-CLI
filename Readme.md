# Marsho CLI File Manager

**Marsho** is a fast and reliable Command Line Interface (CLI) file manager built with Go. It leverages the Cobra and PromptUI libraries to provide an interactive and efficient file browsing experience. Marsho offers a seamless way to navigate directories, view file details, and open files directly from the terminal.

## Features

- **Directory Navigation**: Browse through directories and subdirectories easily using an intuitive prompt-based interface.
- **File Management**: View file details such as size and modification time. Open files directly from the CLI using the OS's default programs.
- **Exit Option**: A built-in "Exit" option allows you to exit the file manager from any directory level.
- **Interactive Interface**: Clean and interactive file/directory selection with highlighted active selections and clear navigation options.

## Future Improvements

- **File Operations**: Implement additional file operations such as copy, move, and delete.
- **File Search**: Add functionality for searching files within directories.
- **Enhanced Filtering**: Provide advanced filtering options for handling large directories more efficiently.

## Installation

1. **Download**: Get the `.zip` file for Windows from the [releases page](https://github.com/AayushBhusari/Marsho/releases).
2. **Extract**: Right-click the downloaded `.zip` file and select "Extract All..." to unzip the file to a location of your choice.
3. **Setup**: Move the extracted `.exe` file to a directory of your choice. To make it accessible from anywhere in the terminal, add this directory to your system's `PATH` environment variable:
   - Open the Start Menu and search for "Environment Variables."
   - Select "Edit the system environment variables."
   - In the System Properties window, click on "Environment Variables..."
   - In the Environment Variables window, locate the `Path` variable under "System variables," select it, and click "Edit..."
   - Click "New" and add the path to the directory where you placed the `.exe` file.
   - Click "OK" to close all dialog boxes.
4. **Verify Installation**: Open Command Prompt and type `marsho` to ensure the file manager runs correctly.

## Usage

To start using Marsho, open your terminal and simply run:

```bash
marsho
```

The CLI will display a list of files and directories in your current directory. You can navigate through directories, view file details, and open files directly from the prompt. Use the "Exit" option to close the file manager when you're done.

## Examples

- **Navigate to a Parent Directory**: Select the `..` option to move up one level in the directory hierarchy.
- **Open a File**: Select a file to view its details and open it with the default application for that file type.
- **Exit the File Manager**: Choose the "Exit" option to quit the file manager.

## Contributing

Contributions to Marsho are welcome! If you have suggestions for new features or improvements, please submit an issue or pull request on [GitHub](https://github.com/AayushBhusari/Marsho/tree/v1.0.0).

```

```
