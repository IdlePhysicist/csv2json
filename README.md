# CSV2JSON
Convert CSV files to JSON files with Go.

*Forked from [Ahmad-Magdy/CSV-To-JSON-Converter](https://github.com/Ahmad-Magdy/CSV-To-JSON-Converter).*

## How to build:

1. After downloading / cloning navigate to the repo directory.
2. Build an executable with `go build`
3. To run `./csv2json data.csv`

**Optional Step:**

4. Link the executable to a handy location.
```shell
cd /usr/local/bin
ln -s ~/path/to/csv2json/csv2json csv2json
```
Then to run simply do the following:
```shell
$ csv2json data.csv
```

## Example:
If i have a csv file contains a people data **people.csv**:
```csv
Id,Name,Age
1,Ahmad,21
2,Ali,50
```
and we need to convert it to json file:

`$ csv2json people.csv`

after writing this command you will get another file in the same directory called **people.json** with the data in the new format:
```json
[
  {
    "Id": 1,
    "Name": "Ahmad",
    "Age": 21
  },
  {
    "Id": 2,
    "Name": "Ali",
    "Age": 50
  }
]
```

and that's it.

## License:
[The MIT License](https://github.com/IdlePhysicist/csv2json/blob/master/LICENSE)
