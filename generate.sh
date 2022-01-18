#!/bin/bash

protoc calculator_proto/calculator.proto --go_out=plugins=grpc:.
