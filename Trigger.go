package main

var TriggerTypes = [...]string {
  "gt",
  "gte",
  "lt",
  "lte",
  "eq",
  "change",
  "frozen",
  "live",
}

type Trigger struct {
  Url string
  TriggerType string
  ThresholdValue string
  EnvironmentId string
  StreamId string
}
