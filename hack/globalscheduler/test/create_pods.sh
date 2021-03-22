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
# locs=("NewYork NewYork NE-1 US" "Bellevue Washington NW-1 US" "Orlando Florida SE-1 US" "Austin Texas SW-1 US" "Chicago Illinois Central-1 US" "Boston Massachusettes NE-2 US" "SanFrancisco California NW-2 US" "Atlanta Georgia SE-2 US" "LasVegas Nevada SW-2 US" "Omaha Nebraska Central-2 US")

FILE="/home/ubuntu/go/src/k8s.io/arktos/globalscheduler/test/yaml/sample_1000_pods.yaml"

function create_pod {
# Create multiple YAML objects from stdin
cat <<EOM >> $FILE
apiVersion: v1
kind: Pod
metadata:
  name: $1
spec:
  resourceType: "vm"
  virtualMachine:
    name: vm$1
    image: "51dc9183-f01b-4ecf-8c55-cd8aecb17db3"
    keyPairName: "demo-keypair"
    securityGroupId: "c46e6ee2-8112-4d4e-9e60-fe4c813f7339"
    flavors:
      - flavorID: "42"
    resourceCommonInfo:
     count: 1
     selector:
       geoLocation:
         city: $3
         province: $4
         area: $2
         country: $5
       regions:
         - region: $2
           availablityZone:
           - "az1"
  nics:
    - name: "8cbc8370-0b22-4927-9b50-f16b5dc57d52"
---
EOM
}

# locsLen=${#locs[@]}
idx=0
for ((i = 0 ; i < $(($1)) ; i++)); do
    if [ $idx -eq 1000 ]
    then
      idx=0
    fi
    name="pod-$(($i))"
    area="area-$(($idx))"
    city="city-$(($idx))"
    province="province-$(($idx))"
    country="US"
    create_pod $name $area $city $province $country
    idx=$((idx+1))
done