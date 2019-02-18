# Trapped cow

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