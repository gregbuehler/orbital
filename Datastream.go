package main

type Datastream struct {
  StreamID string
  Updated string
  tags []string
  Units string
  UnitType string
  UnitSymbol string
  MinValue string
  MaxValue string
  CurrentValue string

  Datapoints []Datapoint
}
