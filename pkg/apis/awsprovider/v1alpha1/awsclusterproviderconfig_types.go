/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSClusterProviderSpec is the providerConfig for AWS in the cluster
// object
// +k8s:openapi-gen=true
type AWSClusterProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The AWS Region the cluster lives in.
	Region string `json:"region,omitempty"`

	// SSHKeyName is the name of the ssh key to attach to the bastion host.
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// CACertificate is a PEM encoded CA Certificate for the control plane nodes.
	CACertificate []byte `json:"caCertificate,omitempty"`

	// CAPrivateKey is a PEM encoded PKCS1 CA PrivateKey for the control plane nodes.
	CAPrivateKey []byte `json:"caKey,omitempty"`

	// NetworkSpec encapsulates all things related to AWS network.
	NetworkSpec NetworkSpec `json:"networkSpec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&AWSClusterProviderSpec{})
}

// NetworkSpec encapsulates all things related to AWS network.
type NetworkSpec struct {
	// VPC configuration.
	// +optional
	VPC VPCSpec `json:"vpc,omitempty"`

	// Subnets configuration.
	// +optional
	Subnets []SubnetSpec `json:"subnets,omitempty"`
}

// VPCSpec configures an AWS VPC.
type VPCSpec struct {
	// ID is the vpc-id of the VPC this provider should use to create resources.
	// When this value is specified the VPC is considered unmanaged.
	// +optional
	ID *string `json:"id,omitempty"`

	// CidrBlock is the CIDR block to be used when the provider creates a managed VPC.
	// Defaults to 10.0.0.0/16.
	// Ignored if ID is specified.
	// +optional
	CidrBlock *string `json:"cidrBlock,omitempty"`
}

// SubnetSpec configures an AWS Subnet.
type SubnetSpec struct {
	// ID is the unique identifier of the subnet this provider should use to create resources.
	// When this value is specified the subnet is considered unmanaged.
	// +optional
	ID string `json:"id,omitempty"`

	// CidrBlock is the CIDR block to be used when the provider creates a managed VPC.
	// Must be specified when VPCSpec CidrBlock is specified.
	// Ignored if ID is specified.
	// +optional
	CidrBlock *string `json:"cidrBlock,omitempty"`

	// AvailabilityZone defines the availability zone to use for this subnet in the cluster's region.
	// Ignored if ID is specified.
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Public defines the subnet as a public subnet. Refer to the AWS documentation for further information.
	// Ignored if ID is specified.
	// +optional
	Public *bool `json:"public,omitempty"`
}
