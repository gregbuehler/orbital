package main

type Feed struct {
  ID string
  Title string
  Private string
  Tags []string
  Description string
  Feed string
  Status string
  Updated string
  Created string
  Creator string
  Version string
  Location string
  Website string
  Icon string
  UserLogin string
  
  Datastreams []Datastream
}
