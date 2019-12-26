[![CircleCI](https://circleci.com/gh/cloudowski/krazy-cow.svg?style=svg)](https://circleci.com/gh/cloudowski/krazy-cow)

# Krazy Cow

```ascii
           (    )
            (oo)
   )\.-----/(O O)
  # ;       / u
    (  .   |} )
     |/ ".;|/;
     "     " "
```

This is **Krazy Cow** - a Kubernetes friendly animal that it also moody and requires special attention. It helps you learn how containers and whole Cloud Native system can work together to bring more speed to your environments. 

It also fun to play with!

# Features

* Configuration provided by yaml file (with all defaults placed in [config/defaultconfig.yaml](config/defaultconfig.yaml)) and environment variables
* Plain http and additional, optional https endpoint - key and certificate files required
* Logging with colored output and different, configurable severity levels (see [https://github.com/op/go-logging](https://github.com/op/go-logging))
* Healthcheck endpoint that returns error when cow mood is below a defined threshold
* Artificial "mood changer" can be attached to a cow to decrease the mood every N seconds
* Logging all requests (*access logs*)
* Different output for text (curl,wget)  and browser clients
* "Shepherd" (redis) can be a destination for access logs 
* When "pasture" location is provided to a cow with tufts (files `tuft*`) then it is being periodically eaten by a cow increasing its mood AND producing a milk; if a shepherd is available then it is sent to a `dairy` (a key with that name in redis); when there are no tufts left in pasture a cow mood is decreased
* Cow can be set free using `/setfree` endpoint. Optionally it can be secured with basic http authentication.
* Two custom metrics are available at `/metrics` endpoint:
  * `cow_requests` with request count
  * `cow_mood` with current mood of a cow


# Endpoints

* `/` - talk with a cow
* `/metrics` - standard Prometheus endpoint
* `/setfree` - set free a cow (cause a process to exit(1))
* `/healthz` - healthcheck; also display current mood of the cow


# Configuration

Configuration is handled by viper ([https://github.com/spf13/viper](https://github.com/spf13/viper)) and thus yaml files and environment variables can be used.

* **With a config file**

Create `cowconfig.yaml` and put it in the same directory as the app (binary file) or in the `/config/` directory. The latter location is used when running a container with a config file mounted from ConfigMap.

* **With environment variable**

Start a cow with environment variable set with the following scheme: 

For the following variable

```yaml
cow:
  say: "Moo"

```

override with `KC_COW_SAY="Hello"`

## TLS

By default tls is **disabled**. To enable make sure your config the following entries:

```yaml
http:
    port: 8080 # <-- this is default port
    tls:
        enabled: true
        cert: config/server.crt # <-- path to a cert file
        key: config/server.key # <-- path to a key file
```

## Authentication

Currently only `/setfree` is protected with http basic authentication. To set it up you need to point a credentials file with `http.auth.credentials` variable and enable it with `http.auth.enabled` set to `true`.

Credentials must be put in a file in this very simple format: **`USERNAME:PASSWORD`**.

## Access Kubernetes API

To enable access to Kubernetes api and enable features of herd discovery a rolebinding is required. To create it simply add it with a command:

```shell
kubectl create clusterrolebinding cowview --clusterrole=view --serviceaccount=default:default
```
