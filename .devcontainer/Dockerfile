# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.234.0/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Debian OS version (use bullseye on local arm64/Apple Silicon): buster, bullseye
ARG VARIANT="bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/rust:${VARIANT} AS build

RUN rustup default stable
RUN rustup target add wasm32-wasi

# TinyGo installation, see https://tinygo.org/getting-started/install/linux/ for instructions
ARG TINYGO_URL="https://github.com/tinygo-org/tinygo/releases/download/v0.30.0/tinygo0.30.0.linux-arm64.tar.gz"
RUN curl -sL "$TINYGO_URL" -o tinygo_arm64.tar.gz && \
    tar -xzf tinygo_arm64.tar.gz && \
    mv tinygo /usr/local/tinygo

## wit-bindgen@7c9c46269
# RUN curl -LO https://github.com/bytecodealliance/wit-bindgen/releases/download/wit-bindgen-cli-0.16.0/wit-bindgen-v0.16.0-aarch64-linux.tar.gz && \
#     tar -xzf wit-bindgen-v0.16.0-aarch64-linux.tar.gz && \
#     mv wit-bindgen-v0.16.0-aarch64-linux/wit-bindgen /usr/local/bin
RUN cargo install wit-bindgen-cli --git https://github.com/bytecodealliance/wit-bindgen --rev 7c9c4626945699efb0379053134f4992d3f36216

## wasm-tools@1.0.55
RUN curl -LO https://github.com/bytecodealliance/wasm-tools/releases/download/wasm-tools-1.0.55/wasm-tools-1.0.55-aarch64-linux.tar.gz && \
    tar -xzf wasm-tools-1.0.55-aarch64-linux.tar.gz && \
    mv wasm-tools-1.0.55-aarch64-linux/wasm-tools /usr/local/bin

## spin@2.2.0
RUN curl -LO https://github.com/fermyon/spin/releases/download/v2.2.0/spin-v2.2.0-linux-aarch64.tar.gz && \
    tar -xzf spin-v2.2.0-linux-aarch64.tar.gz && \
    mv spin /usr/local/bin

## go
RUN curl -LO https://go.dev/dl/go1.21.4.linux-arm64.tar.gz && \
    tar -C /usr/local -xzf go1.21.4.linux-arm64.tar.gz

## wasmstime@17.0.0
RUN curl -LO https://github.com/bytecodealliance/wasmtime/releases/download/v17.0.0/wasmtime-v17.0.0-aarch64-linux.tar.xz && \
    tar xf wasmtime-v17.0.0-aarch64-linux.tar.xz && \
    mv wasmtime-v17.0.0-aarch64-linux/wasmtime /usr/local/bin

## install npm
RUN curl -fsSL https://deb.nodesource.com/setup_21.x | sudo -E bash - &&\
    sudo apt-get install -y nodejs

## jco@1.0.0
RUN npm install @bytecodealliance/jco@1.0.0 -g 

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/usr/local/go"
RUN go install -v golang.org/x/tools/gopls@latest

FROM mcr.microsoft.com/vscode/devcontainers/rust:${VARIANT}
COPY --from=build /usr/local/bin /usr/local/bin
COPY --from=build /usr/local/go /usr/local/go
COPY --from=build /usr/local/cargo/bin /usr/local/cargo/bin
COPY --from=build /usr/local/tinygo /usr/local/tinygo
COPY --from=build /usr/lib/node_modules /usr/lib/node_modules
COPY --from=build /usr/bin /usr/bin
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/usr/local/tinygo/bin:${PATH}"
ENV GOROOT="/usr/local/go"
ENV TINYGOROOT="/usr/local/tinygo"