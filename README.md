# ds-cluster-utility

Utilities to help with darkstores cluster in multiple countries.

#### Commands you will able to use:
- **use <country_code>**: switch cluster context using country code
- **con <country_code>**: connect database using country code
- **country**: list all the countries and details such as port assigned to them (port number will remain constant)
- **backup-db <country_code> <file_name>**: backup demand-planner database from prod
- **restore-db <country_code> <file_name>**: restore database file in local db(will ask password 2 times. You should use demand_planner)

## Prerequiste:
- go compiler
- make
- This requires to have gcloud auth and kubernetes all configured locally
- cloud_sql_proxy should be runnable from anywhere

## Install
```
make install
```
This will install binaries in path /usr/local/bin:

# Configuration

IMPORTANT: Set the env variable DH_DARKSTORE_INFRA_PATH path for the dh-darkstores-infra source code in your local machine. The below example considers the infra code is at path *~/Projects/python/dh-darkstores-infra*
```
export DH_DARKSTORE_INFRA_PATH=~/Projects/python/dh-darkstores-infra
```
Note: Put this in your shell rc file (example:zshrc or bashrc).

# use command usage

Switch context to Kuwait

```
use kw
```
Output:
```
Switched to context "gke_dh-darkstores-live_europe-west1-d_live-europe-west1-v3".
```

Now you should able to use kubectl for Kuwait cluster example:

```
kubectl get pod
```

Another example, Switch context to sg

```
use sg
```

# con command usage

start cloudsql client for kw:

```
con kw
```

To start cloud sql for replica db:

```
con sg r
```

# country command usage

```
country
```

Output:
```
1-code|2-country-code|3-country|8-port
arg|ar|Argentina|54010
bh|bh|Bahrain|54017
bd|bd|Bangladesh|54018
bo|bo|Bolivia (Plurinational State of)|54026
bg|bg|Bulgaria|54034
kh|kh|Cambodia|54038
cl|cl|Chile|54044
cz|cz|Czechia|54059
do|do|Dominican Republic|54063
eg|eg|Egypt|54065
fi|fi|Finland|54075
hk|hk|Hong Kong|54100
hu|hu|Hungary|54101
jo|jo|Jordan|54114
kr|kr|Korea, Republic of|54119
kwt|kw|Kuwait|54120
la|la|Lao People's Democratic Republic|54122
my|my|Malaysia|54134
mm|mm|Myanmar|54152
om|om|Oman|54167
pk|pk|Pakistan|54168
py|py|Paraguay|54173
ph|ph|Philippines|54175
qat|qa|Qatar|54180
ro|ro|Romania|54182
sa|sa|Saudi Arabia|54195
sg|sg|Singapore|54200
se|se|Sweden|54214
tw|tw|Taiwan, Province of China|54217
th|th|Thailand|54220
ae|ae|United Arab Emirates|54233
us|us|United States of America|54235
uy|uy|Uruguay|54237
stg|stg|Staging|54249
test|test|Test|54250
```

# How it works?

The `ds-cluster-list` parses the dh-darkstores-infra terraform source code to form details like country-code, cluster-name, etc.
