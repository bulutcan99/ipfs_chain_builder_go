# ipfs_chain_builder_go
This project revolves around querying a MySQL table to extract column names, their respective types, and cell values, structuring them into JSON format. Additionally, it integrates with IPFS for data storage and CID generation. The implementation involves constructing a chain of hashes for each row's cells, ultimately linked to an IPNS hash.

Requirements in the Makefile:
make start -> You can start the project with dependencies injected.

start: Initiates the project with necessary steps.
go.build: Get a build file for go.
go.mod: Installs dependencies essential for running the project.
Env File:
The env file serves the purpose of storing environment variables or sensitive information, ensuring secure configurations, such as database credentials.

Note: These instructions aim to streamline the setup and operation of the project, enabling users to seamlessly interact with MySQL databases, generate JSON data, utilize IPFS for storage, and manage environmental configurations.
