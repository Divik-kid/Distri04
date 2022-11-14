# Mandatory Handin 4 - Distributed Mutual Exclusion
Peer-to-peer Mutual Exclusion

## How to start the system
Run the program in multiple command prompts using `go run main.go 0` for the first terminal, incrementing the argument for every subsequent terminal.
Supports a maximum of 3 users at a time.

## How the program works
After each of the command prompts have been launched press enter to send requests to the other two command prompts. the command prompt will write which peers approved/disaproved the request, the mutual agreement is decided using an implementation of the Ricart & Agrawala algorithm

the Critical state is represented using a print statement to signify that a peer did access it

## System Requirements

* R1: Implement a system with a set of peer nodes, and a Critical Section, that represents a sensitive system operation. Any node can at any time decide it wants access to the Critical Section. Critical section in this exercise is emulated, for example by a print statement, or writing to a shared file on the network.
* R2: Safety: Only one node at the same time is allowed to enter the Critical Section 
* R3: Liveliness: Every node that requests access to the Critical Section, will get access to the Critical Section (at some point in time)

## Technical Requirements

1. Use Golang to implement the service's nodes
1. Provide a README.md, that explains how to start your system
1. Use gRPC for message passing between nodes
1. Your nodes need to find each other. For service discovery, you can choose one of the following options
    1. Supply a file with IP addresses/ports of other nodes
    2. Enter IP address/ports through the command line
    3. use a package for service discovery, like the Serf package 
1. Demonstrate that the system can be started with at least 3 nodes
1. Demonstrate using your system's logs, a sequence of messages in the system, that leads to a node getting access to the Critical Section. You should provide a discussion of your algorithm, using examples from your logs.