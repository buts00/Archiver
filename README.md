# Shannon-Fano Archiver

This program provides functionality for archiving and extracting files using the Shannon-Fano algorithm.

## Usage

To use this program, follow these steps:

1. Clone the repository.
2. Run the following commands in your terminal:

### Packing a File

To pack a file using the Shannon-Fano algorithm, execute the following command:

```bash
make pack_sf INPUT_FILE=<path_to_input_file>
```

Replace `<path_to_input_file>` with the path to the file you want to archive.

### Unpacking a File

To unpack a file that has been archived using the Shannon-Fano algorithm, execute the following command:

```bash
make unpack_sf INPUT_FILE=<path_to_archived_file>
```

Replace `<path_to_archived_file>` with the path to the archived file you want to extract.

## Example

Suppose you have a file named `example.txt` that you want to pack using the Shannon-Fano algorithm. To do this, run the following command:

```bash
make pack_sf INPUT_FILE=example.txt
```

This will create a new file named `example.sf` with a compressed size.

Similarly, if you have an archived file named `example.sf` and you want to extract its contents, execute the following command:

```bash
make unpack_sf INPUT_FILE=example.sf
```

## Resulting File Size

After packing a file using the Shannon-Fano algorithm, you can expect the resulting archived file to have a reduced size compared to the original file. For example run:
```bash
ls -l
```
To make sure the size reduced compare to original file
```bash
-rw-r--r--  1 user  staff  3471 Feb 13 18:24 example.sf
-rw-r--r--  1 user  staff  5719 Feb 13 18:18 example.txt
```
(From 5719 to 3471 bytes )
