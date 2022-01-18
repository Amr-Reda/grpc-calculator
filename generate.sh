#!/bin/bash

protoc proto/calculator.proto --go_out=plugins=grpc:.
