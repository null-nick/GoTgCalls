# README - Native Libraries Download for GoTgCalls

This folder (`/GoTgCalls/bindings/lib`) is dedicated to loading the native libraries required by **GoTgCalls**. Users must download the appropriate library manually based on their operating system.

## System Requirements

To ensure **GoTgCalls** functions properly, you need to download the specific **shared library package** for your platform, as it contains all the required dependencies within the library itself.

Supported library types: **`-shared_libs.zip`**

### Platforms Currently Unsupported
The following platforms are not currently supported by **GoTgCalls**:
- **Linux ARM**
- **Windows ARM**
- **macOS (Intel and ARM)**

Support for these platforms may be added in the future. Please check the official repository or documentation for updates.

## Download Instructions

1. Visit the [latest release page on GitHub](https://github.com/pytgcalls/ntgcalls/releases/latest).
2. Download the appropriate **`-shared_libs.zip`** file for your operating system:
    - **Linux x86_64**: `ntgcalls.linux-x86_64-shared_libs.zip`
    - **Windows x86_64**: `ntgcalls.windows-x86_64-shared_libs.zip`
    - **macOS ARM64** (currently unsupported): `ntgcalls.macos-arm64-shared_libs.zip`

> **Important**: Ensure that you download the `-shared_libs.zip` files, as these contain all necessary dependencies. Avoid downloading the `-static_libs.zip` versions, which are not self-contained.

## Setup Instructions

### 1. Extract the Files
After downloading the appropriate `.zip` file:
- Extract its contents. You should find a file with one of the following extensions:
    - `.so` (Linux)
    - `.dll` (Windows)
    - `.dylib` (macOS)

### 2. Move the File to the `/GoTgCalls/bindings/lib` Folder
Manually move the extracted library file into the `lib` folder:
```
/GoTgCalls/
├── bindings/
│   ├── lib/
│   │   ├── libntgcalls.so  # Linux example
```

### 3. Set Permissions (Linux)
On Linux, ensure the file has executable permissions using the `chmod` command:
```bash
chmod +x libntgcalls.so
```

### 4. Verify the Setup
Once the file has been placed in the correct folder:
- Run the **GoTgCalls** client.
- Check that there are no errors related to loading the native library.

## Troubleshooting

- **Error: "Library not found"**:  
  Make sure you:
    - Downloaded the correct file for your operating system.
    - Placed the file in the `/GoTgCalls/bindings/lib` folder.

- **Error: Permission denied** (Linux):  
  Ensure the file has executable permissions using the `chmod` command:
  ```bash
  chmod +x libntgcalls.so
  ```

- **Compatibility Issues**:  
  Confirm that the downloaded library version matches the requirements of **GoTgCalls** (check the official documentation or changelog).

---

## `setup_libraries.sh` Script

The `/GoTgCalls/setup_libraries.sh` script automates the process of downloading and setting up the native libraries required for **GoTgCalls**, simplifying the steps outlined above. Here’s an overview of its functionality:

### Description

The `setup_libraries.sh` script automatically performs the following actions:

1. **Detects the Operating System**: The script automatically detects the operating system being used (Linux, Windows, macOS).
2. **Downloads the Correct Package**: Based on the detected OS, the script downloads the appropriate `-shared_libs.zip` package from the GitHub repository.
3. **Extracts the Files**: Once downloaded, the script extracts the contents of the `.zip` file.
4. **Moves the Files**: The script moves the extracted library file into the `/GoTgCalls/bindings/lib` folder.
5. **Sets Permissions** (Linux): If the OS is Linux, the script automatically sets the executable permissions for the native library.

### Usage

To run the script, simply execute the following command in the terminal (from the root of the project):

```bash
cd GoTgCalls/ && ./setup_libraries.sh
```

> **Note**: Ensure that the script has execute permissions. If it does not, run:
```bash
chmod +x setup_libraries.sh
```

The script will automatically handle downloading, extracting, and setting up the libraries.