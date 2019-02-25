# Krazy Cow

```
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


# Endpoints

* `/` - talk with a cow
* `/setfree` - set free a cow (cause a process to exit(1))
* `/healthz` - healthcheck; also display current mood of the cow


# Configuration

* **With a config file**

Create `cowconfig.yaml` and put it in the same directory as the app (binary file) or in the `/config/` directory. The latter location is used when running a container with a config file mounted from ConfigMap.

* **With environment variable**

Start a cow with environment variable set with the following scheme: 

For the following variable

```
cow:
  say: "Moo"

```

override with `KC_COW_SAY="Hello"`

## TLS

By default tls is **disabled**. To enable make sure your config the following entries:

```
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
