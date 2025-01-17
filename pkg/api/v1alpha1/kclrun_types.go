package v1alpha1

const (
	// KCLRunGroup represents the API group for the KCLRun resource.
	KCLRunGroup = "krm.kcl.dev"

	// KCLRunVersion represents the API version for the KCLRun resource.
	KCLRunVersion = "v1alpha1"

	// KCLRunAPIVersion is a combination of the API group and version for the KCLRun resource.
	KCLRunAPIVersion = KCLRunGroup + "/" + KCLRunVersion

	// KCLRunKind represents the kind of resource for the KCLRun resource.
	KCLRunKind = "KCLRun"

	// SourceKey is the key for the source field in the KCLRun resource.
	SourceKey = "source"

	// ParamsKey is the key for the params field in the KCLRun resource.
	ParamsKey = "params"

	// MatchConstraintsKey is the key for the match constraints field in the KCLRun resource.
	MatchConstraintsKey = "matchConstraints"
)
