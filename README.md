# Kubeconfig Tool

A tool for managing and merging multiple kubeconfig files.

## Installation

### Option 1: Download binary

Download the pre-compiled binary from the [releases page](https://github.com/moonlight8978/kubeconfig/releases/latest) and add it to your PATH.

#### Linux and macOS

```bash
# Replace X.Y.Z with the latest version
curl -L https://github.com/moonlight8978/kubeconfig/releases/download/vX.Y.Z/kubeconfig_X.Y.Z_Linux_x86_64.tar.gz | tar xz
sudo mv kubeconfig /usr/local/bin/
```

#### Windows (using PowerShell)

```powershell
# Replace X.Y.Z with the latest version
Invoke-WebRequest -Uri https://github.com/moonlight8978/kubeconfig/releases/download/vX.Y.Z/kubeconfig_X.Y.Z_Windows_x86_64.zip -OutFile kubeconfig.zip
Expand-Archive -Path kubeconfig.zip -DestinationPath .
mv .\kubeconfig.exe $Env:USERPROFILE\bin\
# Make sure $Env:USERPROFILE\bin is in your PATH
```

### Option 2: Using Go Install

If you have Go installed, you can use the `go install` command:

```bash
# Install the latest version
go install github.com/moonlight8978/kubeconfig@latest
```

Or to install a specific version:

```bash
# Install a specific version
go install github.com/moonlight8978/kubeconfig@vX.Y.Z
```

**Note:** Make sure your `$GOPATH/bin` is in your PATH.

### Option 3: Using Make Install (Recommended)

```bash
# Clone the repository
git clone https://github.com/moonlight8978/kubeconfig.git
cd kubeconfig

# Build and install to $GOPATH/bin
make install
```

Make sure your `$GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Option 4: Manual Installation

```bash
# Clone the repository
git clone https://github.com/moonlight8978/kubeconfig.git
cd kubeconfig

# Build only
go build -o kubeconfig bin/main.go

# Move to a directory in your PATH
sudo mv kubeconfig /usr/local/bin/
```

## Troubleshooting Installation

If you don't see the binary in your `~/go/bin` directory:

1. Check if GOPATH is set correctly:

   ```bash
   echo $GOPATH
   ```

   If not set, it defaults to `~/go`. You can set it with:

   ```bash
   export GOPATH=~/go
   ```

2. Make sure the bin directory exists:

   ```bash
   mkdir -p ~/go/bin
   ```

3. Make sure the bin directory is in your PATH:

   ```bash
   export PATH=$PATH:~/go/bin
   ```

   For permanent addition, add this line to your `.bashrc` or `.zshrc`.

## Usage

### Merge Kubeconfigs

```bash
kubeconfig merge --configs ~/.kube/.mergecfg.yaml --output ~/.kube/config
```

Default locations:

- Config list: `~/.kube/.mergecfg.yaml`
- Output: `~/.kube/config`

### Options

```
kubeconfig merge --help
```
