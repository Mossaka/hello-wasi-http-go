# Hello WASI HTTP from Golang!

> This is a folked version of the [hello-wasi-http](https://github.com/sunfishcode/hello-wasi-http) repository. It has been updated to create a component out of Go code instead of Rust code. The modified README.md is below.

This is a simple tutorial to get started with WASI HTTP (`v0.2.0-rc-2023-12-05`) using the
`wasmtime serve` command in [Wasmtime] 16.0.0, `spin` 2.1.0 (`spin` uses `v0.2.0-rc-2023-11-10`) and `jco serve` command in [jco] 0.14.2. It runs an HTTP server and
forwards requests to a Wasm component via the [WASI HTTP] API.

[Wasmtime]: https://wasmtime.dev
[WASI HTTP]: https://github.com/WebAssembly/wasi-http/
[jco]: https://github.com/bytecodealliance/jco

The WASI HTTP API is settling down but as of this writing not quite stable.
This tutorial uses a snapshot of it that's implemented in Wasmtime 16.0.0.

With that said...

## Let's Go!

First, [install `tinygo`](https://github.com/tinygo-org/tinygo/releases),
version 0.30.0, which is a LLVM-based Go compiler alternative. (See [here] for information about building Wasm components from other
languages too!)

[here]: https://component-model.bytecodealliance.org/language-support.html

<!-- Then, [install `wit-bindgen-cli@0.16.0`](https://github.com/bytecodealliance/wit-bindgen) with `cargo install wit-bindgen-cli@0.16.0`, which is a tool for generating Go bindings for WIT interfaces. -->
Then install `wit-bindgen-cli` with `cargo install wit-bindgen-cli --git https://github.com/bytecodealliance/wit-bindgen --rev 7c9c4626945699efb0379053134f4992d3f36216`, which is a tool for generating Go bindings for WIT interfaces.

Lastly, [install `wasm-tools`](https://github.com/bytecodealliance/wasm-tools/releases/) version 1.0.55, which is a tool for building Wasm components.

With that, build the Wasm component from the source in this repository:

```sh
$ go generate
Generating "target_world/2023_11_10/target-world.go"
Generating "target_world/2023_11_10/target-world_types.go"
Generating "target_world/2023_11_10/target_world.c"
Generating "target_world/2023_11_10/target_world.h"
Generating "target_world/2023_12_05/target-world.go"
Generating "target_world/2023_12_05/target-world_types.go"
Generating "target_world/2023_12_05/target_world.c"
Generating "target_world/2023_12_05/target_world.h"
$ tinygo build -o main_2023_12_05.wasm -target=wasi main_2023_12_05.go
```

This builds a Wasm module, `main.wasm`.

Next, we'll need to create a Wasm component.

```sh
$ wasm-tools component embed wit main_2023_12_05.wasm > main_2023_12_05.embed.wasm
$ wasm-tools component new main_2023_12_05.embed.wasm -o main_2023_12_05.component.wasm --adapt wasi_snapshot_preview1.reactor.2023_12_05.wasm
```

This creates a Wasm component, `main_2023_12_05.component.wasm`.

To run it, we'll need Wasmtime `v16.0.0`. Installation instructions are
on [wasmtime](https://github.com/bytecodealliance/wasmtime/releases/tag/v16.0.0) repo.

Then, in a new terminal, we can run `wasmtime serve` on our Wasm component:

```
$ wasmtime serve -Scommon main_2023_12_05.component.wasm
```

This starts up an HTTP server on `0.0.0.0:8080` (the specific address and port
can be configured with the `--addr=` flag).

With that running, in another window, we can now make requests!

```
$ curl http://localhost:8080
Hello world from Go!!!
```

## Notes

`wasmtime serve` uses the [proxy] world, which is a specialized world just for
accepting requests and producing responses. One interesting thing about the proxy
world is that it doesn't have a filesystem or network API. If you add code to the
example that tries to access files or network sockets, it won't be able to build,
because those APIs are not available in this world. This allows proxy components
to run in many different places, including specialized serverless environments
which may not provide traditional filesystem and network access.

But what if you do want to have it serve some files? One option will be to use
[WASI-Virt](https://github.com/bytecodealliance/WASI-Virt), which is a tool
that can bundle a filesystem with a component.

Another option is to use a custom world. The proxy world is meant to be able
to run in many different environments, but if you know your environment and
you know it has a filesystem, you could create your own world, by including
both the "wasi:http/proxy" and "wasi:filesystem/types" or any other APIs you want
the Wasm to be able to access. This would require a custom embedding of Wasmtime,
as it wouldn't run under plain `wasmtime serve`, so it's a little more work to
set up.

In the future, we expect to see standard worlds emerge that combine WASI HTTP
with many other APIs, such as [wasi-cloud-core].

[wasi-cloud-core]: https://github.com/WebAssembly/wasi-cloud-core

If you're interested in tutorials for any of these options, please reach out
and say hi!

[proxy]: https://github.com/WebAssembly/wasi-http/blob/main/wit/proxy.wit

## Running in Spin 2.1

To run this component in Spin 2.1, you'll need to first download the Spin 2.1 runtime from [here](https://github.com/fermyon/spin/releases/tag/v2.1.0)

Then, you'll need to create a `spin.toml` file in the same directory as the `main.component.wasm` file. The `spin.toml` file should look like this:

```toml
spin_manifest_version = 2

[application]
name = "hello-wasi-http"
version = "1.0.0"

[[trigger.http]]
route = "/"
component = "hello"

[component.hello]
source = "main_2023_11_10.component.wasm"
[component.hello.build]
command = """go generate && 
    tinygo build -o main_2023_11_10.wasm -target=wasi main_2023_11_10.go && 
    wasm-tools component embed wit/2023_11_10 main_2023_11_10.wasm > main_2023_11_10.embed.wasm && 
    wasm-tools component new main_2023_11_10.embed.wasm -o main_2023_11_10.component.wasm --adapt wasi_snapshot_preview1.reactor.2023_11_10.wasm
"""

[component.hello2]
source = "main_2023_12_05.component.wasm"
[component.hello2.build]
command = """go generate && 
    tinygo build -o main_2023_12_05.wasm -target=wasi main_2023_12_05.go && 
    wasm-tools component embed wit/2023_12_05 main_2023_12_05.wasm > main_2023_12_05.embed.wasm && 
    wasm-tools component new main_2023_12_05.embed.wasm -o main_2023_12_05.component.wasm --adapt wasi_snapshot_preview1.reactor.2023_12_05.wasm
"""
```

This repo has a `spin.toml` file already.

> Note: The latest `spin` version only supports `wasi-http` version `2023_11_10`, so the only component used in spin HTTP trigger is the 2023_11_10 version of the component.

Then, you can run the component with the following command:

```sh
$ spin up --build
Building component hello with `go generate && 
    tinygo build -o main_2023_11_10.wasm -target=wasi main_2023_11_10.go && 
    wasm-tools component embed wit/2023_11_10 main_2023_11_10.wasm > main_2023_11_10.embed.wasm && 
    wasm-tools component new main_2023_11_10.embed.wasm -o main_2023_11_10.component.wasm --adapt wasi_snapshot_preview1.reactor.2023_11_10.wasm
`
Generating "target_world/2023_11_10/target-world.go"
Generating "target_world/2023_11_10/target-world_types.go"
Generating "target_world/2023_11_10/target_world.c"
Generating "target_world/2023_11_10/target_world.h"
Finished building all Spin components
Logging component stdio to ".spin/logs/"

Serving http://127.0.0.1:3000
Available Routes:
  hello: http://127.0.0.1:3000
```

With that running, in another window, we can now make requests!

```
$ curl http://127.0.0.1:3000
Hello world from Go!!!
```

## Running in jco 0.14.2

To run this component in jco 0.14.2, you'll need to first download the jco 0.14.2 runtime 
```
npm install @bytecodealliance/jco@0.14.2 -g
```

and then run the following command:

```sh
$ jco serve main_2023_12_05.component.wasm
```

With that running, in another window, we can now make requests!

```
$ curl http://localhost:8000
Hello world from Go!!!
```

## Creating this repo

TODO: Add instructions for creating this repo from scratch.
