## Example about setting the IP-Filter of AKS API-Server

Howto to limit the list of authorized IP sources allowed connections to your K8S API Server


example call:

```
err := WhitelistIPOnAPIServer("89.76.89.76", "xx010101-xxxx-yyyy-xxxx-xxx0101010" , "my-group", "my-cluster")
```
