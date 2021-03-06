# proto-gen-dft

PGD is a protoc plugin to generate golang code for proto3 default.
Inspired by [pgv](https://github.com/envoyproxy/protoc-gen-validate).

## TODO:

- [x] support json as default of complex struct
- [] support WKT
- [] rename functions
- [] check default


## Example
```
syntax = "proto3";

option go_package = "github.com/j2gg0s/protoc-gen-dft/examples/basic";
package basic;

import "dft/dft.proto";

message Everything {
  double double_value = 1 [(dft.dft).number = 1.5];
  int64 int64_value = 2 [(dft.dft).number = 3];
  string string_value = 3 [(dft.dft).string = "hello"];

  repeated double double_valus = 4 [(dft.dft).numbers.values = "1.5, 2.5, 3.5"];
  repeated int64 int64_values = 5 [(dft.dft).numbers.values = "1, 2, 3"];
  repeated string string_values = 6 [(dft.dft).strings = {values: "A, B, C"}];

  Foo struct_value = 7 [(dft.dft).json = "{name: \"j2gg0s\", age: 18}"];

  double none_value = 15;
}

message Foo {
  string name = 1;
  int32 age = 2;
}
```

Detail in `examples/basic`.
