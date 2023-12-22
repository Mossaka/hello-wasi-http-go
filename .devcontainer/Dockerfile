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

## wit-bindgen@0.16.0
RUN curl -LO https://github.com/bytecodealliance/wit-bindgen/releases/download/wit-bindgen-cli-0.16.0/wit-bindgen-v0.16.0-aarch64-linux.tar.gz && \
    tar -xzf wit-bindgen-v0.16.0-aarch64-linux.tar.gz && \
    mv wit-bindgen-v0.16.0-aarch64-linux/wit-bindgen /usr/local/bin

## wasm-tools@1.0.48
RUN curl -LO https://github.com/bytecodealliance/wasm-tools/releases/download/wasm-tools-1.0.48/wasm-tools-1.0.48-aarch64-linux.tar.gz && \
    tar -xzf wasm-tools-1.0.48-aarch64-linux.tar.gz && \
    mv wasm-tools-1.0.48-aarch64-linux/wasm-tools /usr/local/bin

## spin@2.1.0
RUN curl -LO https://github.com/fermyon/spin/releases/download/v2.1.0/spin-v2.1.0-linux-aarch64.tar.gz && \
    tar -xzf spin-v2.1.0-linux-aarch64.tar.gz && \
    mv spin /usr/local/bin

## go
RUN curl -LO https://go.dev/dl/go1.21.4.linux-arm64.tar.gz && \
    tar -C /usr/local -xzf go1.21.4.linux-arm64.tar.gz

## wasmstime@15.0.1
RUN curl -LO https://github.com/bytecodealliance/wasmtime/releases/download/v15.0.1/wasmtime-v15.0.1-aarch64-linux.tar.xz && \
    tar xf wasmtime-v15.0.1-aarch64-linux.tar.xz && \
    mv wasmtime-v15.0.1-aarch64-linux/wasmtime /usr/local/bin

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/usr/local/go"
RUN go install -v golang.org/x/tools/gopls@latest

FROM mcr.microsoft.com/vscode/devcontainers/rust:${VARIANT}
COPY --from=build /usr/local/bin /usr/local/bin
COPY --from=build /usr/local/go /usr/local/go
COPY --from=build /usr/local/cargo/bin /usr/local/cargo/bin
COPY --from=build /usr/local/tinygo /usr/local/tinygo
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/usr/local/tinygo/bin:${PATH}"
ENV GOROOT="/usr/local/go"
ENV TINYGOROOT="/usr/local/tinygo"