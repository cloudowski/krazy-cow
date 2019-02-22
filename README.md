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

This is Krazy Cow that is Kubernetes friendly but sometimes moody and needs special attention. It helps you learn how containers and whole Cloud Native system can work together to bring more speed to your environments. 

It also brings fun and joy!


# Endpoints

* `/` - talk with a cow
* `/setfree` - set free a cow (cause a process to exit(1))


# Configuration

## With a config file

Create `cowconfig.yaml`

## With environment variable

Start a cow with environment variable set with the following scheme: 

For the following variable

```
cow:
  say: "Moo"

```

override with `TC_COW_SAY="Hello"`