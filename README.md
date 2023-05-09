# pureget - Caputre Pure Storage API Data

Grab simple GET requests from a pure storage with a user token.

## Usage

Get data from a pure storage array to be used for populating Zabbix.

`pureget -ip <array-ip-fqdn> -endpoint <request> -token <token>`

 
## Examples

To use these examples you need a `token` and the `ip-address` of the pure storage to get.

The returned data will be a JSON object. It would be helpful to run this through `jq` or a similar
parser / formatter.

### PODs

Get all the pods:

`pureget -endpoint /pod -ip flasharray1.testdrive.local -token af71f0a4-45d8-8be1-4d4c-98536fbfe81f`

### POD Lag

Get all the replica-links with lag:

`pureget -endpoint /pod/replica-link?lag=true -ip flasharray1.testdrive.local -token af71f0a4-45d8-8be1-4d4c-98536fbfe81f`
  
### Volumes

Get all pure storage volumes defined:

`pureget -endpoint /volume -ip flasharray1.testdrive.local -token af71f0a4-45d8-8be1-4d4c-98536fbfe81f`
  
  
