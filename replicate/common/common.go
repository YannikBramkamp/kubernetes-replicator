package common

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Replicator interface {
	Run()
	Synced() bool
	NamespaceAdded(ns *v1.Namespace)
}

func PreviouslyPresentKeys(object *metav1.ObjectMeta) (map[string]struct{}, bool) {
	keyList, ok := object.Annotations[ReplicatedKeysAnnotation]
	if !ok {
		return nil, false
	}

	keys := strings.Split(keyList, ",")
	out := make(map[string]struct{})

	for _, k := range keys {
		out[k] = struct{}{}
	}

	return out, true
}
