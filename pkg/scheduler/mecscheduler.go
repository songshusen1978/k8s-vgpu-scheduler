package scheduler

import (
	//"4pd.io/k8s-vgpu/pkg/k8sutil"
	//"4pd.io/k8s-vgpu/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

func (s *Scheduler) StartMecScheduler(informerFactory informers.SharedInformerFactory) {
	informer := informerFactory.Apps().V1().Deployments().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: s.onAddDeployment,
	})
	informerFactory.Start(s.stopCh)
}

func (s *Scheduler) onAddDeployment(obj interface{}) {
	dep, ok := obj.(*appsv1.Deployment)
	if !ok {
		klog.Errorf("unknown add object type, not deployment")
		return
	}
	//node name and annotaions
	if dep.Spec.Template.Spec.NodeName == "" {
		klog.Infof("node name set")
		return
	}

	/*

		nodeID, ok := pod.Annotations[util.AssignedNodeAnnotations]
		if !ok {
			return
		}
		ids, ok := pod.Annotations[util.AssignedIDsAnnotations]
		if !ok {
			return
		}
		if k8sutil.IsPodInTerminatedState(pod) {
			s.delPod(pod)
			return
		}
		podDev := util.DecodePodDevices(ids)
		s.addPod(pod, nodeID, podDev)
	*/
}
