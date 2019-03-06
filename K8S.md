# Features

* pod
  * env from configmap
  * confifg file configmap 
  * initcontainer for pasture
  * healthcheck - mood decreaser
  * restartpolicy - /setfree
* volume - configfile
* secret - auth, tls certs
* service
  * cow to redis
* rs
  * scaling
* deployment
  * scaling
  * update between versions
* pv, pvc
  * pasture as pv
  * shared pasture for many cows
  * resiliency - redis with pv, failure, proof
* envs / namespaces
  * separate ns with redis 
  * redis in cloud - configurable
* rbac
  * discover other cows from herd
* fun
  * gather milk, prepare pasture

PLANNED
* garfana dashboard with prom metrics
  * eaten turfs
  * restarts
  * produced milk - herd ad individual
* build with jenkins
  * publish to harbor
* use vault
  * to generate cert
  * use k8s auth
  * use external auth for /setfree or other endpoint
* connect with consul
* connect with istio
* use networkpolicy to limit traffic


