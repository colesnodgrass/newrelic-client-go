package alerts

// nolint:golint
type AlertsIncidentPreference string

const (
	/* A condition will create a condition-level incident when it violates its critical threshold.
	Other violating conditions will create their own incidents. */
	PER_CONDITION AlertsIncidentPreference = "PER_CONDITION" // nolint:golint
	/* Each target of each condition will create an entity-level incident upon critical violation.
	Other violating targets even on the same condition will create their own incidents. */
	PER_CONDITION_AND_TARGET AlertsIncidentPreference = "PER_CONDITION_AND_TARGET" // nolint:golint
	/* A condition will create a policy-level incident when it violates its critical threshold.
	Other violating conditions will be grouped into this incident. */
	PER_POLICY AlertsIncidentPreference = "PER_POLICY" // nolint:golint
)

// nolint:golint
type AlertsMutingRuleConditionGroupInput struct {
	/* The individual MutingRuleConditions within the group. */
	Conditions []AlertsMutingRuleConditionInput `json:"conditions"`

	/* The operator used to combine all the MutingRuleConditions within the group. */
	Operator AlertsMutingRuleConditionGroupOperator `json:"operator"`
}

// nolint:golint
type AlertsMutingRuleConditionGroupOperator string

const (
	/* Match conditions by AND */
	AND AlertsMutingRuleConditionGroupOperator = "AND" // nolint:golint
	/* Match conditions by OR */
	OR AlertsMutingRuleConditionGroupOperator = "OR" // nolint:golint
)

// nolint:golint
type AlertsMutingRuleConditionInput struct {
	/* The attribute on a violation. Expects one of:

	* **accountId** - The account id
	* **conditionId** - The alert condition id
	* **policyId** - The alert policy id
	* **policyName** - The alert policy name
	* **conditionName** - The alert condition name
	* **conditionType** - The alert condition type, such as app_metric
	* **conditionRunbookUrl** - The alert condition's runbook url
	* **product** - The target product (e.g., SYNTHETICS)
	* **targetId** - The ID of the alerts target
	* **targetName** - The name of the alerts target
	* **nrqlEventType** - The NRQL event type
	* **tag** - Arbitrary tags associated with some entity (e.g., FACET from a NRQL query)
	* **nrqlQuery** - The NRQL query string */
	Attribute string `json:"attribute"`

	/* The operator used to compare the attribute's value with the supplied value(s). */
	Operator AlertsMutingRuleConditionOperator `json:"operator"`

	/* The value(s) to compare against the attribute's value. */
	Values []string `json:"values"`
}

// nolint:golint
type AlertsMutingRuleConditionOperator string

const (
	/* Where any. */
	ANY AlertsMutingRuleConditionOperator = "ANY" // nolint:golint
	/* Where contain value. */
	CONTAINS AlertsMutingRuleConditionOperator = "CONTAINS" // nolint:golint
	/* Where ends with. */
	ENDS_WITH AlertsMutingRuleConditionOperator = "ENDS_WITH" // nolint:golint
	/* Where values equals. */
	EQUALS AlertsMutingRuleConditionOperator = "EQUALS" // nolint:golint
	/* Where in value. */
	IN AlertsMutingRuleConditionOperator = "IN" // nolint:golint
	/* Where blank. */
	IS_BLANK AlertsMutingRuleConditionOperator = "IS_BLANK" // nolint:golint
	/* Where not blank. */
	IS_NOT_BLANK AlertsMutingRuleConditionOperator = "IS_NOT_BLANK" // nolint:golint
	/* Where do not contain value. */
	NOT_CONTAINS AlertsMutingRuleConditionOperator = "NOT_CONTAINS" // nolint:golint
	/* Where does not end with. */
	NOT_ENDS_WITH AlertsMutingRuleConditionOperator = "NOT_ENDS_WITH" // nolint:golint
	/* Where do not equal value. */
	NOT_EQUALS AlertsMutingRuleConditionOperator = "NOT_EQUALS" // nolint:golint
	/* Where not in value. */
	NOT_IN AlertsMutingRuleConditionOperator = "NOT_IN" // nolint:golint
	/* Where does not start with. */
	NOT_STARTS_WITH AlertsMutingRuleConditionOperator = "NOT_STARTS_WITH" // nolint:golint
	/* Where starts with. */
	STARTS_WITH AlertsMutingRuleConditionOperator = "STARTS_WITH" // nolint:golint
)

// nolint:golint
type AlertsMutingRuleInput struct {
	/* The condition that defines which violations to target. */
	Condition AlertsMutingRuleConditionGroupInput `json:"condition"`

	/* The description of the MutingRule. */
	Description string `json:"description"`

	/* Whether the MutingRule is enabled */
	Enabled bool `json:"enabled"`

	/* The name of the MutingRule. */
	Name string `json:"name"`
}

// nolint:golint
type AlertsPoliciesSearchCriteriaInput struct {
	/* The list of policy ids to return. */
	IDs []int `json:"ids"`
}

// nolint:golint
type AlertsPoliciesSearchResultSet struct {

	/* Cursor pointing to the end of the current page of Policy records. Null if final page. */
	NextCursor string `json:"nextCursor"`

	/* Set of Policies returned for the supplied cursor and criteria. */
	Policies []AlertsPolicy `json:"policies"`

	/* Total number of Policy records for the given search criteria. */
	TotalCount int `json:"totalCount"`
}

// nolint:golint
type AlertsPolicy struct {

	/* The account id of the Policy. */
	AccountID int `json:"accountId"`

	/* Primary key for Policies. */
	ID int `json:"id,string"`

	/* Determines how incidents are created for critical violations of the conditions contained in the policy. */
	IncidentPreference AlertsIncidentPreference `json:"incidentPreference"`

	/* A text description of the policy for display purposes. */
	Name string `json:"name"`
}

// nolint:golint
type AlertsPolicyInput struct {
	/* Determines how incidents are created for critical violations of the conditions contained in the policy. */
	IncidentPreference AlertsIncidentPreference `json:"incidentPreference"`

	/* A text description of the policy for display purposes. */
	Name string `json:"name"`
}

// nolint:golint
type AlertsPolicyUpdateInput struct {
	/* Determines how incidents are created for critical violations of the conditions contained in the policy. */
	IncidentPreference AlertsIncidentPreference `json:"incidentPreference"`

	/* A text description of the policy for display purposes. */
	Name string `json:"name"`
}