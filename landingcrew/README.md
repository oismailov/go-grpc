# LandingCrew CLI

LandingCrew CLI is a tool to quickly create work tasks using the command
line.


# Usage

```
configure github

code init
   --type (LambdaRuby,LambdaPython,LambdaNodejs)
   --name <Name>
   --dir

   - Download template from website

code new --github-auth-token|ENV['GITHUB_AUTH_TOKEN'] --github-repo <abhiyerra/landingcrew> --type <type>

    - Check for landingcrew.json
    - Check for Readme
    - Check that there are Github Issues
    - Add LandingCrewWorker as a contributor to the repo.

        Return: {"ID": id}

code get --id <id>

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}
            Bid int32 `json:",omitempty"`

            Error string `json:",omitempty"`
    }

code approve --id <id>

     POST request


code list

    [
    Output struct {
                ID         string
                Identifier string `json:",omitempty"`
                Status     Status
                Fields     map[string]interface{}
                Bid int32 `json:",omitempty"`

                Error string `json:",omitempty"`
        },

    ]

content init --type

{
        "Title": "",
        "Instructions": ""
        "FieldInstructions": [
            {
            "Name": "Name",
            "Instructions": ""
            "Type": ""
            }
        ]
}

content new --type <type> --file=@file.md
content get

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}

            Error string `json:",omitempty"`
    }

content list

data new
data get --id <id>

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}

            Error string `json:",omitempty"`
    }

data list

action new
action get --id <id>

    Return as JSON: type Output struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}

            Error string `json:",omitempty"`
    }

action list
```
