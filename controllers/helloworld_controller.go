package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
    ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	examplev1 "hello-world-operator/api/v1"
)

type HelloWorldReconciler struct {
	client.Client
}

func (r *HelloWorldReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log := log.FromContext(ctx)

	log.Info("Reconciling HelloWorld", "request", req)

	helloWorld := &examplev1.HelloWorld{}
	err := r.Get(ctx, req.NamespacedName, helloWorld)
	if err != nil {
		log.Error(err, "unable to fetch HelloWorld")
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// Here you can add your reconciliation logic

	return reconcile.Result{}, nil
}

func (r *HelloWorldReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplev1.HelloWorld{}).
		Owns(&corev1.Pod{}).
		Complete(r)
}
