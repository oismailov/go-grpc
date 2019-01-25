LandingCrew CLI

LandingCrew CLI allows quick access to a workforce for various tasks.

# Usage

## Code

### Initialize New Code

code init --code-type Type --name Name

### Create New Code from Dir

code new-from-dir --code-type type --dir .

### Create New Code from Github Repo

code new-from-repo --github-auth-token --github-repo abhiyerra/landingcrew --code-type type

    - Check for Readme
    - Check that there are Github Issues
    - Add LandingCrewWorker as a contributor to the repo.

        Return: {"ID": id}

### Get the code

code get --id id

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}
            Bid int32 `json:",omitempty"`

            Error string `json:",omitempty"`
    }

## Decide on Approval

code approve --id id

## Decide on Reject

code reject --id id --reason ""

## code list

## code list-type

# Content

## Initialize New Type of Content

content init --type Type

```json
{
        "Name": "",
        "Type": "",
        "Instructions": "",
}
```

## Create New Content

content new --config ./config.json

## Get Content

content get --id id

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}

            Error string `json:",omitempty"`
    }

## Get Content List

# Data

## Initialize New Data Task

data init

## Create New CSV

data new-csv --config config.json file

```json
{
    "Identifier": "",
    "Name": "",
    "Instructions": "",
    "InstructionVideo": "",
    "Fields": [
        {
            "Type": "",
            "Instructions": ""
        }
    ]
}
```

## Create New Airtable Data

data new --config config.json

```json
{
    "Identifier": "",
    "Name": "",
    "Instructions": "",
    "InstructionVideo": "",
    "Fields": [
        {
            "Name": "",
            "Type": "",
            "Instructions": ""
        }
    ],
    "Stream": {
        "Type": "AirtableDataType",
        "AirtableApiKey": "",
        "AirtableAppKey": "",
        "AirtableTable": ""
    }
}
```

## Create New Database

data new --config config.json


```json
{
    "Identifier": "",
    "Name": "",
    "Instructions": "",
    "InstructionVideo": "",
    "Fields": [
        {
            "Name": "",
            "Type": "",
            "Instructions": ""
        }
    ],
    "Stream": {
        "Type": "StreamDataType",
    }
}
```


data sync --id id

```json


```


## Get Data

data get --id <id>

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Error string `json:",omitempty"`
    }

### Get Csv Data

data get-csv --id <id>


## List Data

data list


## Sync Data

data sync --id id

# Actions

## Create

action new

  --file
  --title --instructions --instruction-video

### action get --id <id>

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}

            Error string `json:",omitempty"`
    }

### action list
