# Kubeconfig Tool

A tool for managing and merging multiple kubeconfig files.

## Installation

### Using Make Install (Recommended)

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

### Alternative: Manual Installation

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
