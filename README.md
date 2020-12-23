# Protein Folding on Digital Ocean Infrastructure
**WARNING:** This script will generate a cluster with a projected cost of 400 per month. 
Make sure to destroy the cluster after you've learned all you can. Running the cluster for a few hours should only cost you a few dollars 😉

## Requirments
- [doctl](https://github.com/digitalocean/doctl)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [golang](https://golang.org/dl/)
- [A DO API Key](https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/)

## Instructions
1. [Setup up doctl](https://github.com/digitalocean/doctl#authenticating-with-digitalocean) with your new api token
2. Add your api token to your environment variables like so

```
$ export DIGITALOCEAN_ACCESS_TOKEN=<your-access-token>
````

3. In the root of the project run 

```
$ go run script.go
```

You're all set! Check your cloud console for more details 🎉🥳
