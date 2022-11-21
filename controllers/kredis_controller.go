/*
Copyright 2022.

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

	kerrors "k8s.io/apimachinery/pkg/api/errors"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	stablev1 "github.com/hkpark130/operator/api/v1"
)

// KRedisReconciler reconciles a KRedis object
type KRedisReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=stable.example.com,resources=kredis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=stable.example.com,resources=kredis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=stable.example.com,resources=kredis/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KRedis object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *KRedisReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	log.Log.Info("HK TEST HK TEST HK TEST HK TEST HK TEST")

	// TODO(user): your logic here
	reqKRedis := &stablev1.KRedis{}
	err := r.Get(ctx, req.NamespacedName, reqKRedis)
	if err != nil {
		if kerrors.IsNotFound(err) {
			log.Log.Info("KRedis resource not found.")
			return ctrl.Result{}, nil
		}
		// custom resource를 가져오지 못 함
		log.Log.Error(err, "Failed to get KRedis")
		return ctrl.Result{}, err
	} // Reconcile 함수 구현

	deployment := &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: reqKRedis.Name, Namespace: reqKRedis.Namespace}, deployment)
	if err != nil && kerrors.IsNotFound(err) {
		log.Log.Info("Creating a new Deployment.", reqKRedis.Name, reqKRedis.Namespace)
	} else if err != nil {
		log.Log.Info("Failed to get Deployment.")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KRedisReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stablev1.KRedis{}).
		Complete(r)
}
