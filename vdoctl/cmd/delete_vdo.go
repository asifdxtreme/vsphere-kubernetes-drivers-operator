package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var deleteVdoCmd = &cobra.Command{
	Use:   "vdo",
	Short: "Delete vSphere Kubernetes Driver Operator",
	Long: `This command deletes the VDO deployment and associated artifacts from the cluster targeted by --kubeconfig flag or KUBECONFIG environment variable.
Currently the command supports vanilla k8s cluster`,
	Example: "vdoctl delete vdo",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Check the vdoDeployment Namespace and confirm if VDO operator is running in the env
		getVdoNamespace(ctx)

		ns := corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: VdoCurrentNamespace,
			},
		}

		err := K8sClient.Delete(ctx, &ns, &client.DeleteOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			cobra.CheckErr(fmt.Errorf("Error occurred deleting VDO,  %v", err))
		}

		clusterRBList := rbacv1.ClusterRoleBindingList{}
		clusterRoleList := rbacv1.ClusterRoleList{}
		labelSelector, _ := labels.Parse("managedby=vdo")

		err = K8sClient.List(ctx, &clusterRBList, &client.ListOptions{LabelSelector: labelSelector})
		if err != nil {
			cobra.CheckErr(fmt.Errorf("Error occurred deleting VDO %v", err))
		}

		for _, crb := range clusterRBList.Items {
			err = K8sClient.Delete(ctx, &crb, &client.DeleteOptions{})
			if err != nil && !apierrors.IsNotFound(err) {
				cobra.CheckErr(fmt.Errorf("Error occurred when deleting VDO,  %v", err))
			}
		}

		err = K8sClient.List(ctx, &clusterRoleList, &client.ListOptions{LabelSelector: labelSelector})
		if err != nil {
			cobra.CheckErr(fmt.Errorf("Error occurred deleting VDO %v", err))
		}

		for _, role := range clusterRoleList.Items {
			err = K8sClient.Delete(ctx, &role, &client.DeleteOptions{})
			if err != nil && !apierrors.IsNotFound(err) {
				cobra.CheckErr(fmt.Errorf("Error occurred when deleting VDO,  %v", err))
			}
		}

		fmt.Println("This command deletes the VDO operator deployment eventually the operator pods will be deleted")

	},
}

func init() {
	deleteCmd.AddCommand(deleteVdoCmd)
}
