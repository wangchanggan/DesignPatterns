package behavioral_patterns

import "testing"

func TestVisitor(t *testing.T) {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}

	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)

		/*{"Radius":10}
		<Circle><Radius>10</Radius></Circle>
		{"Width":100,"Heigh":200}
		<Rectangle><Width>100</Width><Heigh>200</Heigh></Rectangle>*/
	}
}
