package filters

import "fmt"

func UnsupportedOperator(op Operator) error {
	return fmt.Errorf("failed to add filter: unsupported operator %s", op.String())
}

func InvalidScope(filterScope string) error {
	return fmt.Errorf("invalid filter scope: %s", filterScope)
}

func InvalidExpression(expression string) error {
	return fmt.Errorf("invalid filter expression: %s", expression)
}

func InvalidValue(value string) error {
	return fmt.Errorf("invalid filter value %s", value)
}

func InvalidEventName(event string) error {
	return fmt.Errorf("invalid event name in filter: %s", event)
}

func InvalidEventArgument(argument string) error {
	return fmt.Errorf("invalid filter event argument: %s", argument)
}

func InvalidContextField(field string) error {
	return fmt.Errorf("invalid event context field: %s", field)
}

func FailedToRetreiveHostNS() error {
	return fmt.Errorf("failed to retreive host mount namespace")
}
