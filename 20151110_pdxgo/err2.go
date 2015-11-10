  f, err := os.Create("foo")
  if err != nil {
    return err // What does this mean when it's passed up 3 levels?
  }
