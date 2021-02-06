/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/sarweshsuman/kubernetes-custom-controller-examples/group-deployment/api/v1alpha1"
)

// GroupDeploymentReconciler reconciles a GroupDeployment object
type GroupDeploymentReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps.example,resources=groupdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.example,resources=groupdeployments/status,verbs=get;update;patch

func (r *GroupDeploymentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("groupdeployment", req.NamespacedName)

	log.V(1).Info("Reconciling", "GroupDeployment", req)
	log.V(1).Info("Reconciling", "GroupDeployment", req.Name)

	/* Read the GroupDeployment Object with name req.Name */
	var groupDeployment appsv1alpha1.GroupDeployment
	if err := r.Get(ctx, req.NamespacedName, &groupDeployment); err != nil {
		if errors.IsNotFound(err) {
			log.Info("Deleted", "GroupDeployment", req)
			return ctrl.Result{}, client.IgnoreNotFound(err)
		} else {
			log.Error(err, "unable to fetch GroupDeployment object.")
		}
		return ctrl.Result{}, err
	}

	log.V(1).Info("Found object", "Group Deployment", groupDeployment)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *GroupDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.GroupDeployment{}).
		Complete(r)
}
