To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:

myCar := struct {
  Make string
  Model string
} {
  Make: "tesla",
  Model: "model 3"
}

You can even nest anonymous structs as fields within other structs:

type car struct {
  Make string
  Model string
  Height int
  Width int
  // Wheel is a field containing an anonymous struct
  Wheel struct {
    Radius int
    Material string
  }
}
