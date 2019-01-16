# code init --code-type CodeType --name Name

```
- Download template from a website
- Extract it into directory
```

# code new --github-auth-token Token --github-repo abhiyerra/landingcrew --code-type type

```
- Check for landingcrew.json
- Check for Readme
- Check that there are Github Issues
- Add LandingCrewWorker as a contributor to the repo.

    Return: {"ID": id}
```

# code get --id id


```
Output:

Return as JSON: type Output struct {
        ID         string
        Identifier string `json:",omitempty"`
        Status     Status
        Fields     map[string]interface{}
        Bid int32 `json:",omitempty"`

        Error string `json:",omitempty"`
}
```

# code approve --id id

```
POST request
```

# code list

```
Output: [
struct {
            ID         string
            Identifier string `json:",omitempty"`
            Status     Status
            Fields     map[string]interface{}
            Bid int32 `json:",omitempty"`

            Error string `json:",omitempty"`
    },

]
```

# content type-list

```
List ContentType

    [
        {
            "Type": "TypeName",
            "Description": "Description"
        }
    ]
```

# content init --content-type Type --name Name

```
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
```

# content new --type type --file=file.md

# content get

```
Return as JSON: type Output struct {
        ID         string
        Identifier string `json:",omitempty"`
        Status     Status
        Fields     map[string]interface{}

        Error string `json:",omitempty"`
}
```

# content list

Same as `code list`

# data new
# data get --id id

```
Return as JSON: type Output struct {
        ID         string
        Identifier string `json:",omitempty"`
        Status     Status
        Fields     map[string]interface{}

        Error string `json:",omitempty"`
}
```

# data list

Same as `code list`

# action new

```
Ignore

```

# action get --id id

```
Return as JSON: type Output struct {
        ID         string
        Identifier string `json:",omitempty"`
        Status     Status
        Fields     map[string]interface{}

        Error string `json:",omitempty"`
}
```

# action list

Same as `code list`