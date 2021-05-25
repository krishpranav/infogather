# infogather
A simple go web app and a cli tool for gathering information

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation
```
git clone https://github.com/krishpranav/infogather
cd infogather
go get 
```

- cli:
```
go run main.go
```

- web app
```
go run api/server.go
```

- you want to add some of the api keys for experiencing full tool
```js
{
    "melissaKeyCred": "Paste Melissa Key with Credits Here",
    "hibpKey": "Paste Have I Been Pwned Key Here",
    "dataGovKey": "Paste Data.gov Key Here"
}
```

## Package-Types

### Court Cases

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Case Law](https://github.com/krishpranav/infogather/tree/main/pkg/noauth/caselaw)           | Court Case Search                            |  `none`  |

### Discord

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Discord Token Lookup](https://github.com/krishpranav/infogather/tree/main/pkg/noauth/discord)| Discord Token Lookup                        |  `none`  |

### IP Address

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [IPV4 Address Lookup](https://github.com/krishpranav/infogather/tree/main/pkg/noauth/ip)     | IPV4 Address Lookup                          |  `none`  |

### Multi-Use

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Have I Been Pwned](https://github.com/krishpranav/infogather/tree/main/pkg/authpaid/hibp)   | Email and Password Vulnerability - (Breaches)|  `paid`  |
| [Melissa](https://github.com/krishpranav/infogather/tree/main/pkg/authfree/melissa)          | Lookups - Email, Phone Number, IP Address    |  `free`  |


### Username

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [Username Lookup](https://github.com/krishpranav/infogather/tree/main/pkg/noauth/userlookup) | Username Lookup - (Comparable to Sherlock)   |  `none`  |

### Vehicle

| Package                                                                                    | Description                                  |   Auth   |
| :----------------------------------------------------------------------------------------: | -------------------------------------------- | :------: |
| [VIN Lookup](https://github.com/krishpranav/infogather/tree/main/pkg/noauth/vin)             | Vehicle Identification Number Lookup         |  `none`  |