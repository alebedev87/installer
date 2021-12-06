package dds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ConfigserverAttribute is a nested struct in dds response
type ConfigserverAttribute struct {
	NodeClass       string `json:"NodeClass" xml:"NodeClass"`
	NodeDescription string `json:"NodeDescription" xml:"NodeDescription"`
	MaxConnections  int    `json:"MaxConnections" xml:"MaxConnections"`
	Port            int    `json:"Port" xml:"Port"`
	NodeStorage     int    `json:"NodeStorage" xml:"NodeStorage"`
	NodeId          string `json:"NodeId" xml:"NodeId"`
	ConnectString   string `json:"ConnectString" xml:"ConnectString"`
	MaxIOPS         int    `json:"MaxIOPS" xml:"MaxIOPS"`
}