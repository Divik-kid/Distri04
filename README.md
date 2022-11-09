# Distri04
peertopeervirus

Technical Requirements:

-Use Golang to implement the service's nodes
-Provide a README.md, that explains how to start your system
-Use gRPC for message passing between nodes
-Your nodes need to find each other.  For service discovery, you can choose one of the following options
*Supply a file with IP addresses/ports of other nodes
*Enter IP address/ports through the command line
*use a package for service discovery, like the Serf package 
-Demonstrate that the system can be started with at least 3 nodes
-Demonstrate using your system's logs,  a sequence of messages in the system, that leads to a node getting access to the Critical Section. You should provide a discussion of your algorithm, using examples from your logs.