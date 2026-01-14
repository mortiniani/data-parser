# SPDX-License-Identifier: MIT

"""
data-parser
-----------

A simple data parser for reading and processing data from various sources.

Usage
-----

To use the data parser, simply import the parser and use its methods to load and process data.

Installation
------------

To install the data parser, run the following command in your terminal:

```bash
pip install .
```

Prerequisites
------------

* Python 3.7+
* pip

API
---

### `DataParser`

#### `__init__(self, source=None)`

Initializes the data parser with a data source.

* `source`: The data source. Can be a file path, a URL, or a generator.

#### `read(self)`

Reads data from the source and returns it as a pandas DataFrame.

* Returns: A pandas DataFrame containing the parsed data.

#### `parse(self, data)`

Parses the given data and returns a pandas DataFrame.

* `data`: The data to be parsed. Can be a string, a file path, or a pandas DataFrame.

* Returns: A pandas DataFrame containing the parsed data.

#### `save(self, data, filename)`

Saves the parsed data to a file.

* `data`: The data to be saved. Can be a pandas DataFrame or a string.
* `filename`: The file path where the data will be saved.

### `load_data(source, file_format)`

Loads data from a file or URL and returns it as a pandas DataFrame.

* `source`: The data source. Can be a file path or a URL.
* `file_format`: The file format (e.g. 'csv', 'json', etc.).

* Returns: A pandas DataFrame containing the loaded data.

### `parse_csv(data)`

Parses a CSV string and returns a pandas DataFrame.

* `data`: The CSV string to be parsed.

* Returns: A pandas DataFrame containing the parsed data.

### `parse_json(data)`

Parses a JSON string and returns a pandas DataFrame.

* `data`: The JSON string to be parsed.

* Returns: A pandas DataFrame containing the parsed data.

Tests
-----

To run the tests, use the following command:

```bash
python -m unittest tests
```

Development
------------

To contribute to the data parser, fork the repository and submit a pull request.
"""

import pandas as pd
from pathlib import Path
import json
import csv
import urllib.request
import os

class DataParser:
    def __init__(self, source=None):
        self.source = source

    def read(self):
        if isinstance(self.source, str):
            if self.source.startswith('http'):
                data = urllib.request.urlopen(self.source).read()
            else:
                with open(self.source, 'r') as f:
                    data = f.read()
        elif callable(self.source):
            data = self.source()
        else:
            raise ValueError('Invalid source type')

        with pd.option_context('display.max_rows', None, 'display.max_columns', None):
            return pd.read_csv(data)

    def parse(self, data):
        if isinstance(data, str):
            with open(data, 'r') as f:
                return pd.read_csv(f)
        elif isinstance(data, Path):
            return pd.read_csv(data)
        elif isinstance(data, pd.DataFrame):
            return data
        else:
            raise ValueError('Invalid data type')

    def save(self, data, filename):
        if isinstance(data, pd.DataFrame):
            data.to_csv(filename, index=False)
        elif isinstance(data, str):
            with open(filename, 'w') as f:
                f.write(data)
        else:
            raise ValueError('Invalid data type')

def load_data(source, file_format):
    if file_format == 'csv':
        return pd.read_csv(source)
    elif file_format == 'json':
        with open(source, 'r') as f:
            return pd.DataFrame(json.load(f))
    else:
        raise ValueError('Invalid file format')

def parse_csv(data):
    return pd.read_csv(data)

def parse_json(data):
    return pd.DataFrame(json.loads(data))