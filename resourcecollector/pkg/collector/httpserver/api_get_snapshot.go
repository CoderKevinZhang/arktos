/*
Copyright 2020 Authors of Arktos.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httpserver

import (
	"k8s.io/klog"
	"net/http"

	"k8s.io/kubernetes/globalscheduler/pkg/scheduler/utils"
	"k8s.io/kubernetes/resourcecollector/pkg/collector"

	"github.com/emicklei/go-restful"
)

// Schedule get snapshot
func GetSnapshot(req *restful.Request, resp *restful.Response) {
	col, err := collector.GetCollector()
	if err != nil {
		klog.Errorf("get new collector failed, err: %s", err.Error())
		utils.WriteFailedJSONResponse(resp, http.StatusInternalServerError, utils.InternalServerError())
		return
	}

	snapshot, err := col.GetSnapshot()
	if err != nil {
		klog.Errorf("Collector snapshot failed! err : %s", err)
		utils.WriteFailedJSONResponse(resp, http.StatusInternalServerError, utils.InternalServerError())
		return
	}

	resp.WriteAsJson(snapshot)
}
