# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.

# BIN="/usr/local/bin" && \
BIN="." && \
VERSION="1.0.0" && \
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"

# https://docs.buf.build/installation
# docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf lint
