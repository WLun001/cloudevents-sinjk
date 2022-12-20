package main

import "time"

type PipelineRunModel struct {
	PipelineRun struct {
		Kind       string `json:"kind"`
		APIVersion string `json:"apiVersion"`
		Metadata   struct {
			Name              string    `json:"name"`
			GenerateName      string    `json:"generateName"`
			Namespace         string    `json:"namespace"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			Generation        int       `json:"generation"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				TektonDevPipeline string `json:"tekton.dev/pipeline"`
			} `json:"labels"`
			Annotations struct {
				KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
			} `json:"annotations"`
			ManagedFields []interface {
			} `json:"managedFields"`
		} `json:"metadata"`
		Spec struct {
			PipelineRef struct {
				Name string `json:"name"`
			} `json:"pipelineRef"`
			Params []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"params"`
			ServiceAccountName string `json:"serviceAccountName"`
			Timeout            string `json:"timeout"`
			Workspaces         []struct {
				Name                  string `json:"name"`
				PersistentVolumeClaim struct {
					ClaimName string `json:"claimName"`
				} `json:"persistentVolumeClaim"`
			} `json:"workspaces"`
		} `json:"spec"`
		Status struct {
			Conditions []struct {
				Type               string    `json:"type"`
				Status             string    `json:"status"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				Reason             string    `json:"reason"`
				Message            string    `json:"message"`
			} `json:"conditions"`
			StartTime time.Time `json:"startTime"`
			TaskRuns  interface {
			} `json:"taskRuns"`
			PipelineSpec interface {
			} `json:"pipelineSpec"`
			FinallyStartTime time.Time `json:"finallyStartTime"`
		} `json:"status"`
	} `json:"pipelineRun"`
}
