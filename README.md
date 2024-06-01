tool for getting the Principal Token with using different PartnerAccountIds 

use the flag -i for help

```bash
Usage of jwt-go:
 -c string
        A cartId input for the session call

  -o string
        an opaque token input provided by the user (default "")
  -p string
        partnet account id for different sessions
  -t    adding the template field
  ```

  examples: 
  
  ```bash
    jwt-go

    jwt-go -p _
  
    jwt-go -p _ -t 

    jwt-go -o _

    jwt-go -o _ -t 

    jwt-go -c _ 

```

example output: 

```bash
--------------------------
..........................................................................................
```
```json
{
    "Client-Info": "",
    "Authorization": "",
    "Debug-Trace": ""
} 
```

if -o (opeaqueToken input) and -p (partnerAccountId input) are used together, the -o input (opeaqueToken input) will be ignored and only -p (partnerAccountId input) will be applied 

symlink jwt-go so can be used everywhere in the terminal
```bash
sudo ln -s "$(pwd)/jwt-go" /usr/local/bin/jwt-go
```

after setting the JWT_GO_USE_TABLE env to true the table mode will be the default 