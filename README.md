# canparse
CAN bus traffic parsing using CAN database files.

This code borrows heavily from [cantools/cantools](https://github.com/cantools/cantools).

## Supported CAN Database Formats
The plan is to initially support the open-source [KCD format](github.com/julietkilo/kcd), and hopefully
expand the capabilities of the library to support DBC and SYM formats in the future. For now a tool such
as [canmatrix](github.com/ebroecker/canmatrix) can be used to convert other file types to KCD format. 

## Supported CAN Attributes

For now the library will only have a concept of messages, id's, and signals. Multiple busses and nodes
are not supported. It is assumed that every CAN message in the KCD file is being recieved on the bus.
