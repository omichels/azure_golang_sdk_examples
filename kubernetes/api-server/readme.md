## Example about setting the IP-Filter of AKS API-Server

Howto to limit the list of authorized IP sources allowed connections to your K8S API Server


example call:

```
err := WhitelistIPOnAPIServer("89.76.89.76", "xx010101-xxxx-yyyy-xxxx-xxx0101010" , "my-group", "my-cluster")
```


### Authorization

works via environment variables like this:


```
export AZURE_TENANT_ID="xxxxxx-yyyy-zzzz-xxxx-xxxxxxxxxx"
export AZURE_CLIENT_ID="xxxxxx-yyyy-zzzz-xxxx-xxxxxxxxxx"
export AZURE_CLIENT_SECRET="a-very-secure-password"
export AZURE_SUBSCRIPTION_ID="xxxxxx-yyyy-zzzz-xxxx-xxxxxxxxxx"
```
