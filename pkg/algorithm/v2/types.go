// MIT License
//
// Copyright (c) Microsoft Corporation. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE

package v2

import (
	algo "github.com/microsoft/hivedscheduler/pkg/algorithm"
	"github.com/microsoft/hivedscheduler/pkg/api"
	apiv2 "github.com/microsoft/hivedscheduler/pkg/api/v2"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/apis/core"
)

// PodGroupState represents the internal state of pod group.
type PodGroupState string

// PodGroupSchedulingRequest represents request of pod group.
type PodGroupSchedulingRequest struct {
	vc           api.VirtualClusterName
	pinnedCellId api.PinnedCellId
	podRootGroup apiv2.PodGroupSpec
	chain        algo.CellChain
	priority     algo.CellPriority
}

// PodGroupSchedulingStatus represents internal scheduling status of pod group.
type PodGroupSchedulingStatus struct {
	preemptingPods       map[types.UID]*core.Pod
	allocatedPodGroup    AllocatedPodGroup
	virtualPlacement     groupVirtualPlacement
	physicalPlacement    groupPhysicalPlacement
	state                PodGroupState
	lazyPreemptionStatus *api.LazyPreemptionStatus
}

// AllocatedPodGroup represents a tree structure of allocated pod group.
type AllocatedPodGroup struct {
	pods                []*core.Pod
	allocatedChildGroup []*AllocatedPodGroup
}

// podGroupPlacement represents a tree structure of intra VC scheduled placement.
type podGroupPlacement struct {
	podsPlacement        []algo.CellList
	childGroupsPlacement []*podGroupPlacement
}

// PodGroupPhysicalPlacement represents physical placement of pod group.
type PodGroupPhysicalPlacement podGroupPlacement

// PodGroupVirtualPlacement represents virtual placement of pod group.
type PodGroupVirtualPlacement podGroupPlacement