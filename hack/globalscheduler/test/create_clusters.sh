#!/usr/bin/env bash

# Copyright 2020 Authors of Arktos - file created.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FILE="/home/ubuntu/go/src/k8s.io/arktos/globalscheduler/test/yaml/sample_1000_clusters.yaml"

function create_pod {
# Create multiple YAML objects from stdin
cat <<EOM >> $FILE
apiVersion: globalscheduler.com/v1
kind: Cluster
metadata:
  name: $1
  namespace: default
spec:
  cpucapacity: 8
  eipcapacity: 3
  flavors:
  - flavorid: "1"
    totalcapacity: 1000
  geolocation:
    area: $2
    city: $3
    country: $5
    province: $4
  ipaddress: 18.236.217.191
  memcapacity: 256
  operator:
    operator: globalscheduler
  region:
    availabilityzone: az1
    region: $2
  serverprice: 10
  storage:
  - storagecapacity: 256
    typeid: sas
---
EOM
}

for ((i = 0 ; i < $(($1)) ; i++)); do
    name="cluster-$(($i))"
    area="area-$(($i))"
    city="city-$(($i))"
    province="province-$(($i))"
    country="US"
    create_pod $name $area $city $province $country
done
