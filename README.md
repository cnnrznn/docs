# docs

## Project Goals
The goal of this project is to re-implement the operational transform, server-client algorithm (see github.com/cnnrznn/OperationalTransform).
The purpose being for practice using Golang and gRPC. Both the Client and the Server "engines" will be written in Go and they will talk through gRPC pipes.

## The Algorithm
Operational transform is an algorithm for performing realtime collaborative editing, typically of a shared document.
Operations that are performed are "transformed" such that they result in the same view of the document on all clients once all updates are received.

### A Simple Example: Insert
