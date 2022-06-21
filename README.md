# pureget - Caputre Pure Storage API Info

Simple GET, POST for API pure storage routines.

## Usage

Get data from a pure storage array to be used for populating Zabbix.

`pureget -i <array-ip-fqdn> -r request`

Obtain a user token 
  
## Examples

### PODs

Get all the pods
`pureget -endpoint /pod -ip flasharray1.testdrive.local -token af71f0a4-45d8-8be1-4d4c-98536fbfe81f`

### POD Lag

Get all the replica-links with log
`pureget -endpoint /pod/replica-link?lag=true -ip flasharray1.testdrive.local -token af71f0a4-45d8-8be1-4d4c-98536fbfe81f`
  
### Volumes

Get all the replica-links with log
`pureget -endpoint /volume -ip flasharray1.testdrive.local -token af71f0a4-45d8-8be1-4d4c-98536fbfe81f`
  
  
