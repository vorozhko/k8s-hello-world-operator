# Hello World Kubernetes Operator

This project implements a simple Kubernetes operator that manages HelloWorld resources. The operator is built using the Operator SDK and follows the best practices for Kubernetes operators.

## Project Structure

- **cmd/main.go**: Entry point of the operator. Initializes the manager and starts the operator.
- **config/crd/bases/example.com_helloworlds.yaml**: Defines the Custom Resource Definition (CRD) for HelloWorld resources.
- **config/default/kustomization.yaml**: Kustomize configuration for default resources.
- **config/manager/kustomization.yaml**: Kustomize configuration for manager deployment.
- **config/rbac/role.yaml**: Defines the Role for the operator with necessary permissions.
- **config/rbac/role_binding.yaml**: Links the Role to the ServiceAccount.
- **config/rbac/service_account.yaml**: Defines the ServiceAccount for the operator.
- **controllers/helloworld_controller.go**: Implements the HelloWorld controller.
- **api/v1/groupversion_info.go**: Contains group and version information for the API.
- **api/v1/helloworld_types.go**: Defines the HelloWorld resource's Go struct and methods.
- **Dockerfile**: Instructions to build the Docker image for the operator.
- **Makefile**: Commands for building, testing, and deploying the operator.
- **PROJECT**: Metadata about the project, including name and version.
- **go.mod**: Go module definition file.

## Getting Started

To build and run the operator, follow these steps:

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd hello-world-operator
   ```

2. **Build the operator(Optional)**:
   ```
   make build
   ```

3. **Deploy the operator**:
   ```
   make deploy
   ```

4. **Interact with the HelloWorld resource**:
   You can create, update, and delete HelloWorld resources using `kubectl`.
   ```
   kubectl apply -f example/helloworld.example.yaml
   kubectl get helloworlds.example.com
   NAME                AGE
   helloworld-sample   3m33s
   ```

## License

This project is licensed under the MIT License. See the LICENSE file for details.