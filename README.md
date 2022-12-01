# canparse
CAN bus traffic parsing using CAN database files.

This library was partially inspired by and borrows some naming conventions from [cantools/cantools](https://github.com/cantools/cantools).

## Supported CAN Database Formats
The plan is to initially support the open-source [KCD format](https://github.com/julietkilo/kcd), and hopefully
expand the capabilities of the library to support DBC and SYM formats in the future. For now a tool such
as [canmatrix](https://github.com/ebroecker/canmatrix) can be used to convert other file types to KCD format.
