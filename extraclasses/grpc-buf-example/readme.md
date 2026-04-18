grpc Modern Workflow with Buf: 
Using buf for linting,detecting breaking changes,and simplified code generation, in go lang with simple example

The traditional gRPC workflow often feels like a collection of fragile shell scripts and manual protoc installs. Buf changes the game by treating Protocol Buffers like a modern software project—complete with dependency management, linting, and a consistent build system.

Here is the modern workflow for Go using Buf.

1. Project Setup
First, initialize your Go module and install the necessary Buf tools.

Bash
mkdir grpc-buf-example && cd grpc-buf-example
go mod init github.com/youruser/grpc-buf-example

# Install Buf CLI
# (On macOS: brew install bufbuild/buf/buf)
The File Structure
A clean Buf project usually separates the .proto files from the generated code:

Plaintext
.
├── buf.yaml          # Project configuration
├── buf.gen.yaml      # Generation strategy
├── proto/            # Your API definitions
│   └── greet/v1/
│       └── greet.proto
└── gen/              # Generated Go code (output)
2. Defining the API
Create proto/greet/v1/greet.proto. Notice we follow the Buf naming convention (package matches directory structure).

Protocol Buffers
syntax = "proto3";

package greet.v1;

option go_package = "github.com/youruser/grpc-buf-example/gen/greet/v1;greetv1";

service GreetService {
  rpc SayHello(SayHelloRequest) returns (SayHelloResponse);
}

message SayHelloRequest {
  string name = 1;
}

message SayHelloResponse {
  string greeting = 1;
}
3. The Buf Configuration
Instead of a 100-line protoc command, we use two simple YAML files.

buf.yaml (Linting & Breaking Changes)
Run buf config init in your root to create this. It defines where your protos live.

YAML
# For details on buf.yaml configuration, visit https://buf.build/docs/configuration/v2/buf-yaml
# version: v2 is the modern standard
version: v2
modules:
  - path: proto # This replaces the 'build: roots' from v1
lint:
  use:
    - STANDARD
breaking:
  use:
    - FILE
buf.gen.yaml (Simplified Generation)
This file tells Buf which plugins to use. No more path-mangling!

YAML
version: v1
plugins:
  - plugin: go
    out: gen
    opt: paths=source_relative
  - plugin: go-grpc
    out: gen
    opt: paths=source_relative
4. The Workflow in Action

Step A: Linting
Keep your API clean and consistent. Buf will catch things like missing comments or non-camelCase fields.

Bash
buf lint
Step B: Detecting Breaking Changes
This is Buf's "superpower." It compares your current local files against a previous git commit or the main branch.

Bash
# Compare local changes against the main branch
buf breaking --against '.git#branch=main'
Why this matters: It prevents you from accidentally deleting a field or changing a field ID that would crash your production mobile apps or microservices.

Step C: Code Generation
Forget complex protoc -I... commands. Just run:

Bash
buf generate
Buf automatically finds all protos in your configured directories and outputs them to the gen/ folder as specified.

5. Using the Generated Code in Go
Now you can simply import your generated package like any other Go library.

Go
package main

import (
    "fmt"
    "github.com/youruser/grpc-buf-example/gen/greet/v1" // The generated code
)

func main() {
    resp := &greetv1.SayHelloResponse{
        Greeting: "Hello from the modern stack!",
    }
    fmt.Println(resp.Greeting)
}
Key Advantages of this Workflow:
No local dependencies: You don't need protoc installed on your machine; Buf can handle the heavy lifting.

Atomic Generation: It generates everything in one pass.

Remote Plugins: You can use "Remote Plugins" so you don't even need the Go gRPC plugins installed locally—Buf pulls them from the cloud.

>buf lint

proto\greet\v1\greet.proto:3:1:Files with package "greet.v1" must be within a directory "greet\v1" relative to root but were in directory "proto\greet\v1".

That error is a classic "Buf rite of passage." It happens because Buf expects your directory structure to match your Protobuf package name exactly, starting from where your buf.yaml is located.

Since your package is greet.v1, Buf expects the file to be at greet/v1/greet.proto, but you have it at proto/greet/v1/greet.proto.

The Fix: Point Buf to the "proto" folder
You don't need to move your files. You just need to tell Buf that the proto directory is the root of your Protobuf source tree.

Update your buf.yaml file to include the roots configuration:

YAML
version: v1
# Add this 'build' section:
build:
  roots:
    - proto
lint:
  use:
    - DEFAULT
break:
  use:
    - FILE
Why this works
By default, Buf looks at the directory containing buf.yaml as the starting point (the "root").

Without the fix: Buf sees proto/greet/v1/greet.proto and thinks the package should be proto.greet.v1.

With the fix: Buf enters the proto folder first. Now it sees greet/v1/greet.proto, which perfectly matches your package greet.v1; declaration.

Pro-Tip: Update your Generation Config
Since you changed the root to proto, make sure your buf.gen.yaml is also aware of where to look. Usually, if you run buf generate from the project root, it will now correctly map the paths.

If you ever run into issues with the generated code location, you can run the command explicitly pointing to the template:

Bash
buf generate proto
Try running buf lint again—it should come back clean now!

# For details on buf.yaml configuration, visit https://buf.build/docs/configuration/v2/buf-yaml

version: v2

build:

  roots:

    - proto

lint:

  use:

    - STANDARD

breaking:

  use:

    - FILE

>buf lint

Failure: decode buf.yaml: yaml: unmarshal errors:

  line 3: field build not found

Ah, you’ve run into a version mismatch! You are using Buf Configuration v2, but the build: roots: syntax you used belongs to v1.

In v2, Buf simplified the configuration. You no longer use a build block with roots. Instead, you define Modules directly.

The Fix: Update your buf.yaml to v2 Syntax
Change your buf.yaml to look like this:

YAML
# version: v2 is the modern standard
version: v2
modules:
  - path: proto # This replaces the 'build: roots' from v1
lint:
  use:
    - STANDARD
breaking:
  use:
    - FILE
Why the error happened
In v1: You defined a single "root" directory where protos lived.

In v2: Buf moved to a "workspace" model. Everything under the modules list is treated as a self-contained unit of Protobuf files. By setting path: proto, you are telling Buf: "Treat the proto folder as the base for all package name calculations."

One more thing: Check your buf.gen.yaml
If you are using v2, your generation file also looks slightly different (the plugins are now nested under managed or defined more cleanly). For a simple setup, ensure your buf.gen.yaml matches the v2 style as well:

YAML
version: v2
managed:
  enabled: true
  override:
    - file_option: go_package_prefix
      value: github.com/youruser/grpc-buf-example/gen
plugins:
  - local: protoc-gen-go
    out: gen
    opt: paths=source_relative
  - local: protoc-gen-go-grpc
    out: gen
    opt: paths=source_relative
Try running buf lint now. It should successfully see that proto is your module root and validate that greet/v1/greet.proto matches package greet.v1.

give me server and client code

Now that you have your boilerplate and Buf configuration sorted, let’s implement a simple Greeter server and client in Go.

Before running these, ensure you have run buf generate so the gen/ directory contains your code.

1. The Server (server/main.go)
The server implements the interface generated by Buf.

Go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	// Replace with your actual module path from go.mod
	pb "github.com/youruser/grpc-buf-example/gen/greet/v1"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SayHelloResponse{Greeting: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
2. The Client (client/main.go)
The client dials the server and sends a single request.

Go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// Replace with your actual module path from go.mod
	pb "github.com/youruser/grpc-buf-example/gen/greet/v1"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.SayHelloRequest{Name: "Buf User"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())
}
3. Running the Example
Sync Dependencies:
Run go mod tidy to pull in google.golang.org/grpc and the protobuf dependencies.

Start the Server:

Bash
go run server/main.go
Run the Client (in a new terminal):

Bash
go run client/main.go
A Quick Tip on "Managed Mode"
In your buf.gen.yaml (v2), we used managed: enabled: true. This is a lifesaver because it allows you to keep your .proto files clean of Go-specific options like option go_package. Buf will inject those values during generation based on your config, making your Protobuf files "pure" and easier to share across different languages (like Python or Java) without Go-specific clutter.

Does the server start up correctly, or are you seeing any "package not found" errors in the imports?